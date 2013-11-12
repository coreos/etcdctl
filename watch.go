package main

import (
	"flag"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"os"
	"os/signal"
)

const WatchUsage = `usage: etcdctl [etcd flags] watch <key> [watch flags]
special flags: -f forever watch a key until CTRL+C
               -i watch from the given index`

const WatchAllUsage = `usage: etcdctl [etcd flags] watchAll <key> [watchAll flags]
special flags: -f forever watch a key until CTRL+C
               -i watch from the given index`

var (
	watchFlag = flag.NewFlagSet("watch", flag.ExitOnError)
	forever   = watchFlag.Bool("f", false, "forever watch at the key")
	index     = watchFlag.Int64("i", 0, "watch from the given index")
)

func init() {
	registerCommand("watch", WatchUsage, 1, 3, watch)
	registerCommand("watchAll", WatchAllUsage, 1, 3, watchAll)
}

func watch(args []string) error {
	key := args[0]
	watchFlag.Parse(args[1:])

	if *forever {

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		stop := make(chan bool)

		go func() {
			<-c
			stop <- true
			os.Exit(0)
		}()

		receiver := make(chan *etcd.Response)
		go client.Watch(key, uint64(*index), receiver, stop)

		for {
			resp := <-receiver
			if debug {
				fmt.Println(<-curlChan)
			}
			output(resp)
		}

	} else {
		resp, err := client.Watch(key, uint64(*index), nil, nil)
		if debug {
			fmt.Println(<-curlChan)
		}
		if err != nil {
			return err
		}
		output(resp)
	}

	return nil
}

func watchAll(args []string) error {
	key := args[0]
	watchFlag.Parse(args[1:])

	if *forever {

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		stop := make(chan bool)

		go func() {
			<-c
			stop <- true
			os.Exit(0)
		}()

		receiver := make(chan *etcd.Response)
		go client.WatchAll(key, uint64(*index), receiver, stop)

		for {
			resp := <-receiver
			if debug {
				fmt.Println(<-curlChan)
			}
			output(resp)
		}

	} else {
		resp, err := client.WatchAll(key, uint64(*index), nil, nil)
		if debug {
			fmt.Println(<-curlChan)
		}
		if err != nil {
			return err
		}
		output(resp)
	}

	return nil
}
