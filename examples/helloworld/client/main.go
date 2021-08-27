package main

import (
	"context"
	"fmt"
	"os"

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
	msg, err := hc.Hello(context.Background(), "World")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", msg)
}
