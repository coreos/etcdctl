package command

import (
	"errors"
	"os"
	"os/signal"

	"github.com/coreos/go-etcd/etcd"
	"github.com/joshi4/cobra"
)

var watchCmd *cobra.Command

//flags
var foreverFlag bool
var watchRecursiveFlag bool
var watchAfterIndexFlag int

func init() {
	watchCmd = &cobra.Command{
		Use:   "watch",
		Short: "watch a key for changes",
		Run: func(cmd *cobra.Command, args []string) {
			handleKey(cmd, args, watchCommandFunc)
		},
	}
	watchCmd.Flags().BoolVar(&foreverFlag, "forever", false, "forever watch a key unitl CTRL+C")
	watchCmd.Flags().BoolVar(&watchRecursiveFlag, "recursive", false, "returns all values for key and child keys")
	watchCmd.Flags().IntVar(&watchAfterIndexFlag, "after-index", 0, "watch after the given index")

}

// WatchCommand returns the CLI command for "watch".
func WatchCommand() *cobra.Command {
	return watchCmd

}

// watchCommandFunc executes the "watch" command.
func watchCommandFunc(cmd *cobra.Command, args []string, client *etcd.Client) (*etcd.Response, error) {
	if len(args) == 0 {
		return nil, errors.New("Key required")
	}
	key := args[0]
	recursive := watchRecursiveFlag
	forever := foreverFlag

	index := 0
	if watchAfterIndexFlag != 0 {
		index = watchAfterIndexFlag + 1
	}

	if forever {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt)
		stop := make(chan bool)

		go func() {
			<-sigch
			os.Exit(0)
		}()

		receiver := make(chan *etcd.Response)
		errCh := make(chan error, 1)

		go func() {
			_, err := client.Watch(key, uint64(index), recursive, receiver, stop)
			errCh <- err
		}()

		for {
			select {
			case resp := <-receiver:
				//outputFlag is a flag(Persistent) var defined in root_command.go.
				printAll(resp, outputFlag)
			case err := <-errCh:
				handleError(-1, err)
			}
		}

	} else {
		var resp *etcd.Response
		var err error
		resp, err = client.Watch(key, uint64(index), recursive, nil, nil)

		if err != nil {
			handleError(ErrorFromEtcd, err)
		}

		if err != nil {
			return nil, err
		}
		printAll(resp, outputFlag)
	}

	return nil, nil
}
