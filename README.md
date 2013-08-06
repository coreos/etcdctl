etcdctl
========

This is our spec/working area for a simple command line client for etcd. This client will be bundled with CoreOS. This is all brainstorming at the moment. Please contribute!

## Example usage

Setting a key on `/foo/bar`: 

    $ etcdctl /foo/bar -v "Hello world"
    Hello world

    $ etcdctl /foo/bar -v "Hello world" --ttl 100
    Hello world
    99
 
Getting a key:

    $ etcdctl /foo/bar
    Hello world

Deleting a key:

    $ etcdctl /foo/bar -d
    Hello world

Atomic Test and Set

    $ etcdctl /foo/bar --pv "Hello world" -v "Hello etcd"
    Hello etcd    
    
Tailing a key:
	
	$ etcdctl /foo/bar -w
	.... client hangs forever until ctrl+C or the watching value changes 

    $ etcdctl /foo/bar --wf
    Hello world
    .... client hangs forever until ctrl+C printing values as key change

## Flags 

-c      a list of machines in one cluster

-d      delete a key

-v      the value to set
--pv    the previous to test against

-w      watch change of a key
--wf    keep on watching changes of a key 
--index the index to watch from

--cert  client certificate
--key   client private key 

--detail detailed return information 

## Environment Variables 

ETCD_CLUSTER - The etcd cluster to join to; overridden by -c.
ETCD_KEY     - The client key path; overridden by --key
ETCD_CERT    - The client certificate path; overridden by --cert


## Return Codes

0	Success

1	Malformed etcdctl arguments

2	Failed to connect to host

3	Failed to auth (client cert rejected, ca validation failure, etc)

4	400 error from etcd

5	500 error from etcd
