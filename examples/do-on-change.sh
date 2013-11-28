#!/bin/bash -e

# Simple example of using the watch command to execute a script on the
# change of an etcd key.

# The key to watch
KEY=config

out=$(./etcdctl -o extended get ${KEY} | tail -n 1)
index=${out##*# index:}

while true; do
	out=$(./etcdctl -o extended watch ${KEY} --after-index $index)
	config=${out%?# index:*}
	index=${out##*# index:}

	# Print out the example command line to execute
	echo "echo ${config} | config-update"
done

