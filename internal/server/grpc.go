package server

import (
	"github.com/raylin666/go-gin-api/pkg/grpc"
	"go-server/config"
	auth_pb "go-server/grpc/auth/rpc/auth"
	auth_server "go-server/grpc/auth/rpc/server"
	auth_svc "go-server/grpc/auth/rpc/svc"
	system_server "go-server/grpc/system/rpc/server"
	system_svc "go-server/grpc/system/rpc/svc"
	system_pb "go-server/grpc/system/rpc/system"
	go_grpc "google.golang.org/grpc"
)

// 创建 GRPC 服务
func NewGrpcServer() {
	// 创建 gRPC 系统服务
	grpc.NewServer(grpc.Server{
		Network: config.Get().Grpc.System.Network,
		Host:    config.Get().Grpc.System.Host,
		Port:    config.Get().Grpc.System.Port,
		RegisterServer: func(g *go_grpc.Server) {
			ctx := system_svc.NewContext()
			srv := system_server.NewSystemServer(ctx)
			system_pb.RegisterSystemServer(g, srv)
		},
	})

	// 创建 gRPC 鉴权服务
	grpc.NewServer(grpc.Server{
		Network: config.Get().Grpc.Auth.Network,
		Host:    config.Get().Grpc.Auth.Host,
		Port:    config.Get().Grpc.Auth.Port,
		RegisterServer: func(g *go_grpc.Server) {
			ctx := auth_svc.NewContext()
			srv := auth_server.NewAuthServer(ctx)
			auth_pb.RegisterAuthServer(g, srv)
		},
	})
}
