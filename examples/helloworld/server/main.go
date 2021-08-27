package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/renevo/rpc"
	"github.com/renevo/rpc/examples/helloworld"
)

func main() {
	srv := rpc.NewServer()
	srv.Register(&helloworld.Server{})

	ln, err := net.Listen("tcp", "0.0.0.0:2311")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	go func() {
		srv.Accept(context.Background(), ln)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

}
