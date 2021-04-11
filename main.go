package main

import grpc_server "go-grpcserver/grpc/system_services"

func main()  {
	// 启动系统服务
	grpc_server.NewServer()
}
