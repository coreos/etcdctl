package command

import (
	"fmt"
	"github.com/spf13/cobra"
	// flag "github.com/spf13/pflag"

	"strings"
)

// var for root command
var etcdctlRootCmd *cobra.Command

type peerList []string

// variable for the global flags.
var debugFlag bool
var noSyncFlag bool
var outputFlag string
var peersFlag peerList

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (i *peerList) String() string {
	return fmt.Sprint(*i)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (i *peerList) Set(value string) error {

	//Allow state in peer flag or not : currently allowing state. ( ASK Brandon/others during code review )

	// // If we wanted to allow the flag to be set multiple times,
	// // accumulating values, we would delete this if statement.
	// // That would permit usages such as
	// //	-deltaT 10s -deltaT 15s
	// // and other combinations.
	// if len(*i) > 0 {
	// 	return errors.New("interval flag already set")
	// }
	for _, addr := range strings.Split(value, ",") {
		*i = append(*i, addr)
	}
	return nil
}

func (i *peerList) Type() string {
	return "peerList"
}

func init() {

	etcdctlRootCmd = &cobra.Command{
		Use:   "etcdctl",
		Short: "A simple command line client for etcd.",
	}
	etcdctlRootCmd.Flags().BoolVarP(&debugFlag, "debug", "", false, "output cURL commands which can be used to reproduce the request")
	etcdctlRootCmd.Flags().BoolVarP(&noSyncFlag, "no-sync", "", true, "don't synchronize cluster information before sending request")
	etcdctlRootCmd.PersistentFlags().StringVarP(&outputFlag, "output", "o", "simple", "output response in the given format (`simple` or `json` or `extended`)")
	etcdctlRootCmd.Flags().Var(&peersFlag, "peers", "a comma-delimited list of machine addresses in the cluster (default: \"127.0.0.1:4001\")")
} // end of init

func CreateCommandTree() {

	etcdctlRootCmd.AddCommand(LsCommand())
	etcdctlRootCmd.AddCommand(SetCommand())
	etcdctlRootCmd.AddCommand(SetDirCommand())
	etcdctlRootCmd.AddCommand(MakeCommand())
	etcdctlRootCmd.AddCommand(MakeDirCommand())
	etcdctlRootCmd.AddCommand(RemoveCommand())
	etcdctlRootCmd.AddCommand(RemoveDirCommand())
	etcdctlRootCmd.AddCommand(GetCommand())
	etcdctlRootCmd.Execute()

}
