package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/renevo/rpc"
	"github.com/renevo/rpc/examples/helloworld"
)

func main() {
	client, err := rpc.Dial(context.Background(), "tcp", "localhost:2311")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	hc := &helloworld.Client{Client: client}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	msg, err := hc.Hello(rpc.ContextWithHeaders(ctx, rpc.Header{}.Set("X-Client-Token", "MY-TOKEN-HERE")), "World")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", msg)
}
