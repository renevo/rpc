package rpc

import "context"

// A ServerCodec implements reading of RPC requests and writing of
// RPC responses for the server side of an RPC session.
// The server calls ReadRequestHeader and ReadRequestBody in pairs
// to read requests from the connection, and it calls WriteResponse to
// write a response back. The server calls Close when finished with the
// connection. ReadRequestBody may be called with a nil
// argument to force the body of the request to be read and discarded.
type ServerCodec interface {
	ReadRequestHeader(context.Context, *Request) error
	ReadRequestBody(context.Context, interface{}) error
	WriteResponse(context.Context, *Response, interface{}) error

	// Close can be called multiple times and must be idempotent.
	Close() error
}

// A ClientCodec implements writing of RPC requests and
// reading of RPC responses for the client side of an RPC session.
// The client calls WriteRequest to write a request to the connection
// and calls ReadResponseHeader and ReadResponseBody in pairs
// to read responses. The client calls Close when finished with the
// connection. ReadResponseBody may be called with a nil
// argument to force the body of the response to be read and then
// discarded.
type ClientCodec interface {
	WriteRequest(context.Context, *Request, interface{}) error
	ReadResponseHeader(context.Context, *Response) error
	ReadResponseBody(context.Context, interface{}) error

	// Close can be called multiple times and must be idempotent.
	Close() error
}
