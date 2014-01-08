package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/coreos/etcdctl/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "etcdctl"
	app.Version = releaseVersion
	app.Usage = "A simple command line client for etcd."
	app.Flags = []cli.Flag{
		cli.BoolFlag{"debug", "output cURL commands which can be used to reproduce the request"},
		cli.StringFlag{"output, o", "simple", "output response in the given format (`simple` or `json`)"},
		cli.StringFlag{"peers, C", "127.0.0.1:4001", "a comma seperated list of machine addresses in the cluster"},
	}
	app.Commands = []cli.Command{
		command.NewMakeCommand(),
		command.NewMakeDirCommand(),
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
