package rpc

type contextKey string

var (
	contextRequestID = contextKey("ID")
)
