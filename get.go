package main

import (
	"fmt"
)

const GetUsage = `usage: etcdctl [etcd flags] get <key>`

func init() {
	registerCommand("get", GetUsage, 2, 2, get)
}

func get(args []string) error {
	key := args[1]
	resps, err := client.Get(key)
	if err != nil {
		return err
	}
	for _, resp := range resps {
		if resp.Value != "" {
			fmt.Println(resp.Value)
		}
	}
	return nil
}
