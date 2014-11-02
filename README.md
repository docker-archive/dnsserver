# dnsserver - a simple DNS service toolkit

This provides a very basic API for programming a DNS service that serves over
UDP. Single A (non-RR) records and SRV records are currently supported,
although this may change in the future.

## Documentation

You can find comprehensive documentation in the source comments or at the URL below:

[http://godoc.org/github.com/docker/dnsserver](http://godoc.org/github.com/docker/dnsserver)

## Usage

There is an example program [here](https://github.com/docker/dnsserver/blob/master/examples/server/server.go).

## Dependencies

[github.com/miekg/dns](https://github.com/miekg/dns)

## License

We use the Apache 2.0 License as with all public Docker projects. You can read
this [here](https://github.com/docker/dnsserver/blob/master/LICENSE).

## Contribution Guidelines

* Fork the project and send a pull request
* Please do not modify any license or version information.
* Thanks!

## Maintainer

Erik Hollensbe <github@hollensbe.org>
