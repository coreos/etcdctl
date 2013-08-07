package main

import (
	"flag"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"os"
)

var (
	cluster = flag.String("C", "0.0.0.0:4001", "a list of machine addresses in the cluster")
	client  = etcd.NewClient()
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		os.Exit(1)
	}

	commandName := args[0]

	command, ok := commands[commandName]

	if !ok {
		fmt.Println("wrong command provided")
		os.Exit(MalformedEtcdctlArguments)
	}

	if len(args) > command.maxArgs || len(args) < command.minArgs {
		fmt.Println("wrong arguments provided")
		fmt.Println(command.usage)
		os.Exit(MalformedEtcdctlArguments)
	}

	if !client.SyncCluster() {
		fmt.Println("cannot sync with the given cluster")
		os.Exit(FailedToConnectToHost)
	}

	err := command.f(args)

	if err != nil {
		fmt.Println(err)
		os.Exit(ErrorFromEtcd)
	}
}
