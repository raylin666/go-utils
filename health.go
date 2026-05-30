// Package ut 提供通用健康检查接口定义
// 使用场景：
//   - 服务健康监控
//   - 服务发现与注册
//   - 负载均衡健康检查
package ut

import (
	"context"
	"time"
)

// HealthChecker 定义健康检查接口
// 所有连接类模块应实现此接口，用于服务监控和故障发现
type HealthChecker interface {
	// HealthCheck 执行健康检查
	// 参数：
	//   - ctx: 上下文，用于超时控制
	// 返回值：
	//   - nil: 服务健康
	//   - error: 服务异常时的错误信息
	// 使用示例：
	//   err := client.HealthCheck(ctx)
	//   if err != nil {
	//       log.Printf("服务异常: %v", err)
	//   }
	HealthCheck(ctx context.Context) error

	// IsConnected 检查连接状态
	// 返回值：
	//   - true: 连接正常
	//   - false: 连接断开
	// 使用场景：
	//   - 快速检查连接状态（不执行网络请求）
	//   - 判断是否需要重连
	IsConnected() bool
}

// HealthStatus 健康状态结构体
// 用于返回详细的健康检查结果
type HealthStatus struct {
	// Status 健康状态：healthy、unhealthy、degraded
	Status string `json:"status"`

	// Message 健康检查消息
	Message string `json:"message"`

	// Timestamp 检查时间戳
	Timestamp time.Time `json:"timestamp"`

	// Details 详细信息（可选）
	Details map[string]interface{} `json:"details,omitempty"`
}

// HealthStatus 常量定义
const (
	// HealthStatusHealthy 健康状态
	HealthStatusHealthy = "healthy"

	// HealthStatusUnhealthy 不健康状态
	HealthStatusUnhealthy = "unhealthy"

	// HealthStatusDegraded 降级状态
	HealthStatusDegraded = "degraded"
)

// DetailedHealthChecker 详细健康检查接口
// 支持返回详细健康状态信息
type DetailedHealthChecker interface {
	HealthChecker

	// DetailedHealthCheck 执行详细健康检查
	// 返回值：
	//   - HealthStatus: 详细健康状态信息
	//   - error: 检查失败时的错误
	// 使用场景：
	//   - 监控系统需要详细信息
	//   - 健康检查API响应
	DetailedHealthCheck(ctx context.Context) (HealthStatus, error)
}
