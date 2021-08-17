package rpc

// Request is a header written before every RPC call. It is used internally
// but documented here as an aid to debugging, such as when analyzing
// network traffic.
type Request struct {
	ServiceMethod string // format: "Service.Method"
	ID            string // RequestID
	Header        Header // Request headers
}

// Response is a header written before every RPC return. It is used internally
// but documented here as an aid to debugging, such as when analyzing
// network traffic.
type Response struct {
	ServiceMethod string // echoes that of the Request
	ID            string // echoes that of the request
	Header        Header // Response Headers
	Error         string // error, if any.
}
