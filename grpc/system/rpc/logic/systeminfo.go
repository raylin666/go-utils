package logic

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"go-server/grpc/system/rpc/srv"
	"go-server/grpc/system/rpc/system"
	"go-server/internal/utils"
)

type SystemInfoLogic struct {
	ctx    context.Context
	srvCtx *svc.Context
}

func NewSystemInfoLogic(ctx context.Context, srvCtx *svc.Context) *SystemInfoLogic {
	return &SystemInfoLogic{
		ctx:    ctx,
		srvCtx: srvCtx,
	}
}

// 获取系统信息
func (l *SystemInfoLogic) GetSystemInfo(request *system.GetSystemInfoRequest) (*system.GetSystemInfoResponse, error) {
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

	return &system.GetSystemInfoResponse{
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
