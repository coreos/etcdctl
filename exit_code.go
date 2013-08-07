package main

const (
	SUCCESS = iota
	MalformedEtcdctlArguments
	FailedToConnectToHost
	FailedToAuth
	ErrorFromEtcd
)
