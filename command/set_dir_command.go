package command

import (
	"errors"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

// NewSetDirCommand returns the CLI command for "setDir".
func NewSetDirCommand() cli.Command {
	return cli.Command{
		Name:  "setdir",
		Usage: "create a new or existing directory",
		Flags: []cli.Flag{
			cli.IntFlag{"ttl", 0, "key time-to-live"},
		},
		Action: func(c *cli.Context) {
			handleKey(c, setDirCommandFunc)
		},
	}
}

// setDirCommandFunc executes the "setDir" command.
func setDirCommandFunc(c *cli.Context, client *etcd.Client) (*etcd.Response, error) {
	if len(c.Args()) == 0 {
		return nil, errors.New("Key required")
	}
	key := c.Args()[0]
	ttl := c.Int("ttl")

	return client.SetDir(key, uint64(ttl))
}
