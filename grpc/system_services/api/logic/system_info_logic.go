package logic

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	system_services "go-grpcserver/grpc/system_services/proto"
	"go-grpcserver/internal/utils"
)

type SystemInfoLogic struct {
	ctx context.Context
}

func NewSystemInfoLogic(ctx context.Context) *SystemInfoLogic {
	return &SystemInfoLogic{
		ctx: ctx,
	}
}

// 获取系统信息
func (l *SystemInfoLogic) GetSystemInfo(request *system_services.GetSystemInfoRequest) (*system_services.GetSystemInfoResponse, error) {
	var (
		cpuGHz = "0GHz"
	)
	c, _ := cpu.Info()
	for _, v := range c {
		if v.Mhz > 0 {
			cpuGHz = fmt.Sprintf("%v0GHz", v.Mhz/1000)
		}
	}

	cpuCounts, _ := cpu.Counts(true)

	m, _ := mem.VirtualMemory()

	w, _ := disk.Usage("/")

	return &system_services.GetSystemInfoResponse{
		CpuPercent:  utils.GetCpuPercent(),
		MemPercent:  utils.GetMemPercent(),
		DiskPercent: utils.GetDiskPercent(),
		CpuGHz:      cpuGHz,
		CpuCounts:   int32(cpuCounts),
		MemTotal:    utils.FormatFileSize(int64(m.Total)),
		MemUsed:     utils.FormatFileSize(int64(m.Used)),
		DiskTotal:   utils.FormatFileSize(int64(w.Total)),
		DiskUsed:    utils.FormatFileSize(int64(w.Total) - int64(w.Free)),
	}, nil
}
