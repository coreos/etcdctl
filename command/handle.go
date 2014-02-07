package command

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/coreos/etcdctl/third_party/github.com/codegangsta/cli"
	"github.com/coreos/etcdctl/third_party/github.com/coreos/go-etcd/etcd"
)

type handlerFunc func(*cli.Context, *etcd.Client) (*etcd.Response, error)
type printFunc func(*etcd.Response, string)

// dumpCURL blindly dumps all curl output to os.Stderr
func dumpCURL(client *etcd.Client) {
	client.OpenCURL()
	for {
		fmt.Fprintf(os.Stderr, "Curl-Example: %s\n", client.RecvCURL())
	}
}

// rawhandle wraps the command function handlers and sets up the
// environment but performs no output formatting.
func rawhandle(c *cli.Context, fn handlerFunc) (*etcd.Response, error) {
	peers := c.GlobalString("peers")
	client := etcd.NewClient(trimsplit(peers, ","))

	if c.GlobalBool("debug") {
		go dumpCURL(client)
	}

	// Sync cluster.
	ok := client.SyncCluster()
	if c.GlobalBool("debug") {
		fmt.Fprintf(os.Stderr, "Cluster-Peers: %s\n",
			strings.Join(client.GetCluster(), " "))
	}

	if !ok {
		fmt.Println("Cannot sync with the cluster")
		os.Exit(FailedToConnectToHost)
	}

	// Execute handler function.
	return fn(c, client)
}

// handlePrint wraps the command function handlers to parse global flags
// into a client and to properly format the response objects.
func handlePrint(c *cli.Context, fn handlerFunc, pFn printFunc) {
	resp, err := rawhandle(c, fn)

	// Print error and exit, if necessary.
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(ErrorFromEtcd)
	}

	if resp != nil {
		pFn(resp, c.GlobalString("output"))
	}
}

// handleKey handles a request that wants to do operations on a single key.
func handleKey(c *cli.Context, fn handlerFunc) {
	handlePrint(c, fn, printKey)
}

// printKey writes the etcd response to STDOUT in the given format.
func printKey(resp *etcd.Response, format string) {
	// printKey is only for keys, error on directories
	if resp.Node.Dir == true {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Cannot print key [%s: Is a directory]", resp.Node.Key))
		os.Exit(1)
	}

	// Format the result.
	switch format {
	case "simple":
		fmt.Println(resp.Node.Value)
	case "extended":
		// Extended prints in a rfc2822 style format
		fmt.Println("Key:", resp.Node.Key)
		fmt.Println("Created-Index:", resp.Node.CreatedIndex)
		fmt.Println("Modified-Index:", resp.Node.ModifiedIndex)

		if resp.PrevNode != nil {
			fmt.Println("PrevNode.Value:", resp.PrevNode.Value)
		}

		fmt.Println("TTL:", resp.Node.TTL)
		fmt.Println("Etcd-Index:", resp.EtcdIndex)
		fmt.Println("Raft-Index:", resp.RaftIndex)
		fmt.Println("Raft-Term:", resp.RaftTerm)
		fmt.Println("")
		fmt.Println(resp.Node.Value)
	case "json":
		b, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	default:
		fmt.Fprintln(os.Stderr, "Unsupported output format:", format)
	}
}
