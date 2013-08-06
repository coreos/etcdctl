package main

import (
	"errors"
)

type Command struct {
	usage   string
	minArgs int
	maxArgs int
	f       func(args []string) error
}

var commands = make(map[string]*Command)

func registerCommand(name string, usage string, minArgs int, maxArgs int, f func(args []string) error) error {
	_, ok := commands[name]
	if ok {
		return errors.New("duplicate command name")
	}

	commands[name] = &Command{
		usage:   usage,
		minArgs: minArgs,
		maxArgs: maxArgs,
		f:       f,
	}
	return nil
}
