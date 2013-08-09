package main

import (
	"flag"
	"fmt"
	"crypto/md5"
	"io"
	"os"
	"github.com/coreos/go-etcd/etcd"
)

const setUsage = `usage: etcdctl [etcd flags] <command>

Commands:

  setadd <key> <value> [flags]
    --ttl to set add a value with a ttl to the set
  setremove <key> <value>
  setmembers <key>
  setismember <key> <value>

`

var (
	setAddFlag = flag.NewFlagSet("setadd", flag.ExitOnError)
	addTtl = setAddFlag.Int64("ttl", 0, "ttl of the key")

	setRemoveFlag = flag.NewFlagSet("setremove", flag.ExitOnError)
	// FIXME: add a '--all' flag to setremove
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
	registerCommand("setadd", setUsage, 3, 4, SetAdd)
	registerCommand("setremove", setUsage, 3, 3, SetRemove)
	registerCommand("setismember", setUsage, 3, 3, SetIsMember)
	registerCommand("setmembers", setUsage, 2, 2, SetMembers)
}

func SetAdd(args []string) error {

	setKey := args[1]
	value := args[2]

	// Create the set unless it exists
	if ! setExists(setKey) {
		headKey := getHeadKey(setKey)
		_, err := client.Set(headKey, "1", 0)
		if err != nil {
			return err
		}
	}

	key := fmt.Sprintf("%s/%s", setKey, hash(value))
	_, err := client.Set(key, value, uint64(*addTtl))
	if err != nil {
		return err
	}

	fmt.Println(value)

	return nil
}

func SetRemove(args []string) error {

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

func SetMembers(args []string) error {
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
			fmt.Printf("\"%s\"\n", resp.Value)
		}
	}

	return nil
}

func SetIsMember(args []string) error {
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
