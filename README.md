etcdctl
========

[![Build Status](https://travis-ci.org/coreos/etcdctl.png)](https://travis-ci.org/coreos/etcdctl)

`etcdctl` is a command line client for [etcd][etcd]. It can be used in scripts or for administrators to explore an etcd cluster.

[etcd]: https://github.com/coreos/etcd

## Getting etcdctl

The latest release is available as a binary at [Github][github-release] along with etcd.

[github-release]: https://github.com/coreos/etcd/releases/

You can also build etcdctl from source:

```
./build
```

## Usage

### Key/Value

Setting a key on `/foo/bar`:

```
etcdctl set /foo/bar "Hello world"
Hello world
```

Getting a key:

```
etcdctl get /foo/bar
Hello world
```

Deleting a key:

```
etcdctl delete /foo/bar
Hello world
```

Tailing a key:

```
etcdctl watch /foo/bar -f
Hello world
.... client hangs forever until ctrl+C printing values as key change
```

### Sets

`etcdctl` implements _sets_ on top of the key-value store that etcd
provides. These are useful for storing lists of unique items. A common
use-case is servers of a particular type that register themselves under
an etcd key, so that they can be detected and used by clients.

Adding members to a set:

```
etcdctl sadd /queues amqp://user:password@rabbitmq1
amqp://user:password@rabbitmq1
etcdctl sadd /queues amqp://user:password@rabbitmq2 --ttl=60
amqp://user:password@rabbitmq2
```
    
List all members:

```
etcdctl smembers /queues
amqp://user:password@rabbitmq1
amqp://user:password@rabbitmq2
```

To delete a member:

```
etcdctl sdel /queues amqp://user:password@rabbitmq1
```

## Return Codes

0	Success

1	Malformed etcdctl arguments

2	Failed to connect to host

3	Failed to auth (client cert rejected, ca validation failure, etc)

4	400 error from etcd

5	500 error from etcd

## Project Details

### Versioning

etcdctl uses [semantic versioning][semver].
Releases will follow lockstep with the etcd release cycle.

[semver]: http://semver.org/

### License

etcdctl is under the Apache 2.0 license. See the [LICENSE][license] file for details.

[license]: https://github.com/coreos/etcdctl/blob/master/LICENSE
