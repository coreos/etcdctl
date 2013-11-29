#!/bin/bash -e

# Simple example of using the watch command to execute a script on the
# change of an etcd key.

# The key to watch
KEY=config

index_re='s/^# modified-index: \(.*\)/\1/p'

index=$(./etcdctl -o extended get ${KEY} | sed -n -e "${index_re}")

while true; do
	out=$(./etcdctl -o extended watch ${KEY} --after-index ${index})
	config=$(echo "${out}" | grep -v "^#")
	index=$(echo "${out}" | sed -n -e "${index_re}")

	# Print out the example command line to execute
	echo "echo \"${config}\" | config-update"
done

