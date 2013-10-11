package main

import (
	"flag"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"os"
)

var (
	client *etcd.Client

	printVersion bool
)

func main() {
	flag.BoolVar(&printVersion, "version", false, "print the version and exit")

	cluster := ClusterValue{"http://localhost:4001"}
	flag.Var(&cluster, "C", "a comma seperated list of machine addresses in the cluster e.g. 127.0.0.1:4001,127.0.0.1:4002")
	flag.Parse()

	if printVersion {
		fmt.Println(releaseVersion)
		os.Exit(0)
	}

	client = etcd.NewClient(cluster.GetMachines())

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
