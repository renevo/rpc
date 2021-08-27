package rpc

import (
	"context"
	"io"
	"net"
	"time"
)

// DefaultServer is the default instance of *Server.
var DefaultServer = NewServer()

// ServeConn runs the DefaultServer on a single connection.
// ServeConn blocks, serving the connection until the client hangs up.
// The caller typically invokes ServeConn in a go statement.
// ServeConn uses the gob wire format (see package gob) on the
// connection. To use an alternate codec, use ServeCodec.
// See NewClient's comment for information about concurrent access.
func ServeConn(ctx context.Context, conn io.ReadWriteCloser) {
	DefaultServer.ServeConn(ctx, conn)
}

// ServeCodec is like ServeConn but uses the specified codec to
// decode requests and encode responses.
func ServeCodec(ctx context.Context, codec ServerCodec) {
	DefaultServer.ServeCodec(ctx, codec)
}

// ServeRequest is like ServeCodec but synchronously serves a single request.
// It does not close the codec upon completion.
func ServeRequest(ctx context.Context, codec ServerCodec) error {
	return DefaultServer.ServeRequest(ctx, codec)
}

// Accept accepts connections on the listener and serves requests
// to DefaultServer for each incoming connection.
// Accept blocks; the caller typically invokes it in a go statement.
func Accept(ctx context.Context, lis net.Listener) {
	DefaultServer.Accept(ctx, lis)
}

// Dial connects to an RPC server at the specified network address.
func Dial(ctx context.Context, network, address string) (*Client, error) {
	dialer := &net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 10 * time.Second,
	}
	conn, err := dialer.DialContext(ctx, network, address)
	if err != nil {
		return nil, err
	}

	// TODO: Add tcp optimizations here

	return NewClient(conn), nil
}
