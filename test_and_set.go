package main

import (
	"errors"
	"flag"
	"fmt"
)

const TestAndSetUsage = `usage: etcdctl [etcd flags] testAndSet <key> <prevValue> <value> [testAndSet flags]
special flags: --ttl to set a key with ttl`

var (
	testAndSetFlag = flag.NewFlagSet("testAndSet", flag.ExitOnError)
	testAndSetTtl  = testAndSetFlag.Int64("ttl", 0, "ttl of the key")
)

func init() {
	registerCommand("testAndSet", TestAndSetUsage, 4, 6, testAndSet)
}

func testAndSet(args []string) error {
	key := args[1]
	prevValue := args[2]
	value := args[3]
	testAndSetFlag.Parse(args[4:])
	resp, success, err := client.TestAndSet(key, prevValue, value, uint64(*testAndSetTtl))

	if err != nil {
		return err
	}

	if success {
		fmt.Println(resp.Value)
		return nil
	}

	return errors.New("TestAndSet: prevValue does not match the current value of the given key")
}
