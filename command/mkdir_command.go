package command

import (
	"errors"

	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/cobra"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
)

var (
	mkDirCmd     *cobra.Command
	mkDirTTLFlag int
)

func init() {
	mkDirCmd = &cobra.Command{
		Use:   "mkdir",
		Short: "make a new directory",
		Run: func(cmd *cobra.Command, args []string) {
			handleDir(cmd, args, makeDirCommandFunc)
		},
	}
	mkDirCmd.Flags().IntVar(&mkDirTTLFlag, "ttl", 0, "directory time-to-live")
}

// NewMakeDirCommand returns the Cobra command for "mkdir".
func MakeDirCommand() *cobra.Command {
	return mkDirCmd
}

// makeDirCommandFunc executes the "mkdir" command.
func makeDirCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("key required")
	}
	key := args[0]
	return client.CreateDir(key, uint64(mkDirTTLFlag))
}
