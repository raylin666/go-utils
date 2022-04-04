package middleware

import (
	"context"
	"net/http"
)

// Handler defines the handler invoked by Middleware.
type Handler func(ctx context.Context, req interface{}) (interface{}, error)

// Middleware is HTTP/gRPC transport middleware.
type Middleware func(Handler) Handler

// Chain returns a Middleware that specifies the chained handler for endpoint.
func Chain(m ...Middleware) Middleware {
	return func(next Handler) Handler {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next)
		}
		return next
	}
}

// HTTPHandler is a function which receives an http.Handler and returns another http.Handler.
type HTTPHandler func(http.Handler) http.Handler

// HTTPChain returns a HTTPHandler that specifies the chained handler for HTTP Router.
func HTTPChain(m ...HTTPHandler) HTTPHandler {
	return func(next http.Handler) http.Handler {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next)
		}
		return next
	}
}