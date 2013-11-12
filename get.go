package main

import (
	"flag"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
)

const GetUsage = `usage: etcdctl [etcd flags] get <key> [get flags]
special flags: --sort to return the result in sorted order
               --consistent to send request to the leader, thereby guranteeing that any earlier writes will be seen by the read`

const GetAllUsage = `usage: etcdctl [etcd flags] getAll <key> [getAll flags]
special flags: --sort set to true to return the result in sorted order
               --consistent to send request to the leader, thereby guranteeing that any earlier writes will be seen by the read`

var (
	getFlag = flag.NewFlagSet("get", flag.ExitOnError)
	sorted  = getFlag.Bool("sort", false,
		"Return the results in sorted order or not (default to false)")
	consistent = getFlag.Bool("consistent", false,
		"Send the request to the leader or not (default to false)")
)

func init() {
	registerCommand("get", GetUsage, 1, 2, get)
	registerCommand("getAll", GetAllUsage, 1, 2, getAll)
}

func get(args []string) error {
	if *consistent {
		client.SetConsistency(etcd.STRONG_CONSISTENCY)
	} else {
		client.SetConsistency(etcd.WEAK_CONSISTENCY)
	}

	key := args[0]
	getFlag.Parse(args[1:])
	resp, err := client.Get(key, *sorted)
	if debug {
		fmt.Println(<-curlChan)
	}

	if err != nil {
		return err
	}

	output(resp)

	return nil
}

func getAll(args []string) error {
	if *consistent {
		client.SetConsistency(etcd.STRONG_CONSISTENCY)
	} else {
		client.SetConsistency(etcd.WEAK_CONSISTENCY)
	}

	key := args[0]
	getFlag.Parse(args[1:])
	resp, err := client.GetAll(key, *sorted)
	if debug {
		fmt.Println(<-curlChan)
	}

	if err != nil {
		return err
	}

	output(resp)

	return nil
}
