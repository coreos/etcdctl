package main

import (
	"strings"
)

type ClusterValue struct {
	machines string
}

func (c *ClusterValue) String() string {
	return c.machines
}

func (c *ClusterValue) Set(value string) error {
	if len(value) == 0 {
		return nil
	}

	c.machines = value

	return nil
}

func (c *ClusterValue) GetMachines() []string {
	return strings.Split(c.machines, ",")
}
