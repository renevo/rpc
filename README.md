# RPC Package

[![GoDoc](https://pkg.go.dev/badge/github.com/renevo/rpc)][API]

Based off of *copied from* the `net/rpc` package with the following changes:

* Server and Client `context.Context` implementations.
* Header support
  * Client Header Injection
  * Server Header Inspection and Injection
* Server Middleware
* Requests now have a unique ID rather than a Sequence number

The following features have been removed:

* DialHTTP
* ServeHTTP


## Stability

This API is mostly stable, however it will not be given a 1.x release until it is.
