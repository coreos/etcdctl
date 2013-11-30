#!/bin/bash -e

# Simple example of using the watch command to execute a script on the
# change of an etcd key.

# The key to watch
KEY=config

index_re='s/^Modified-Index: \(.*\)/\1/p'

index=$(./etcdctl -o extended get ${KEY} | sed -n -e "${index_re}")

while true; do
	out=$(./etcdctl -o extended watch ${KEY} --after-index ${index})
	index=$(echo "${out}" | sed -n -e "${index_re}")
	config=$(echo "${out}" | sed -e '1,/^$/d')

	# Print out the example command line to execute
	echo "echo \"${config}\" | config-update"
done

