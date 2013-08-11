package main

import (
	"flag"
	"fmt"
	"crypto/md5"
	"io"
	"os"
	"github.com/coreos/go-etcd/etcd"
)

const sUsage = `usage: etcdctl [etcd flags] <command>

Commands:

  sadd <key> <value> [flags]
    --ttl to set add a value with a ttl to the set
  sdel <key> <value>
  smembers <key>
  sismember <key> <value>

`

var (
	saddFlag = flag.NewFlagSet("sadd", flag.ExitOnError)
	saddTtl = saddFlag.Int64("ttl", 0, "ttl of the key")
)

func hash(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getHeadKey(key string) string {
	return fmt.Sprintf("%s/set-%s", key, hash(key))
}

func setExists(key string) bool {
	headKey := getHeadKey(key)
	_, err := client.Get(headKey)
	return err == nil
}

func init() {
	registerCommand("sadd", sUsage, 3, 4, sadd)
	registerCommand("sdel", sUsage, 3, 3, sdel)
	registerCommand("sismember", sUsage, 3, 3, sismember)
	registerCommand("smembers", sUsage, 2, 2, smembers)
}

func sadd(args []string) error {

	setKey := args[1]
	value := args[2]
	saddFlag.Parse(args[3:])

	// Create the set unless it exists
	if ! setExists(setKey) {
		headKey := getHeadKey(setKey)
		_, err := client.Set(headKey, "1", 0)
		if err != nil {
			return err
		}
	}

	key := fmt.Sprintf("%s/%s", setKey, hash(value))
	_, err := client.Set(key, value, uint64(*saddTtl))
	if err != nil {
		return err
	}

	fmt.Println(value)

	return nil
}

func sdel(args []string) error {

	setKey := args[1]

	if ! setExists(setKey) {
		return fmt.Errorf("%s is not a set", setKey)
	}

	value := args[2]
	key := fmt.Sprintf("%s/%s", setKey, hash(value))
	_, err := client.Delete(key)
	if err != nil {
		err := err.(etcd.EtcdError)
		if err.ErrorCode == 100 {
			return etcd.EtcdError{
				ErrorCode: 100,
				Message: "Not In Set",
				Cause: setKey,
			}
		}
		return err
	}

	return nil
}

func smembers(args []string) error {
	setKey := args[1]

	if ! setExists(setKey) {
		return fmt.Errorf("%s is not a set", setKey)
	}

	resps, err := client.Get(setKey)
	if err != nil {
		return err
	}

	headKey := getHeadKey(setKey)
	for _, resp := range resps {
		if resp.Key != headKey {
			fmt.Printf("%s\n", resp.Value)
		}
	}

	return nil
}

func sismember(args []string) error {
	setKey := args[1]
	value := args[2]

	if ! setExists(setKey) {
		return fmt.Errorf("%s is not a set", setKey)
	}

	key := fmt.Sprintf("%s/%s", setKey, hash(value))
	_, err := client.Get(key)
	if err != nil {
		fmt.Println("false")
		os.Exit(1)
	} else {
		fmt.Println("true")
		os.Exit(0)
	}

	return nil
}
