// Package timeutil 提供统一的超时配置管理
// 功能说明：
//   - 定义全局默认超时常量
//   - 提供超时 Context 创建工具函数
//
// 使用场景：
//   - 数据库操作超时控制
//   - HTTP请求超时控制
//   - Redis操作超时控制
//   - 外部服务调用超时控制
package timeutil

import (
	"context"
	"time"
)

// 默认超时时间常量
// 可直接使用或作为参考值
const (
	// DefaultConnectTimeout 连接超时时间
	// 适用场景：数据库连接、HTTP连接、Redis连接
	DefaultConnectTimeout = 5 * time.Second

	// DefaultReadTimeout 读操作超时时间
	// 适用场景：数据读取、HTTP响应读取
	DefaultReadTimeout = 10 * time.Second

	// DefaultWriteTimeout 写操作超时时间
	// 适用场景：数据写入、HTTP请求发送
	DefaultWriteTimeout = 10 * time.Second

	// DefaultQueryTimeout 查询操作超时时间
	// 适用场景：数据库查询、Redis查询
	DefaultQueryTimeout = 15 * time.Second

	// DefaultOperationTimeout 通用操作超时时间
	// 适用场景：一般性操作、外部服务调用
	DefaultOperationTimeout = 30 * time.Second

	// DefaultHealthTimeout 健康检查超时时间
	// 适用场景：连接健康检查、服务状态检查
	DefaultHealthTimeout = 3 * time.Second
)

// WithTimeout 创建带自定义超时的 Context
// 参数：
//   - parent: 父 Context
//   - timeout: 超时时间
//
// 返回值：
//   - context.Context: 带超时的 Context
//   - context.CancelFunc: 取消函数（必须调用）
//
// 使用示例：
//
//	ctx, cancel := WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
func WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, timeout)
}

// WithConnectTimeout 创建带连接超时的 Context
// 使用 DefaultConnectTimeout (5秒)
func WithConnectTimeout(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, DefaultConnectTimeout)
}

// WithReadTimeout 创建带读操作超时的 Context
// 使用 DefaultReadTimeout (10秒)
func WithReadTimeout(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, DefaultReadTimeout)
}

// WithQueryTimeout 创建带查询操作超时的 Context
// 使用 DefaultQueryTimeout (15秒)
func WithQueryTimeout(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, DefaultQueryTimeout)
}

// WithOperationTimeout 创建带通用操作超时的 Context
// 使用 DefaultOperationTimeout (30秒)
func WithOperationTimeout(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, DefaultOperationTimeout)
}