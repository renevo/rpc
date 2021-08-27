package rpc

import (
	"context"
	"testing"
)

type Tester int

func (t Tester) Test(ctx context.Context, arg int, reply *int) error {
	*reply = 5

	return nil
}

func TestServerRegister(t *testing.T) {
	s := &Server{}
	if err := s.Register(new(Tester)); err != nil {
		t.Fatalf("Server Failed to Register: %v", err)
	}

	s.serviceMap.Range(func(k, v interface{}) bool {
		t.Logf("Service %q", k)
		svc := v.(*service)

		for name, method := range svc.method {
			t.Logf("Method: %q (%s,%s)", name, method.ArgType, method.ReplyType)
		}

		return true
	})
}
