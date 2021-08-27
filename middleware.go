package rpc

import "context"

// MiddlewareHandler used for calling Middleware
type MiddlewareHandler func(context.Context, ResponseWriter, *Request)

// middleware interface is anything which implements a MiddlewareFunc named Middleware.
type middleware interface {
	Middleware(MiddlewareHandler) MiddlewareHandler
}

// MiddlewareFunc is a function which receives a MiddlewareHandler and returns another MiddlewareHandler.
// Typically, the returned handler is a closure which does something with the Response and Request passed
// to it, and then calls the handler passed as parameter to the MiddlewareFunc.
type MiddlewareFunc func(MiddlewareHandler) MiddlewareHandler

// Middleware allows MiddlewareFunc to implement the middleware interface.
func (mwf MiddlewareFunc) Middleware(h MiddlewareHandler) MiddlewareHandler {
	return mwf(h)
}

// Use appends a MiddlewareFunc to the chain. Middleware can be used to intercept or otherwise modify requests and/or responses, and are executed in the order that they are applied to the Server.
func (s *Server) Use(mwf ...MiddlewareFunc) {
	for _, fn := range mwf {
		s.middlewares = append(s.middlewares, fn)
	}
}

// ResponseWriter provides an interface for middleware to write response headers and errors
type ResponseWriter interface {
	Header() Header
	WriteError(err error)
	Err() error
}

type responseWriter struct {
	*Response
}

func (r responseWriter) Header() Header {
	return r.Response.Header
}

func (r responseWriter) WriteError(err error) {
	if err == nil {
		return
	}

	r.Response.callErr = err
}

func (r responseWriter) Err() error {
	return r.Response.callErr
}
