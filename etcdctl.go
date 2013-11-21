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
		cli.StringSliceFlag{"C", &cli.StringSlice{"127.0.0.1:4001"}, "a comma seperated list of machine addresses in the cluster"},
	}
	app.Commands = []cli.Command{
		command.NewCreateCommand(),
		command.NewCreateDirCommand(),
		command.NewDeleteCommand(),
		command.NewGetCommand(),
		command.NewSetCommand(),
		command.NewSetDirCommand(),
		command.NewUpdateCommand(),
		command.NewUpdateDirCommand(),
		command.NewWatchCommand(),
	}
	app.Run(os.Args)
}
