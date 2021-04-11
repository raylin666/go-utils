package server

import (
	"context"
	"go-grpcserver/grpc/system_services/internal/logic"
	system_services "go-grpcserver/grpc/system_services/proto"
)

// 系统信息服务
type SystemService struct {}

// 获取系统信息
func (server *SystemService) GetSystemInfo(ctx context.Context, request *system_services.GetSystemInfoRequest) (*system_services.GetSystemInfoResponse, error) {
	l := logic.NewSystemInfoLogic(ctx)
	return l.GetSystemInfo(request)
}

