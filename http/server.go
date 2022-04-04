package http

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/gorilla/mux"
	"github.com/raylin666/go-utils/middleware"
	"github.com/raylin666/go-utils/server"
	"github.com/raylin666/go-utils/server/encoder"
	"github.com/raylin666/go-utils/server/protocol"
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

// WithServerMiddleware with service middleware option.
func WithServerMiddleware(m ...middleware.Middleware) ServerOption {
	return func(o *Server) {
		o.middlewares = m
	}
}

// WithServerHTTPMiddlewares with HTTP middleware option.
func WithServerHTTPMiddlewares(middlewares ...HTTPMiddlewareHandler) ServerOption {
	return func(o *Server) {
		o.httpMiddlewares = middlewares
	}
}

// WithServerRequestDecoder with request decoder.
func WithServerRequestDecoder(dec encoder.DecodeRequestFunc) ServerOption {
	return func(o *Server) {
		o.dec = dec
	}
}

// WithServerResponseEncoder with response encoder.
func WithServerResponseEncoder(en encoder.EncodeResponseFunc) ServerOption {
	return func(o *Server) {
		o.enc = en
	}
}

// WithServerErrorEncoder with error encoder.
func WithServerErrorEncoder(en encoder.EncodeErrorFunc) ServerOption {
	return func(o *Server) {
		o.ene = en
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
	middlewares     []middleware.Middleware
	httpMiddlewares []HTTPMiddlewareHandler
	router          *mux.Router
	dec             encoder.DecodeRequestFunc
	enc             encoder.EncodeResponseFunc
	ene             encoder.EncodeErrorFunc
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
		Handler:   HTTPMiddlewareFilterChain(srv.httpMiddlewares...)(srv),
		TLSConfig: srv.tlsConf,
	}

	srv.router = mux.NewRouter()
	srv.router.Use(srv.middleware())
	return srv
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
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
		addr, err := protocol.ExtractAddress(s.address, lis)
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

func (s *Server) middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx, cancel := context.WithCancel(req.Context())
			defer cancel()
			if s.timeout > 0 {
				ctx, cancel = context.WithTimeout(ctx, s.timeout)
				defer cancel()
			}
			pathTemplate := req.URL.Path
			if route := mux.CurrentRoute(req); route != nil {
				// /path/123 -> /path/{id}
				pathTemplate, _ = route.GetPathTemplate()
			}
			var trans = new(server.Transport)
			trans.Option.Endpoint = s.endpoint.String()
			trans.Option.Operation = pathTemplate
			trans.Option.ReqHeader = server.HeaderCarrier(req.Header)
			trans.Option.ReplyHeader = server.HeaderCarrier(w.Header())
			trans.Option.Request = req
			trans.Option.PathTemplate = pathTemplate
			ctx = server.NewServerContext(ctx, trans)
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
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
