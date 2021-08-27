# RPC Package

[![Go Reference](https://pkg.go.dev/badge/github.com/renevo/rpc.svg)](https://pkg.go.dev/github.com/renevo/rpc)
[![Go Report Card](https://goreportcard.com/badge/github.com/renevo/rpc)](https://goreportcard.com/report/github.com/renevo/rpc)

Based off of (*copied from*) the `net/rpc` package with the following changes:

```bash
go get -u github.com/renevo/rpc
```

* Server and Client `context.Context` implementations.
* Header support
  * Client Header Injection
  * Server Header Inspection and Injection
* Server Middleware
* Requests now have a unique ID rather than a Sequence number

The following features have been removed:

* DialHTTP
* ServeHTTP
* HandleHTTP

The HTTP functionality could be added to this relatively easily through server.ServeRequest

## Stability

This API is mostly stable, however it will not be given a 1.x release until it is.
