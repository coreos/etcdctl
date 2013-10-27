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
	outputFormat string
	debug        bool
	cluster      = ClusterValue{"http://localhost:4001"}
)

func output(resp *etcd.Response) {
	switch outputFormat {
	case "json":
		b, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", b)
	case "full":
		fmt.Printf("%v %v\n", resp.Index, resp.Value)
	default:
		fmt.Println(resp.Value)
	}
}

func main() {
	flag.BoolVar(&printVersion, "version", false, "print the version and exit")
	flag.Var(&cluster, "C", "a comma seperated list of machine addresses in the cluster e.g. 127.0.0.1:4001,127.0.0.1:4002")
	flag.StringVar(&outputFormat, "format", "", "Output server response in the given format, either `json` or `full`")
	flag.BoolVar(&debug, "debug", false, "Output cURL commands which can be used to re-produce the request")
	flag.Parse()

	if printVersion {
		fmt.Println(releaseVersion)
		os.Exit(0)
	}

	if debug {
		etcd.SetPrintCurl(true)
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
