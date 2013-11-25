#!/bin/bash -e

# Simple example of using the watch command to execute a script on the
# change of an etcd key.

# The key to watch
KEY=config

# Use the IFS to parse the response and split on newline
IFS=$'\n'
set $(./etcdctl -o extended get ${KEY})

# Setup the initial index to run everything once
nextindex=$2

while true; do
	set $(./etcdctl -o extended watch ${KEY} --index $nextindex)

	config=$1
	nextindex=`expr $2 + 1`

	# Print out the example command line to execute
	echo "echo ${config} | config-update"
done

