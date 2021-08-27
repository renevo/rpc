package helloworld

import (
	"context"
	"fmt"

	"github.com/renevo/rpc"
)

type Client struct {
	*rpc.Client
}

func (c *Client) Hello(ctx context.Context, name string) (string, error) {
	var msg string
	err := c.Client.Call(ctx, "Server.Hello", name, &msg)
	return msg, err
}

type Server struct {
}

func (s *Server) Hello(ctx context.Context, name string, msg *string) error {
	*msg = fmt.Sprintf("Hello, %s!", name)
	return nil
}
