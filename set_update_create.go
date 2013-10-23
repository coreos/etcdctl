package main

import (
	"flag"
	"fmt"
)

const SetUsage = `usage: etcdctl [etcd flags] set <key> <value> [set flags]
special flags: --ttl to set a key with ttl`

const CreateUsage = `usage: etcdctl [etcd flags] create <key> <value> [create flags]
special flags: --ttl to create a key with ttl`

const UpdateUsage = `usage: etcdctl [etcd flags] update <key> <value> [udpate flags]
special flags: --ttl to update a key with ttl`

const SetDirUsage = `usage: etcdctl [etcd flags] setDir <key> <value> [setDir flags]
special flags: --ttl to set a directory with ttl`

const CreateDirUsage = `usage: etcdctl [etcd flags] createDir <key> <value> [createDir flags]
special flags: --ttl to create a directory with ttl`

const UpdateDirUsage = `usage: etcdctl [etcd flags] updateDir <key> <value> [udpateDir flags]
special flags: --ttl to update a directory with ttl`

var (
	setFlag = flag.NewFlagSet("set", flag.ExitOnError)
	ttl     = setFlag.Int64("ttl", 0, "ttl of the key")
)

func init() {
	registerCommand("set", SetUsage, 2, 3, set)
	registerCommand("create", CreateUsage, 2, 3, create)
	registerCommand("update", UpdateUsage, 2, 3, update)
	registerCommand("setDir", SetDirUsage, 1, 2, setDir)
	registerCommand("createDir", CreateDirUsage, 1, 2, createDir)
	registerCommand("updateDir", UpdateDirUsage, 1, 2, updateDir)
}

func set(args []string) error {
	key := args[0]
	value := args[1]
	setFlag.Parse(args[2:])
	resp, err := client.Set(key, value, uint64(*ttl))
	if err != nil {
		return err
	}
	fmt.Println(resp.Value)

	return nil
}

func create(args []string) error {
	key := args[0]
	value := args[1]
	setFlag.Parse(args[2:])
	resp, err := client.Create(key, value, uint64(*ttl))
	if err != nil {
		return err
	}
	fmt.Println(resp.Value)

	return nil
}

func update(args []string) error {
	key := args[0]
	value := args[1]
	setFlag.Parse(args[2:])
	resp, err := client.Update(key, value, uint64(*ttl))
	if err != nil {
		return err
	}
	fmt.Println(resp.Value)

	return nil
}

func setDir(args []string) error {
	key := args[0]
	setFlag.Parse(args[1:])
	resp, err := client.SetDir(key, uint64(*ttl))
	if err != nil {
		return err
	}
	fmt.Println(resp.Value)

	return nil
}

func createDir(args []string) error {
	key := args[0]
	setFlag.Parse(args[1:])
	resp, err := client.CreateDir(key, uint64(*ttl))
	if err != nil {
		return err
	}
	fmt.Println(resp.Value)

	return nil
}

func updateDir(args []string) error {
	key := args[0]
	setFlag.Parse(args[1:])
	resp, err := client.UpdateDir(key, uint64(*ttl))
	if err != nil {
		return err
	}
	fmt.Println(resp.Value)

	return nil
}
