package rpc

import "context"

type contextKey string

var (
	contextID            = contextKey("ID")
	contextHeader        = contextKey("Header")
	contextServiceMethod = contextKey("ServiceMethod")
)

// ContextHeader returns the rpc.Header from the supplied context.Context
func ContextHeader(ctx context.Context) Header {
	v := ctx.Value(contextHeader)
	if h, ok := v.(Header); ok {
		return h
	}

	return Header{}
}

// ContextID returns the rpc request ID from the supplied context.Context
func ContextID(ctx context.Context) string {
	v := ctx.Value(contextID)
	if id, ok := v.(string); ok {
		return id
	}

	return ""
}

// ContextServiceMethod returns the service method name fromt he supplied context.Context
func ContextServiceMethod(ctx context.Context) string {
	v := ctx.Value(contextServiceMethod)
	if sm, ok := v.(string); ok {
		return sm
	}

	return ""
}

// ContextWithHeaders will set the current headers for a client context. This has no effect on a server.
func ContextWithHeaders(ctx context.Context, h Header) context.Context {
	return context.WithValue(ctx, contextHeader, h)
}
