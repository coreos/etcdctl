package command

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

type handlerFunc func(*cli.Context, *etcd.Client) (*etcd.Response, error)

var curlChan = make(chan string, 10)

// rawhandle wraps the command function handlers and sets up the
// environment but performs no output formatting.
func rawhandle(c *cli.Context, fn handlerFunc) (*etcd.Response, error) {
	peers := c.GlobalStringSlice("C")
	client := etcd.NewClient(peers)

	// Set channel to receive cURL output.
	etcd.SetCurlChan(curlChan)

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

	if c.GlobalBool("debug") {
		select {
		case s := <-curlChan:
			fmt.Fprintln(os.Stderr, s)
		default:
		}
	}

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
		fmt.Println(resp.Value)
		// This style emulates tools like getfacl
		fmt.Println("# modified-index:", resp.ModifiedIndex)
		fmt.Println("# ttl:", resp.TTL)
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
