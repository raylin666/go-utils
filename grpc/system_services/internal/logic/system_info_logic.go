package logic

import (
	"context"
	system_services "go-grpcserver/grpc/system_services/proto"
)

type SystemInfoLogic struct {
	ctx    context.Context
}

func NewSystemInfoLogic(ctx context.Context) *SystemInfoLogic {
	return &SystemInfoLogic{
		ctx: ctx,
	}
}

// 获取系统信息
func (l *SystemInfoLogic) GetSystemInfo(request *system_services.GetSystemInfoRequest) (*system_services.GetSystemInfoResponse, error) {
	return &system_services.GetSystemInfoResponse{
		CpuPercent: 10,
	}, nil
}



