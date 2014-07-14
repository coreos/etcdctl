package main

import (
	"github.com/joshi4/etcdctl/command"
)

func main() {

	command.CreateCommandTree()
	//TODO: Add Flags for top level etcdctl itself.

	// continue in a similar vein adding the different sub commands.
	// The Root command only see's it's sub commands, their flags and sub-sub commands are in the command package files.

	// in raw handle , when looking at the flags for etcdctl - add themabove here or keep etcdctl in the command package as well and have main just run the execute function .

	// 	app := cli.NewApp()
	// 	app.Name = "etcdctl"
	// 	app.Version = releaseVersion
	// 	app.Usage = "A simple command line client for etcd."
	// 	app.Flags = []cli.Flag{
	// 		cli.BoolFlag{"debug", "output cURL commands which can be used to reproduce the request"},
	// cli.BoolFlag{"no-sync", "don't synchronize cluster information before sending request"},
	// cli.StringFlag{"output, o", "simple", "output response in the given format (`simple` or `json`)"},
	// cli.StringSliceFlag{"peers, C", &cli.StringSlice{}, "a comma-delimited list of machine addresses in the cluster (default: \"127.0.0.1:4001\")"},
	// 	}
	// 	app.Commands = []cli.Command{
	// 		command.NewMakeCommand(),
	// 		command.NewMakeDirCommand(),
	// 		command.NewRemoveCommand(),
	// 		command.NewRemoveDirCommand(),
	// 		command.NewGetCommand(),
	// 		command.NewLsCommand(),
	// 		command.NewSetCommand(),
	// 		command.NewSetDirCommand(),
	// 		command.NewUpdateCommand(),
	// 		command.NewUpdateDirCommand(),
	// 		command.NewWatchCommand(),
	// 		command.NewExecWatchCommand(),
	// 	}
	// 	app.Run(os.Args)

}
