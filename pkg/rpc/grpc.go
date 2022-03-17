package rpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

var _ GrpcServer = (*grpcServer)(nil)

type GrpcServer interface {
	RegisterService(f func(s grpc.ServiceRegistrar, srv interface{}), srv interface{})
	Serve(lis net.Listener) error
	Stop()
}

type grpcServer struct {
	*grpc.Server
}

type ServerOption grpc.ServerOption

func WithGRPCStreamInterceptor(interceptors ...StreamServerInterceptor) ServerOption {
	var i = make([]grpc.StreamServerInterceptor, len(interceptors))
	for k, irt := range interceptors {
		i[k] = grpc.StreamServerInterceptor(irt)
	}
	return grpc.StreamInterceptor(createChainStreamServer(i...))
}

func WithGRPCUnaryInterceptor(interceptors ...UnaryServerInterceptor) ServerOption {
	var i = make([]grpc.UnaryServerInterceptor, len(interceptors))
	for k, irt := range interceptors {
		i[k] = grpc.UnaryServerInterceptor(irt)
	}
	return grpc.UnaryInterceptor(createChainUnaryServer(i...))
}

func NewGRPCServer(opts ...ServerOption) GrpcServer {
	var s = new(grpcServer)
	var o = make([]grpc.ServerOption, len(opts))
	for k, sopt := range opts {
		o[k] = sopt.(grpc.ServerOption)
	}
	s.Server = grpc.NewServer(o...)
	return s
}

func GRPCDialContext(address string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	if len(opts) <= 0 {
		opts = append(opts, grpc.WithBlock())
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return grpc.DialContext(context.Background(), address, opts...)
}

func (s *grpcServer) RegisterService(f func(s grpc.ServiceRegistrar, srv interface{}), srv interface{}) {
	f(s.Server, srv)
}

func (s *grpcServer) Serve(lis net.Listener) error {
	return s.Server.Serve(lis)
}

func (s *grpcServer) Stop() {
	s.Server.Stop()
}
