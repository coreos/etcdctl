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
		Usage: "delete a key",
		Flags: []cli.Flag{
			cli.BoolFlag{"dir", "deletes the key if it is an empty directory or a key-value pair"},
			cli.BoolFlag{"recursive", "deletes the key and all child keys(if it is a directory)"},
		},
		Action: func(c *cli.Context) {
			handleKey(c, deleteCommandFunc)
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
	dir := c.Bool("dir")

	if recursive || !dir {
		return client.Delete(key, recursive)
	}

	return client.DeleteDir(key)
}
