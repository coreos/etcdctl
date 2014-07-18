package command

import (
	"errors"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/joshi4/cobra"
	"os"
)

var updateCmd *cobra.Command

//flags
var updateTTLFlag int

func init() {
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "update an existing key with a given value",
		Run: func(cmd *cobra.Command, args []string) {
			handleKey(cmd, args, updateCommandFunc)
		},
	}

	updateCmd.Flags().IntVarP(&updateTTLFlag, "ttl", "", 0, "key time-to-live")

}

// UpdateCommand returns the updateCommand to be added onto the root.
func UpdateCommand() *cobra.Command {
	return updateCmd
}

// updateCommandFunc executes the "update" command.
func updateCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("Key required")
	}
	key := args[0]
	value, err := argOrStdin(args, os.Stdin, 1)
	if err != nil {
		return nil, errors.New("Value required")
	}

	ttl := updateTTLFlag

	return client.Update(key, value, uint64(ttl))
}
