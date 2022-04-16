package rpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func createChainStreamServer(interceptors ...StreamServerInterceptor) grpc.StreamServerInterceptor {
	var grpc_interceptors []grpc.StreamServerInterceptor
	for _, interceptor := range interceptors {
		grpc_interceptors = append(grpc_interceptors, grpc.StreamServerInterceptor(interceptor))
	}
	return grpc_middleware.ChainStreamServer(grpc_interceptors...)
}

func createChainUnaryServer(interceptors ...UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	var grpc_interceptors []grpc.UnaryServerInterceptor
	for _, interceptor := range interceptors {
		grpc_interceptors = append(grpc_interceptors, grpc.UnaryServerInterceptor(interceptor))
	}
	return grpc_middleware.ChainUnaryServer(grpc_interceptors...)
}
