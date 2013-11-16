package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"os"
	"sort"
)

var (
	client       *etcd.Client
	printVersion bool
	outputFormat string
	debug        bool
	cluster      = ClusterValue{"http://localhost:4001"}
	curlChan     chan string
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
		fmt.Printf("Index: %v\nValue: %v\n", resp.ModifiedIndex, resp.Value)
	case "value-only":
		fmt.Println(resp.Value)
	}
}

func main() {
	flag.BoolVar(&printVersion, "version", false, "print the version and exit")
	flag.Var(&cluster, "C", "a comma seperated list of machine addresses in the cluster e.g. 127.0.0.1:4001,127.0.0.1:4002")
	flag.StringVar(&outputFormat, "format", "json", "Output server response in the given format, either `json`, `full`, or `value-only`")
	flag.BoolVar(&debug, "debug", false, "Output cURL commands which can be used to re-produce the request")
	flag.Parse()

	if printVersion {
		fmt.Println(releaseVersion)
		os.Exit(0)
	}

	if debug {
		// Making the channel buffered to avoid potential deadlocks
		curlChan = make(chan string, 10)
		etcd.SetCurlChan(curlChan)
	}

	client = etcd.NewClient(cluster.GetMachines())

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: etcdctl [flags] [command] [flags for command]\n")
		fmt.Println("Available flags include:\n")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Available commands:")
		fmt.Println()
		slice := make([]string, len(commands))
		i := 0
		for k, _ := range commands {
			slice[i] = k
			i++
		}
		sort.Strings(slice)
		for _, c := range slice {
			fmt.Println(c)
		}
		fmt.Println()
		fmt.Println(`To see the full usage for a specific command, run "etcdctl [command]"`)
		fmt.Println()
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
