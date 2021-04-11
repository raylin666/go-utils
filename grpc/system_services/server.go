package grpc_server

import (
	"fmt"
	"go-grpcserver/grpc/system_services/api/server"
	system_services "go-grpcserver/grpc/system_services/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func NewServer()  {
	// 监听本地的 10000 端口
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer() // 创建gRPC服务器
	system_services.RegisterSystemServiceServer(s, &server.SystemService{}) // 在 GRPC 服务端注册服务

	reflection.Register(s) // 在给定的 GRPC 服务器上注册服务器反射服务
	// Serve 方法在 lis 上接受传入连接，为每个连接创建一个 ServerTransport 和 server 的 goroutine。
	// 该 goroutine 读取 GRPC 请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}


