package command

import (
	"errors"
	"os"

	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/cobra"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
)

var (
	mkCmd     *cobra.Command
	mkTtlFlag int
)

func init() {
	mkCmd = &cobra.Command{
		Use:   "mk",
		Short: "make a new key with a given value",
		Run: func(cmd *cobra.Command, args []string) {
			handleKey(cmd, args, makeCommandFunc)
		},
	}
	mkCmd.Flags().IntVar(&mkTtlFlag, "ttl", 0, "key time-to-live")
}

// returns the mkCommand.
func MakeCommand() *cobra.Command {
	return mkCmd
}

// makeCommandFunc executes the "make" command.
func makeCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("key required")
	}
	key := args[0]
	value, err := argOrStdin(args, os.Stdin, 1)
	if err != nil {
		return nil, errors.New("value required")
	}
	return client.Create(key, value, uint64(mkTtlFlag))
}
