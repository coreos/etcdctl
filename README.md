etcd-cli
========

This is our spec/working area for a simple command line client for etcd. This client will be bundled with CoreOS. This is all brainstorming at the moment. Please contribute!

## Example usage

Setting a key on `/foo/bar`: 

    $ etcd-cli /foo/bar "Hello world"
    Hello world
    
Getting a key:

    $ etcd-cli /foo/bar
    Hello world
    
Tailing a key:

    $ etcd-cli /foo/bar -f
    Hello world
    .... client hangs forever until ctrl+C printing values as key change
