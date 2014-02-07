package command

import (
	"errors"

	"github.com/coreos/etcdctl/third_party/github.com/codegangsta/cli"
	"github.com/coreos/etcdctl/third_party/github.com/coreos/go-etcd/etcd"
)

// NewGetCommand returns the CLI command for "get".
func NewGetCommand() cli.Command {
	return cli.Command{
		Name:	"get",
		Usage:	"retrieve the value of a key",
		Flags: []cli.Flag{
			cli.BoolFlag{"sort", "returns result in sorted order"},
			cli.BoolFlag{"consistent", "send request to the leader, thereby guranteeing that any earlier writes will be seen by the read"},
			cli.BoolFlag{"recursive", "returns all values for key and child keys"},
		},
		Action: func(c *cli.Context) {
			handleKey(c, getCommandFunc)
		},
	}
}

// getCommandFunc executes the "get" command.
func getCommandFunc(c *cli.Context, client *etcd.Client) (*etcd.Response, error) {
	if len(c.Args()) == 0 {
		return nil, errors.New("Key required")
	}
	key := c.Args()[0]
	recursive := c.Bool("recursive")
	consistent := c.Bool("consistent")
	sorted := c.Bool("sort")

	// Setup consistency on the client.
	if consistent {
		client.SetConsistency(etcd.STRONG_CONSISTENCY)
	} else {
		client.SetConsistency(etcd.WEAK_CONSISTENCY)
	}

	// Retrieve the value from the server.
	return client.Get(key, sorted, recursive)
}
