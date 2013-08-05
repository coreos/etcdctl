etcdctl
========

This is our spec/working area for a simple command line client for etcd. This client will be bundled with CoreOS. This is all brainstorming at the moment. Please contribute!

## Example usage

Setting a key on `/foo/bar`: 

    $ etcdctl /foo/bar "Hello world"
    Hello world
    
Getting a key:

    $ etcdctl /foo/bar
    Hello world
    
Tailing a key:

    $ etcdctl /foo/bar -f
    Hello world
    .... client hangs forever until ctrl+C printing values as key change

## Return Codes

0	Success

1	Malformed etcdctl arguments

2	Failed to connect to host

3	Failed to auth (client cert rejected, ca validation failure, etc)

4	400 error from etcd

5	500 error from etcd
