package command

import (
	"errors"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

// NewCreateCommand returns the CLI command for "create".
func NewCreateCommand() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "create a new key with a given value",
		Flags: []cli.Flag{
			cli.IntFlag{"ttl", 0, "key time-to-live"},
		},
		Action: func(c *cli.Context) {
			handleKey(c, createCommandFunc)
		},
	}
}

// createCommandFunc executes the "create" command.
func createCommandFunc(c *cli.Context, client *etcd.Client) (*etcd.Response, error) {
	if len(c.Args()) == 0 {
		return nil, errors.New("Key required")
	} else if len(c.Args()) == 1 {
		return nil, errors.New("Value required")
	}
	key := c.Args()[0]
	value := c.Args()[1]
	ttl := c.Int("ttl")

	return client.Create(key, value, uint64(ttl))
}
