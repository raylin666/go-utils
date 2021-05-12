package server

import (
	"context"
	"go-server/grpc/system/rpc/logic"
	"go-server/grpc/system/rpc/srv"
	"go-server/grpc/system/rpc/system"
)

// 系统服务
type SystemServer struct {
	srvCtx *svc.Context
}

func NewSystemServer(ctx *svc.Context) *SystemServer {
	return &SystemServer{
		srvCtx: ctx,
	}
}

// 获取系统信息
func (server *SystemServer) GetSystemInfo(ctx context.Context, request *system.GetSystemInfoRequest) (*system.GetSystemInfoResponse, error) {
	l := logic.NewSystemInfoLogic(ctx, server.srvCtx)
	return l.GetSystemInfo(request)
}

