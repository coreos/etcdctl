etcdctl
========

This is our spec/working area for a simple command line client for etcd. This client will be bundled with CoreOS. This is all brainstorming at the moment. Please contribute!

## Example usage

Setting a key on `/foo/bar`: 

    $ etcdctl -k /foo/bar -v "Hello world"
    Hello world

    $ etcdctl -k /foo/bar -v "Hello world" --ttl 100
    Hello world
    99
 
Getting a key:

    $ etcdctl -k /foo/bar
    Hello world

Deleting a key:

    $ etcdctl -k /foo/bar -d
    Hello world

Atomic Test and Set

    $ etcdctl -k /foo/bar -pv "Hello world" -v "Hello etcd"
    Hello etcd    
    
Tailing a key:
	
	$ etcdctl -k /foo/bar -w
	.... client hangs forever until ctrl+C or the watching value changes 

    $ etcdctl /foo/bar --wf
    Hello world
    .... client hangs forever until ctrl+C printing values as key change

## Return Codes

0	Success

1	Malformed etcdctl arguments

2	Failed to connect to host

3	Failed to auth (client cert rejected, ca validation failure, etc)

4	400 error from etcd

5	500 error from etcd
