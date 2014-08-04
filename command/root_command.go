package command

import (
	"fmt"
	"github.com/coreos/etcdctl/Godeps/_workspace/src/github.com/coreos/cobra"

	"strings"
)

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
	etcdctlRootCmd.PersistentFlags().BoolVar(&debugFlag, "debug", false, "output cURL commands which can be used to reproduce the request")
	etcdctlRootCmd.PersistentFlags().BoolVar(&noSyncFlag, "no-sync", true, "don't synchronize cluster information before sending request")
	etcdctlRootCmd.PersistentFlags().StringVarP(&outputFlag, "output", "o", "simple", "output response in the given format (`simple` or `json` or `extended`)")
	etcdctlRootCmd.PersistentFlags().VarP(&peersFlag, "peers", "C", "a comma-delimited list of machine addresses in the cluster (default: \"127.0.0.1:4001\")")
}

func CreateCommandTree() {

	etcdctlRootCmd.AddCommand(LsCommand())
	etcdctlRootCmd.AddCommand(SetCommand())
	etcdctlRootCmd.AddCommand(SetDirCommand())
	etcdctlRootCmd.AddCommand(MakeCommand())
	etcdctlRootCmd.AddCommand(MakeDirCommand())
	etcdctlRootCmd.AddCommand(RemoveCommand())
	etcdctlRootCmd.AddCommand(RemoveDirCommand())
	etcdctlRootCmd.AddCommand(GetCommand())
	etcdctlRootCmd.AddCommand(UpdateCommand())
	etcdctlRootCmd.AddCommand(UpdateDirCommand())
	etcdctlRootCmd.AddCommand(WatchCommand())
	etcdctlRootCmd.AddCommand(ExecWatchCommand())
	etcdctlRootCmd.Execute()

}
