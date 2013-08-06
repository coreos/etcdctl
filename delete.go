package main

import (
	"fmt"
)

const DeleteUsage = `usage: etcdctl [etcd flags] delete <key>`

func init() {
	registerCommand("delete", DeleteUsage, 2, 2, delete)
}

func delete(args []string) error {
	key := args[1]

	resp, err := client.Delete(key)
	if err != nil {
		return err
	}
	fmt.Println(resp.PrevValue)

	return nil
}
