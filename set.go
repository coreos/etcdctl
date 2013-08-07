package main

import (
	"flag"
	"fmt"
)

const SetUsage = `usage: etcdctl [etcd flags] set <key> <value> [set flags]
special flags: --ttl to set a key with ttl`

var (
	setFlag = flag.NewFlagSet("set", flag.ExitOnError)
	ttl     = setFlag.Int64("ttl", 0, "ttl of the key")
)

func init() {
	registerCommand("set", SetUsage, 3, 5, set)
}

func set(args []string) error {
	key := args[1]
	value := args[2]
	setFlag.Parse(args[3:])
	resp, err := client.Set(key, value, uint64(*ttl))
	if err != nil {
		return err
	}
	fmt.Println(resp.Value)

	return nil
}
