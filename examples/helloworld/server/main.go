package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv.Use(func(next rpc.MiddlewareHandler) rpc.MiddlewareHandler {
		return func(ctx context.Context, rw rpc.ResponseWriter, req *rpc.Request) {
			start := time.Now()
			fmt.Fprintf(os.Stdout, "Execute: %q\n", req.ServiceMethod)
			fmt.Fprintf(os.Stdout, "Client Token: %q\n", req.Header.Get("X-Client-Token"))
			next(ctx, rw, req)
			fmt.Fprintf(os.Stdout, "Executed %q in %v\n", req.ServiceMethod, time.Since(start))
		}
	})

	go func() {
		srv.Accept(context.Background(), ln)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

}
