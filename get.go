package main

import (
	"flag"
	"fmt"
)

const GetUsage = `usage: etcdctl [etcd flags] get <key> [get flags]
special flags: --sort set to true to return the result in sorted order`

const GetAllUsage = `usage: etcdctl [etcd flags] getAll <key> [getAll flags]
special flags: --sort set to true to return the result in sorted order`

var (
	getFlag = flag.NewFlagSet("get", flag.ExitOnError)
	sorted  = getFlag.Bool("sort", false,
		"Return the results in sorted order or not (default to false)")
)

func init() {
	registerCommand("get", GetUsage, 1, 2, get)
	registerCommand("getAll", GetAllUsage, 1, 2, getAll)
}

func get(args []string) error {
	key := args[0]
	getFlag.Parse(args[1:])
	resp, err := client.Get(key, *sorted)
	if err != nil {
		return err
	}

	if resp.Value != "" {
		fmt.Println(resp.Value)
	}

	return nil
}

func getAll(args []string) error {
	key := args[0]
	getFlag.Parse(args[1:])
	resp, err := client.GetAll(key, *sorted)
	if err != nil {
		return err
	}

	if resp.Value != "" {
		fmt.Println(resp.Value)
	}

	return nil
}
