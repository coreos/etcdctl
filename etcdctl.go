package main

import (
	"os"

	"github.com/coreos/etcdctl/command"
	"github.com/coreos/etcdctl/third_party/github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "etcdctl"
	app.Version = releaseVersion
	app.Usage = "A simple command line client for etcd."
	app.Flags = []cli.Flag{
		cli.BoolFlag{"debug", "output cURL commands which can be used to reproduce the request"},
		cli.BoolFlag{"no-sync", "don't synchronize cluster information before sending request"},
		cli.StringFlag{"output, o", "simple", "output response in the given format (`simple` or `json`)"},
		cli.StringFlag{"peers, C", "", "a comma-delimited list of machine addresses in the cluster (default: \"127.0.0.1:4001\")"},
	}
	app.Commands = []cli.Command{
		command.NewMakeCommand(),
		command.NewMakeDirCommand(),
		command.NewCopyCommand(),
		command.NewRemoveCommand(),
		command.NewRemoveDirCommand(),
		command.NewGetCommand(),
		command.NewLsCommand(),
		command.NewSetCommand(),
		command.NewSetDirCommand(),
		command.NewUpdateCommand(),
		command.NewUpdateDirCommand(),
		command.NewWatchCommand(),
		command.NewExecWatchCommand(),
	}
	app.Run(os.Args)
}
