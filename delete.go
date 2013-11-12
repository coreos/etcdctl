package main

import (
	"fmt"
)

const DeleteUsage = `usage: etcdctl [etcd flags] delete <key>`
const DeleteAllUsage = `usage: etcdctl [etcd flags] deleteAll <key>`

func init() {
	registerCommand("delete", DeleteUsage, 1, 1, delete)
	registerCommand("deleteAll", DeleteAllUsage, 1, 1, deleteAll)
}

func delete(args []string) error {
	key := args[0]

	resp, err := client.Delete(key)
	if debug {
		fmt.Println(<-curlChan)
	}
	if err != nil {
		return err
	}
	output(resp)

	return nil
}

func deleteAll(args []string) error {
	key := args[0]

	resp, err := client.DeleteAll(key)
	if debug {
		fmt.Println(<-curlChan)
	}
	if err != nil {
		return err
	}
	output(resp)

	return nil
}
