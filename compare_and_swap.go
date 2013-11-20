package main

import (
	"flag"
	"fmt"
)

const CompareAndSwapUsage = `usage: etcdctl [etcd flags] compareAndSwap <key> <value> [compareAndSwap flags]
either prevvalue or previndex needs to be given
special flags: --ttl to set a key with ttl
			   --prevvalue to set the previous value
			   --previndex to set the previous index`

var (
	compareAndSwapFlag   = flag.NewFlagSet("compareAndSwap", flag.ExitOnError)
	compareAndSwapTtl    = compareAndSwapFlag.Uint64("ttl", 0, "ttl of the key")
	compareAndSwapPvalue = compareAndSwapFlag.String("prevvalue", "", "previous value")
	compareAndSwapPindex = compareAndSwapFlag.Uint64("previndex", 0, "previous index")
)

func init() {
	// The minimum number of arguments is 3 because
	// there needs to be either pvalue or pindex
	registerCommand("compareAndSwap", CompareAndSwapUsage, 3, 6, compareAndSwap)
}

func compareAndSwap(args []string) error {
	key := args[0]
	value := args[1]
	compareAndSwapFlag.Parse(args[2:])
	resp, err := client.CompareAndSwap(key, value,
		*compareAndSwapTtl, *compareAndSwapPvalue, *compareAndSwapPindex)
	if debug {
		fmt.Println(<-curlChan)
	}
	if err != nil {
		return err
	}

	output(resp)
	return nil
}
