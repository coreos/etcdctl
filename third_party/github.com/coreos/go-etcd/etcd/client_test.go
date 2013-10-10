package etcd

import (
	"fmt"
	"testing"
	"net/url"
	"net"
)

// To pass this test, we need to create a cluster of 3 machines
// The server should be listening on 127.0.0.1:4001, 4002, 4003
func TestSync(t *testing.T) {
	fmt.Println("Make sure there are three nodes at 0.0.0.0:4001-4003")

	c := NewClient(nil)

	success := c.SyncCluster()
	if !success {
		t.Fatal("cannot sync machines")
	}

	for _, m := range(c.GetCluster()) {
		u, err := url.Parse(m)
		if err != nil {
			t.Fatal(err)
		}
		if u.Scheme != "http" {
			t.Fatal("scheme must be http")
		}
		
		host, _, err := net.SplitHostPort(u.Host)
		if err != nil {
			t.Fatal(err)
		}
		if host != "127.0.0.1" {
			t.Fatal("Host must be 127.0.0.1")
		}
	}

	badMachines := []string{"abc", "edef"}

	success = c.SetCluster(badMachines)

	if success {
		t.Fatal("should not sync on bad machines")
	}

	goodMachines := []string{"127.0.0.1:4002"}

	success = c.SetCluster(goodMachines)

	if !success {
		t.Fatal("cannot sync machines")
	} else {
		fmt.Println(c.cluster.Machines)
	}

}
