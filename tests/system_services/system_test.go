package system_services

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"go-grpcserver/internal/utils"
	"testing"
)

var ctx context.Context

func TestGetSystemInfo(t *testing.T) {
	//var us server.SystemService
	//res, _ := us.GetSystemInfo(ctx, &system_services.GetSystemInfoRequest{})
	c, _ := cpu.Info()
	for _, v := range c {
		if v.Mhz > 0 {
			fmt.Println(fmt.Sprintf("%v0GHz", v.Mhz/1000))
		}
	}
	fmt.Println(cpu.Counts(true))

	m, _ := mem.VirtualMemory()
	fmt.Println(utils.FormatFileSize(int64(m.Total)))
	fmt.Println(utils.FormatFileSize(int64(m.Used)))

	//d, _ := disk.Partitions(true)
	w, _ := disk.Usage("/")
	fmt.Println(utils.FormatFileSize(int64(w.Total)))
	fmt.Println(utils.FormatFileSize(int64(w.Used)))
	t.Log("success")
}
