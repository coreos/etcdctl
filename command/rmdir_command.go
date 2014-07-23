package command

import (
	"errors"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/cobra"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
)

var rmDirCmd *cobra.Command

// flags
// there is a recursive in the rm command.
func init() {
	rmDirCmd = &cobra.Command{
		Use:   "rmdir",
		Short: "removes the key if it is an empty directory or a key-value pair",
		Run: func(cmd *cobra.Command, args []string) {
			handleDir(cmd, args, removeDirCommandFunc)
		},
	}

}

// RemoveDirCommand returns the CLI command for "rmdir".
func RemoveDirCommand() *cobra.Command {
	return rmDirCmd
}

// removeDirCommandFunc executes the "rmdir" command.
func removeDirCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("Key required")
	}
	key := args[0]

	return client.DeleteDir(key)
}
