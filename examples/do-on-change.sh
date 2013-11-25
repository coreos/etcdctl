#!/bin/bash -e

# Simple example of using the watch command to execute a script on the
# change of an etcd key.

# The key to watch
KEY=config

# Use the IFS to parse the response and split on newline
IFS=$'\n'

# Setup the initial index to run everything once
set $(./etcdctl -o extended get ${KEY})
index=$2

while true; do
	set $(./etcdctl -o extended watch ${KEY} --after-index $index)

	config=$1
	index=$2

	# Print out the example command line to execute
	echo "echo ${config} | config-update"
done

