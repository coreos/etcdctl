package command

import (
	"errors"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

// NewUpdateCommand returns the CLI command for "update".
func NewUpdateCommand() cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "update an existing key with a given value",
		Flags: []cli.Flag{
			cli.IntFlag{"ttl", 0, "key time-to-live"},
		},
		Action: func(c *cli.Context) {
			handleKey(c, updateCommandFunc)
		},
	}
}

// updateCommandFunc executes the "update" command.
func updateCommandFunc(c *cli.Context, client *etcd.Client) (*etcd.Response, error) {
	if len(c.Args()) == 0 {
		return nil, errors.New("Key required")
	} else if len(c.Args()) == 1 {
		return nil, errors.New("Value required")
	}
	key := c.Args()[0]
	value := c.Args()[1]
	ttl := c.Int("ttl")

	return client.Update(key, value, uint64(ttl))
}
