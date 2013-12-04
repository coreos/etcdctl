package command

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

// NewExecWatchCommand returns the CLI command for "exec-watch".
func NewExecWatchCommand() cli.Command {
	return cli.Command{
		Name:  "exec-watch",
		Usage: "watch a key for changes and exec an executable",
		Flags: []cli.Flag{
			cli.IntFlag{"after-index", 0, "watch after the given index"},
		},
		Action: func(c *cli.Context) {
			handle(c, execWatchCommandFunc)
		},
	}
}

// execWatchCommandFunc executes the "exec-watch" command.
func execWatchCommandFunc(c *cli.Context, client *etcd.Client) (*etcd.Response, error) {
	_ = io.Copy
	_ = exec.Command
	args := c.Args()
	argsLen := len(args)

	if argsLen < 2 {
		return nil, errors.New("Key and command to exec required")
	}
	key := args[0]

	index := 0
	if c.Int("after-index") != 0 {
		index = c.Int("after-index") + 1
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	stop := make(chan bool)

	go func() {
		<-sigch
		stop <- true
		os.Exit(0)
	}()

	receiver := make(chan *etcd.Response)
	client.SetConsistency(etcd.WEAK_CONSISTENCY)
	go client.Watch(key, uint64(index), false, receiver, stop)

	cmdOffset := 0
	for i, v := range(args) {
		if v == "--" {
			cmdOffset = i + 1
		}
	}

	for {
		resp := <-receiver
		cmd := exec.Command(args[cmdOffset], args[cmdOffset+1:]...)
		cmd.Env = environResponse(resp, os.Environ())

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = cmd.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		go io.Copy(os.Stdout, stdout)
		go io.Copy(os.Stderr, stderr)
		cmd.Wait()
	}

	return nil, nil
}

func environResponse(resp *etcd.Response, env []string) []string {
	env = append(env, "ETCD_WATCH_MODIFIED_INDEX="+fmt.Sprintf("%d", resp.Node.ModifiedIndex))
	env = append(env, "ETCD_WATCH_KEY="+resp.Node.Key)
	env = append(env, "ETCD_WATCH_VALUE="+resp.Node.Value)
	return env
}
