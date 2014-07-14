package command

import (
	"errors"
	"github.com/coreos/go-etcd/etcd"
	"github.com/spf13/cobra"
)

var setDirCmd *cobra.Command
var setDirTtlFlag int

func init() {

	setDirCmd = &cobra.Command{
		Use:   "setdir",
		Short: "create a new or existing directory",

		Run: func(cmd *cobra.Command, args []string) {
			handleDir(cmd, args, setDirCommandFunc)
		},
	}

	setDirCmd.Flags().IntVarP(&setDirTtlFlag, "ttl", "", 0, "key time-to-live ")
}

// NewSetDirCommand returns the CLI command for "setDir".
func SetDirCommand() *cobra.Command {

	return setDirCmd
}

func setDirCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("Key required")
	}
	key := args[0]
	ttl := ttlFlag

	return client.SetDir(key, uint64(ttl))
}
