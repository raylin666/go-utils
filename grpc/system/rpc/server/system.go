package server

import (
	"context"
	"go-server/grpc/system/rpc/logic"
	"go-server/grpc/system/rpc/system"
)

// 系统服务
type System struct {}

// 获取系统信息
func (server *System) GetSystemInfo(ctx context.Context, request *system.GetSystemInfoRequest) (*system.GetSystemInfoResponse, error) {
	l := logic.NewSystemInfoLogic(ctx)
	return l.GetSystemInfo(request)
}

