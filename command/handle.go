package command

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

type handlerFunc func(*cli.Context, *etcd.Client) (*etcd.Response, error)

// dumpCURL blindly dumps all curl output to os.Stderr
func dumpCURL(client *etcd.Client) {
	client.OpenCURL()
	for {
		fmt.Fprintln(os.Stderr, client.RecvCURL())
	}
}

// rawhandle wraps the command function handlers and sets up the
// environment but performs no output formatting.
func rawhandle(c *cli.Context, fn handlerFunc) (*etcd.Response, error) {
	peers := c.GlobalStringSlice("C")
	client := etcd.NewClient(peers)

	if c.GlobalBool("debug") {
		go dumpCURL(client)
	}

	// Sync cluster.    
	if !client.SyncCluster() {
		fmt.Println("cannot sync with the given cluster")
		os.Exit(FailedToConnectToHost)
	}

	// Execute handler function.
	return fn(c, client)
}

// handle wraps the command function handlers to parse global flags
// into a client and to properly format the response objects.
func handle(c *cli.Context, fn handlerFunc) {
	resp, err := rawhandle(c, fn)

	// Print error and exit, if necessary.
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(ErrorFromEtcd)
	}

	if resp != nil {
		printResponse(resp, c.GlobalString("output"))
	}
}

// printResponse writes the etcd response to STDOUT in the given format.
func printResponse(resp *etcd.Response, format string) {
	// Format the result.
	switch format {
	case "simple":
		fmt.Println(resp.Value)
	case "extended":
		// Extended prints in a rfc2822 style format
		fmt.Println("Key:", resp.Key)
		fmt.Println("Modified-Index:", resp.ModifiedIndex)
		fmt.Println("TTL:", resp.TTL)
		fmt.Println("")
		fmt.Println(resp.Value)
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
