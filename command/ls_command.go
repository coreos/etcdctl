package command

import (
	"fmt"

	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/cobra"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
)

var lsCmd *cobra.Command
var lsRecursiveFlag bool

func init() {
	lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "retrieve a directory",
		Run: func(cmd *cobra.Command, args []string) {
			handleLs(cmd, args, lsCommandFunc)
		},
	}

	lsCmd.Flags().BoolVar(&lsRecursiveFlag, "recursive", false, "returns all values for key and child keys")
}

func LsCommand() *cobra.Command {
	return lsCmd
}

// handleLs handles a request that intends to do ls-like operations.
func handleLs(cmd *cobra.Command, args []string, fn handlerFunc) {
	handlePrint(cmd, args, fn, printLs)
}

// printLs writes a response out in a manner similar to the `ls` command in unix.
// Non-empty directories list their contents and files list their name.
func printLs(resp *etcd.Response, format string) {
	if !resp.Node.Dir {
		fmt.Println(resp.Node.Key)
	}
	for _, node := range resp.Node.Nodes {
		rPrint(node)
	}
}

// lsCommandFunc executes the "ls" command.
func lsCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	key := "/"
	if len(args) != 0 {
		key = args[0]
	}
	recursive := lsRecursiveFlag

	// Retrieve the value from the server.
	return client.Get(key, false, recursive)
}

// rPrint recursively prints out the nodes in the node structure.
func rPrint(n *etcd.Node) {
	fmt.Println(n.Key)
	for _, node := range n.Nodes {
		rPrint(node)
	}
}
