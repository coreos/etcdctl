package command

import (
	"errors"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

// NewDeleteCommand returns the CLI command for "delete".
func NewDeleteCommand() cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "remove a key",
		Flags: []cli.Flag{
			cli.BoolFlag{"recursive", "deletes the key and all child keys"},
		},
		Action: func(c *cli.Context) {
			handle(c, deleteCommandFunc)
		},
	}
}

// deleteCommandFunc executes the "delete" command.
func deleteCommandFunc(c *cli.Context, client *etcd.Client) (*etcd.Response, error) {
	if len(c.Args()) == 0 {
		return nil, errors.New("Key required")
	}
	key := c.Args()[0]
	recursive := c.Bool("recursive")
	if recursive {
		return client.DeleteAll(key)
	} else {
		return client.Delete(key)
	}
}
