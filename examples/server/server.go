/*
This implements a basic DNS service that has two A records:

test.docker
test2.docker

This also implements two SRV records for 'test' on tcp that point at 80 and 443
for test.docker.

You can supply an argument with an IP:Port pair to listen on, or with no
arguments it listens on 127.0.0.1:53 udp. In that event, it needs to run as
root.

Testing:

$ host test.docker
$ host test2.docker
$ host -t srv _test._tcp.docker
*/

package main

import (
	"net"
	"os"

	dnsserver "github.com/docker/dnsserver"
)

func main() {
	// this is your domain. All records will be scoped under it, e.g.,
	// 'test.docker' below.
	ds := dnsserver.NewDNSServer("docker")

	// set an A record. RRs are not supported.
	ds.SetA("test", net.ParseIP("127.0.0.2"))
	ds.SetA("test2", net.ParseIP("127.0.0.3"))

	// Set a SRV record. The host values will be qualified for you with the
	// domain you provide.
	ds.SetSRV("test", "tcp", []dnsserver.SRVRecord{
		{Port: 80, Host: "test"},
		{Port: 443, Host: "test"},
	})

	if len(os.Args) > 1 {
		if err := ds.Listen(os.Args[1]); err != nil {
			panic(err)
		}
	} else {
		if err := ds.Listen("127.0.0.1:53"); err != nil {
			panic(err)
		}
	}
}
