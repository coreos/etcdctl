package command

import (
	"errors"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/cobra"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
)

var updateDirCmd *cobra.Command

//flags
var updateDirTTLFlag int

func init() {
	updateDirCmd = &cobra.Command{
		Use:   "updatedir",
		Short: "update an existing directory",
		Run: func(cmd *cobra.Command, args []string) {
			handleDir(cmd, args, updateDirCommandFunc)
		},
	}

	updateDirCmd.Flags().IntVar(&updateDirTTLFlag, "ttl", 0, "key time-to-live")

}

// UpdateDirCommand returns the sub command for "updatedir".
func UpdateDirCommand() *cobra.Command {
	return updateDirCmd
}

// updateDirCommandFunc executes the "updateDir" command.
func updateDirCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("Key required")
	}
	key := args[0]
	ttl := updateDirTTLFlag

	return client.UpdateDir(key, uint64(ttl))
}
