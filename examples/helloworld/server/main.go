package main

import (
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:2311")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	go func() {
		rpc.Accept(ln)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

}
