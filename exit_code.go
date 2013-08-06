package main

const (
	SUCCESS = iota
	MalformedEtcdctlArguments
	FailedToConnectToHost
	FailedToAuth
	Error400FromEtcd
	Error500FromEtcd
)
