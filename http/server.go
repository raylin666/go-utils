package http

import (
	"context"
	"crypto/tls"
	"errors"
	ut "github.com/raylin666/go-utils"
	"github.com/raylin666/go-utils/middleware"
	"github.com/raylin666/go-utils/server"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var _ server.Server = (*Server)(nil)
var _ server.Endpointer = (*Server)(nil)

type ServerOption func(*Server)

// WithServerNetwork with server network.
func WithServerNetwork(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

// WithServerAddress with server address.
func WithServerAddress(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

// WithServerTimeout with server timeout.
func WithServerTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// WithServerEndpoint with server endpoint.
func WithServerEndpoint(endpoint *url.URL) ServerOption {
	return func(o *Server) {
		o.endpoint = endpoint
	}
}

// WithServerTLSConfig with TLS config.
func WithServerTLSConfig(c *tls.Config) ServerOption {
	return func(o *Server) {
		o.tlsConf = c
	}
}

// WithServerHTTPMiddlewares with HTTP middleware option.
func WithServerHTTPMiddlewares(middlewares ...middleware.HTTPHandler) ServerOption {
	return func(o *Server) {
		o.httpMiddlewares = middlewares
	}
}

type Server struct {
	*http.Server

	once            sync.Once
	err             error
	network         string
	address         string
	timeout         time.Duration
	lis             net.Listener
	endpoint        *url.URL
	tlsConf         *tls.Config
	httpMiddlewares []middleware.HTTPHandler
	router 			Router
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}

func NewServer(opts ...ServerOption) *Server {
	var srv = &Server{
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
	}
	for _, opt := range opts {
		opt(srv)
	}

	srv.Server = &http.Server{
		Handler:   middleware.HTTPChain(srv.httpMiddlewares...)(srv),
		TLSConfig: srv.tlsConf,
	}

	return srv
}

// WithRouter 注册路由, 这里只提供路由接口, 可以接入 gin/mux 等路由器
func (s *Server) WithRouter(r Router) {
	s.router = r
}

// Endpoint return a real address to registry endpoint.
// examples:
// 	http://127.0.0.1:8000?isSecure=false
func (s *Server) Endpoint() (*url.URL, error) {
	s.once.Do(func() {
		if s.endpoint != nil {
			return
		}
		lis, err := net.Listen(s.network, s.address)
		if err != nil {
			s.err = err
			return
		}
		addr, err := ut.ExtractAddress(s.address, lis)
		if err != nil {
			lis.Close()
			s.err = err
			return
		}
		s.lis = lis

		s.endpoint = server.NewEndpoint("http", addr, s.tlsConf != nil)
	})
	if s.err != nil {
		return nil, s.err
	}
	return s.endpoint, nil
}

func (s *Server) Start(ctx context.Context) error {
	if _, err := s.Endpoint(); err != nil {
		return err
	}
	s.BaseContext = func(net.Listener) context.Context {
		return ctx
	}

	var err error
	if s.tlsConf != nil {
		err = s.ServeTLS(s.lis, "", "")
	} else {
		err = s.Serve(s.lis)
	}
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
