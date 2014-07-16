package command

import (
	"errors"
	"github.com/joshi4/cobra"

	"github.com/coreos/go-etcd/etcd"
)

var mkDirCmd *cobra.Command
var mkDirTtlFlag int

func init() {

	mkDirCmd = &cobra.Command{
		Use:   "mkdir",
		Short: "make a new directory",
		Run: func(cmd *cobra.Command, args []string) {
			handleDir(cmd, args, makeDirCommandFunc)
		},
	}

	// ttl does not seem to be applicable for directories ??
	mkDirCmd.Flags().IntVarP(&mkDirTtlFlag, "ttl", "", 0, "directory time-to-live")

}

// NewMakeDirCommand returns the CLI command for "mkdir".
func MakeDirCommand() *cobra.Command {
	return mkDirCmd
}

// makeDirCommandFunc executes the "mkdir" command.
func makeDirCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("Key required")
	}
	key := args[0]
	ttl := mkDirTtlFlag

	return client.CreateDir(key, uint64(ttl))
}
