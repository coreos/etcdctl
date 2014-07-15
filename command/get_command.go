package command

import (
	"errors"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"github.com/spf13/cobra"
	"os"
)

var getCmd *cobra.Command

//flags
var getSortFlag bool
var getConsistentFlag bool

func init() {
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "retrieve the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			handleGet(cmd, args, getCommandFunc)
		},
	}

	getCmd.Flags().BoolVarP(&getConsistentFlag, "consistent", "", false,
		"send request to the leader, thereby guranteeing that any earlier writes will be seen by the read")
	getCmd.Flags().BoolVarP(&getSortFlag, "sort", "", false, "returns result in sorted order")
}

// NewGetCommand returns the CLI command for "get".
func GetCommand() *cobra.Command {
	return getCmd
}

// handleGet handles a request that intends to do get-like operations.
func handleGet(cmd *cobra.Command, args []string, fn handlerFunc) {
	handlePrint(cmd, args, fn, printGet)
}

// printGet writes error message when getting the value of a directory.
func printGet(resp *etcd.Response, format string) {
	if resp.Node.Dir {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("%s: is a directory", resp.Node.Key))
		os.Exit(1)
	}

	printKey(resp, format)
}

// getCommandFunc executes the "get" command.
func getCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("Key required")
	}
	key := args[0]
	consistent := getConsistentFlag
	sorted := getSortFlag

	// Setup consistency on the client.
	if consistent {
		client.SetConsistency(etcd.STRONG_CONSISTENCY)
	} else {
		client.SetConsistency(etcd.WEAK_CONSISTENCY)
	}

	// Retrieve the value from the server.
	return client.Get(key, sorted, false)
}
