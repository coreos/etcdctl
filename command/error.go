package command

const (
	SUCCESS = iota
	MalformedEtcdctlArguments
	FailedToConnectToHost
	FailedToAuth
	ErrorFromEtcd
)
