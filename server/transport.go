package server

import (
	"context"
	"net/url"
)

const (
	HTTPServerType = "HTTP"
	GRPCServerType = "gRPC"
)

// Server is transport server.
type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

// Endpointer is registry endpoint.
type Endpointer interface {
	Endpoint() (*url.URL, error)
}

// Header is the storage medium used by a Header.
type Header interface {
	Get(key string) string
	Set(key string, value string)
	Keys() []string
}
