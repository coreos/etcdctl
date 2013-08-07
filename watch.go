package main

import (
	"flag"
	"fmt"
	"github.com/coreos/etcd/store"
	"os"
	"os/signal"
)

const WatchUsage = `usage: etcdctl [etcd flags] watch <key> [watch flags]
special flags: -f forever watch a key until CTRL+C
               -i watch from the given index`

var (
	watchFlag = flag.NewFlagSet("watch", flag.ExitOnError)
	forever   = watchFlag.Bool("f", false, "forever watch at the key")
	index     = watchFlag.Int64("i", 0, "watch from the given index")
)

func init() {
	registerCommand("watch", WatchUsage, 2, 6, watch)
}

func watch(args []string) error {
	key := args[1]
	watchFlag.Parse(args[2:])

	if *forever {

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		go func() {
			<-c
			stop <- true
			os.Exit(0)
		}()

		receiver := make(chan *store.Response)
		stop := make(chan bool)
		go client.Watch(key, uint64(*index), receiver, stop)

		for {
			resp := <-receiver
			fmt.Println(resp.Action, " ", resp.Key, " ", resp.Value)
		}

	} else {
		resp, err := client.Watch(key, uint64(*index), nil, nil)
		if err != nil {
			return err
		}
		fmt.Println(resp.Action, " ", resp.Key, " ", resp.Value)
	}

	return nil
}
