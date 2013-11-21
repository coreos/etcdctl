package command

import (
	"errors"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

// NewCreateDirCommand returns the CLI command for "createDir".
func NewCreateDirCommand() cli.Command {
	return cli.Command{
		Name:  "createDir",
		Usage: "create a new directory",
		Flags: []cli.Flag{
			cli.IntFlag{"ttl", 0, "key time-to-live"},
		},
		Action: func(c *cli.Context) {
			handle(c, createDirCommandFunc)
		},
	}
}

// createDirCommandFunc executes the "createDir" command.
func createDirCommandFunc(c *cli.Context, client *etcd.Client) (*etcd.Response, error) {
	if len(c.Args()) == 0 {
		return nil, errors.New("Key required")
	}
	key := c.Args()[0]
	ttl := c.Int("ttl")

	return client.CreateDir(key, uint64(ttl))
}
