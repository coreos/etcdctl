package command

import (
	"errors"

	"github.com/coreos/etcdctl/third_party/github.com/codegangsta/cli"
	"github.com/coreos/etcdctl/third_party/github.com/coreos/go-etcd/etcd"
)

// NewRemoveCommand returns the CLI command for "rm".
func NewRemoveCommand() cli.Command {
	return cli.Command{
		Name:	"rm",
		Usage:	"remove a key",
		Flags: []cli.Flag{
			cli.BoolFlag{"dir", "removes the key if it is an empty directory or a key-value pair"},
			cli.BoolFlag{"recursive", "removes the key and all child keys(if it is a directory)"},
		},
		Action: func(c *cli.Context) {
			handleKey(c, removeCommandFunc)
		},
	}
}

// removeCommandFunc executes the "rm" command.
func removeCommandFunc(c *cli.Context, client *etcd.Client) (*etcd.Response, error) {
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
