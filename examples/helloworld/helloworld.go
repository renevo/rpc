package helloworld

import (
	"context"
	"fmt"
	"os"

	"github.com/renevo/rpc"
)

// Client is a sample hellowworld client
type Client struct {
	*rpc.Client
}

// Hello to the server!
func (c *Client) Hello(ctx context.Context, name string) (string, error) {
	var msg string
	err := c.Client.Call(ctx, "Server.Hello", name, &msg)
	return msg, err
}

// Server is a sample hellow world server
type Server int

// Hello from the client
func (Server) Hello(ctx context.Context, name string, msg *string) error {
	fmt.Fprintf(os.Stdout, "Hello Request ID: %q\n", rpc.ContextID(ctx))

	*msg = fmt.Sprintf("Hello, %s!", name)
	return nil
}
