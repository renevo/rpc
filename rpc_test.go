package rpc

import (
	"context"
	"net"
	"sync"
	"testing"
)

type testRPC int

func (testRPC) Test(ctx context.Context, in string, out *string) error {
	*out = in
	return nil
}

// TestRPC is a super simple functional test of calling over TCP
func TestRPC(t *testing.T) {
	srv := NewServer()
	if err := srv.RegisterName("Test", new(testRPC)); err != nil {
		t.Fatalf("failed to register service: %v", err)
		return
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		srv.Accept(context.Background(), ln)
		wg.Done()
	}()

	client, err := Dial(context.Background(), ln.Addr().Network(), ln.Addr().String())
	if err != nil {
		ln.Close()
		t.Fatalf("failed to dial: %v", err)
		return
	}

	in := "test"
	var out string
	if err := client.Call(context.Background(), "Test.Test", in, &out); err != nil {
		ln.Close()
		t.Fatalf("failed to call Test.Test: %v", err)
		return
	}

	if in != out {
		t.Errorf("Unexpected output: expected %q got %q", in, out)
	}

	ln.Close()
	wg.Wait()
}
