package http

import "net/http"

// HTTPMiddlewareHandler is a function which receives an http.Handler and returns another http.Handler.
type HTTPMiddlewareHandler func(http.Handler) http.Handler

// HTTPMiddlewareFilterChain returns a HTTPMiddlewareHandler that specifies the chained handler for HTTP Router.
func HTTPMiddlewareFilterChain(filters ...HTTPMiddlewareHandler) HTTPMiddlewareHandler {
	return func(next http.Handler) http.Handler {
		for i := len(filters) - 1; i >= 0; i-- {
			next = filters[i](next)
		}
		return next
	}
}
