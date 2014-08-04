package command

import (
	"errors"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/cobra"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
)

var setDirCmd *cobra.Command
var setDirTtlFlag int

func init() {

	setDirCmd = &cobra.Command{
		Use:   "setdir",
		Short: "create a new directory",

		Run: func(cmd *cobra.Command, args []string) {
			handleDir(cmd, args, setDirCommandFunc)
		},
	}

	setDirCmd.Flags().IntVar(&setDirTtlFlag, "ttl", 0, "key time-to-live ")
}

// NewSetDirCommand returns the Cobra command for "setDir".
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
