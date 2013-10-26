package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"os"
)

var (
	client       *etcd.Client
	printVersion bool
	outputJSON   bool
	cluster      = ClusterValue{"http://localhost:4001"}
)

func output(resp *etcd.Response) {
	if outputJSON {
		b, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", b)
	} else {
		fmt.Println(resp.Value)
	}
}

func main() {
	flag.BoolVar(&printVersion, "version", false, "print the version and exit")
	flag.Var(&cluster, "C", "a comma seperated list of machine addresses in the cluster e.g. 127.0.0.1:4001,127.0.0.1:4002")
	flag.BoolVar(&outputJSON, "json", false, "Output server response in JSON format")
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

	commandArgs := args[1:]

	if len(commandArgs) > command.maxArgs || len(commandArgs) < command.minArgs {
		fmt.Println("wrong arguments provided")
		fmt.Println(command.usage)
		os.Exit(MalformedEtcdctlArguments)
	}

	if !client.SyncCluster() {
		fmt.Println("cannot sync with the given cluster")
		os.Exit(FailedToConnectToHost)
	}

	err := command.f(commandArgs)

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		os.Exit(ErrorFromEtcd)
	}
}
