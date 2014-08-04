package command

import (
	"errors"
	"os"

	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/cobra"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
)

var (
	setCmd            *cobra.Command
	ttlFlag           int
	swapWithValueFlag string
	swapWithIndexFlag int
)

func init() {
	setCmd = &cobra.Command{
		Use:   "set",
		Short: "set the value of a key.",
		Run: func(cmd *cobra.Command, args []string) {
			handleKey(cmd, args, setCommandFunc)
		},
	}

	setCmd.Flags().IntVar(&ttlFlag, "ttl", 0, "key time-to-live ")
	setCmd.Flags().StringVar(&swapWithValueFlag, "swap-with-value", "", "previous value")
	setCmd.Flags().IntVar(&swapWithIndexFlag, "swap-with-index", 0, "previous index")
}

// NewSetCommand returns the cobra command for "set".
func SetCommand() *cobra.Command {
	return setCmd
}

// setCommandFunc executes the "set" command.
func setCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("key required")
	}
	key := args[0]
	value, err := argOrStdin(args, os.Stdin, 1)
	if err != nil {
		return nil, errors.New("value required")
	}
	ttl := uint64(ttlFlag)
	prevIndex := uint64(swapWithIndexFlag)

	if swapWithValueFlag == "" && prevIndex == 0 {
		return client.Set(key, value, ttl)
	} else {
		return client.CompareAndSwap(key, value, ttl, swapWithValueFlag, prevIndex)
	}
}
