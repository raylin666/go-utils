package rpc

import (
	"context"
	"github.com/raylin666/go-utils/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"sync"
)

var _ server.Server = (*GRPCServer)(nil)

type GRPCServerOption func(*GRPCServer)

// WithGRPCServerNetwork with server network.
func WithGRPCServerNetwork(network string) GRPCServerOption {
	return func(grpcServer *GRPCServer) {
		grpcServer.network = network
	}
}

// WithGRPCServerAddress with server address.
func WithGRPCServerAddress(address string) GRPCServerOption {
	return func(grpcServer *GRPCServer) {
		grpcServer.address = address
	}
}

func WithGRPCStreamInterceptor(interceptors ...StreamServerInterceptor) GRPCServerOption {
	return func(grpcServer *GRPCServer) {
		grpcServer.StreamServerInterceptor = append(grpcServer.StreamServerInterceptor, interceptors...)
	}
}

func WithGRPCUnaryInterceptor(interceptors ...UnaryServerInterceptor) GRPCServerOption {
	return func(grpcServer *GRPCServer) {
		grpcServer.UnaryServerInterceptor = append(grpcServer.UnaryServerInterceptor, interceptors...)
	}
}

type GRPCServer struct {
	*grpc.Server

	once                    sync.Once
	network                 string
	address                 string
	err                     error
	lis                     net.Listener
	StreamServerInterceptor []StreamServerInterceptor
	UnaryServerInterceptor  []UnaryServerInterceptor
}

func NewGRPCServer(opts ...GRPCServerOption) *GRPCServer {
	var srv = &GRPCServer{
		network: "tcp",
	}
	for _, opt := range opts {
		opt(srv)
	}

	var serverOption []grpc.ServerOption
	if len(srv.StreamServerInterceptor) > 0 {
		serverOption = append(serverOption, grpc.StreamInterceptor(createChainStreamServer(srv.StreamServerInterceptor...)))
	}
	if len(srv.UnaryServerInterceptor) > 0 {
		serverOption = append(serverOption, grpc.UnaryInterceptor(createChainUnaryServer(srv.UnaryServerInterceptor...)))
	}

	srv.Server = grpc.NewServer(serverOption...)

	return srv
}

// GRPCDialContext 创建到给定目标的客户端连接
func GRPCDialContext(ctx context.Context, address string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	if len(opts) <= 0 {
		opts = append(opts, grpc.WithBlock())
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return grpc.DialContext(ctx, address, opts...)
}

// Preset 预加载服务数据参数
func (s *GRPCServer) Preset() error {
	s.once.Do(func() {
		lis, err := net.Listen(s.network, s.address)
		if err != nil {
			s.err = err
			return
		}
		s.lis = lis
	})

	if s.err != nil {
		return s.err
	}

	return nil
}

func (s *GRPCServer) Start(ctx context.Context) error {
	err := s.Preset()
	if err != nil {
		return err
	}

	err = s.Server.Serve(s.lis)
	if err != nil {
		return err
	}

	return nil
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	s.Server.Stop()
	return nil
}
