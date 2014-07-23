package command

import (
	"errors"
	"github.com/coreos/cobra"
	"os"

	"github.com/coreos/go-etcd/etcd"
)

var setCmd *cobra.Command
var ttlFlag int
var swapWithValueFlag string
var swapWithIndexFlag int

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

// NewSetCommand returns the CLI command for "set".
func SetCommand() *cobra.Command {
	return setCmd
}

// setCommandFunc executes the "set" command.
func setCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("Key required")
	}
	key := args[0]
	value, err := argOrStdin(args, os.Stdin, 1)
	if err != nil {
		return nil, errors.New("Value required")
	}

	ttl := ttlFlag
	prevValue := swapWithValueFlag
	prevIndex := swapWithIndexFlag

	if prevValue == "" && prevIndex == 0 {
		return client.Set(key, value, uint64(ttl))
	} else {
		return client.CompareAndSwap(key, value, uint64(ttl), prevValue, uint64(prevIndex))
	}
}
