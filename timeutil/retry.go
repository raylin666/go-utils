// Package timeutil 提供通用重试机制，支持指数退避策略
// 功能说明：
//   - 支持自定义重试次数和延迟时间
//   - 支持指数退避策略（延迟时间逐步增加）
//   - 支持上下文超时控制
//   - 支持自定义重试条件判断
// 使用场景：
//   - 数据库连接重试
//   - HTTP请求重试
//   - Redis操作重试
//   - 外部服务调用重试
package timeutil

import (
	"context"
	"fmt"
	"math"
	"time"
)

// RetryConfig 重试配置
// 用于控制重试行为和策略
type RetryConfig struct {
	// MaxAttempts 最大重试次数
	// 默认值：3
	// 说明：包含首次尝试，如设置为3表示最多尝试3次
	MaxAttempts int

	// InitialDelay 初始延迟时间
	// 默认值：100ms
	// 说明：第一次重试前的等待时间
	InitialDelay time.Duration

	// MaxDelay 最大延迟时间
	// 默认值：5s
	// 说明：延迟时间上限，避免过长等待
	MaxDelay time.Duration

	// Multiplier 延迟时间乘数（指数退避）
	// 默认值：2.0
	// 说明：每次重试延迟时间乘以此值，实现指数退避
	// 示例：InitialDelay=100ms, Multiplier=2.0
	//       第1次重试: 100ms
	//       第2次重试: 200ms
	//       第3次重试: 400ms
	Multiplier float64

	// RetryCondition 重试条件判断函数
	// 参数：error - 操作返回的错误
	// 返回值：bool - true表示需要重试，false表示不需要重试
	// 说明：可用于判断特定错误是否需要重试
	// 示例：只重试网络错误，不重试参数错误
	RetryCondition func(error) bool
}

// DefaultRetryConfig 默认重试配置
// 提供合理的默认值，适用于大多数场景
var DefaultRetryConfig = RetryConfig{
	MaxAttempts:     3,              // 最多尝试3次
	InitialDelay:    100 * time.Millisecond, // 初始延迟100ms
	MaxDelay:        5 * time.Second,        // 最大延迟5s
	Multiplier:      2.0,                   // 延迟时间翻倍
	RetryCondition:  nil,                   // 默认所有错误都重试
}

// RetryError 重试失败错误
// 包含重试次数和最后一次错误信息
type RetryError struct {
	// Attempts 实际尝试次数
	Attempts int

	// LastErr 最后一次错误
	LastErr error

	// TotalDuration 总耗时
	TotalDuration time.Duration
}

func (e *RetryError) Error() string {
	return fmt.Sprintf("retry failed: attempts=%d, duration=%v, last_error=%v", 
		e.Attempts, e.TotalDuration, e.LastErr)
}

func (e *RetryError) Unwrap() error {
	return e.LastErr
}

// RetryWithBackoff 使用指数退避策略执行重试
// 功能说明：
//   - 使用指数退避策略重试操作
//   - 延迟时间逐步增加，避免频繁重试
//   - 支持上下文超时控制
//   - 支持自定义重试条件判断
//
// 参数：
//   - ctx: 上下文，用于超时控制
//   - config: 重试配置
//   - fn: 需要重试的操作函数
//
// 返回值：
//   - error: nil表示成功，非nil表示重试失败
//
// 使用示例：
//   err := RetryWithBackoff(ctx, DefaultRetryConfig, func() error {
//       return client.Connect()
//   })
//   if err != nil {
//       if retryErr, ok := err.(*RetryError); ok {
//           log.Printf("retry failed after %d attempts", retryErr.Attempts)
//       }
//   }
//
// 自定义配置示例：
//   config := RetryConfig{
//       MaxAttempts:     5,
//       InitialDelay:    200 * time.Millisecond,
//       MaxDelay:        10 * time.Second,
//       Multiplier:      1.5,
//       RetryCondition: func(err error) bool {
//           // 只重试网络错误
//           return IsNetworkError(err)
//       },
//   }
//   err := RetryWithBackoff(ctx, config, fn)
func RetryWithBackoff(ctx context.Context, config RetryConfig, fn func() error) error {
	if config.MaxAttempts <= 0 {
		config.MaxAttempts = DefaultRetryConfig.MaxAttempts
	}
	if config.InitialDelay <= 0 {
		config.InitialDelay = DefaultRetryConfig.InitialDelay
	}
	if config.MaxDelay <= 0 {
		config.MaxDelay = DefaultRetryConfig.MaxDelay
	}
	if config.Multiplier <= 0 {
		config.Multiplier = DefaultRetryConfig.Multiplier
	}

	var delay = config.InitialDelay
	var startTime = time.Now()
	var attempts = 0

	for i := 0; i < config.MaxAttempts; i++ {
		attempts++
		
		err := fn()
		if err == nil {
			return nil
		}

		// 判断是否需要重试
		if config.RetryCondition != nil && !config.RetryCondition(err) {
			return &RetryError{
				Attempts:      attempts,
				LastErr:       err,
				TotalDuration: time.Since(startTime),
			}
		}

		// 最后一次尝试失败，不再等待
		if i == config.MaxAttempts - 1 {
			return &RetryError{
				Attempts:      attempts,
				LastErr:       err,
				TotalDuration: time.Since(startTime),
			}
		}

		// 等待下次重试
		select {
		case <-ctx.Done():
			return &RetryError{
				Attempts:      attempts,
				LastErr:       fmt.Errorf("retry timeout: %w", ctx.Err()),
				TotalDuration: time.Since(startTime),
			}
		case <-time.After(delay):
			// 计算下次延迟时间（指数退避）
			delay = time.Duration(float64(delay) * config.Multiplier)
			if delay > config.MaxDelay {
				delay = config.MaxDelay
			}
		}
	}

	return nil
}

// RetryWithFixedDelay 使用固定延迟策略执行重试
// 功能说明：
//   - 每次重试使用相同的延迟时间
//   - 适用于需要稳定重试间隔的场景
//
// 参数：
//   - ctx: 上下文，用于超时控制
//   - maxAttempts: 最大重试次数
//   - delay: 固定延迟时间
//   - fn: 需要重试的操作函数
//
// 返回值：
//   - error: nil表示成功，非nil表示重试失败
//
// 使用示例：
//   err := RetryWithFixedDelay(ctx, 3, 1*time.Second, func() error {
//       return client.Ping()
//   })
func RetryWithFixedDelay(ctx context.Context, maxAttempts int, delay time.Duration, fn func() error) error {
	config := RetryConfig{
		MaxAttempts:  maxAttempts,
		InitialDelay: delay,
		MaxDelay:     delay,
		Multiplier:   1.0, // 固定延迟，不使用指数退避
	}
	return RetryWithBackoff(ctx, config, fn)
}

// RetryWithJitter 使用抖动策略执行重试
// 功能说明：
//   - 在指数退避基础上添加随机抖动
//   - 避免多个客户端同时重试导致的"惊群效应"
//   - 抖动范围为延迟时间的0.5-1.5倍
//
// 参数：
//   - ctx: 上下文，用于超时控制
//   - config: 重试配置
//   - fn: 需要重试的操作函数
//
// 返回值：
//   - error: nil表示成功，非nil表示重试失败
//
// 使用场景：
//   - 多客户端并发场景
//   - 分布式系统重试
//   - 避免重试风暴
func RetryWithJitter(ctx context.Context, config RetryConfig, fn func() error) error {
	if config.MaxAttempts <= 0 {
		config.MaxAttempts = DefaultRetryConfig.MaxAttempts
	}
	if config.InitialDelay <= 0 {
		config.InitialDelay = DefaultRetryConfig.InitialDelay
	}
	if config.MaxDelay <= 0 {
		config.MaxDelay = DefaultRetryConfig.MaxDelay
	}
	if config.Multiplier <= 0 {
		config.Multiplier = DefaultRetryConfig.Multiplier
	}

	var delay = config.InitialDelay
	var startTime = time.Now()
	var attempts = 0

	for i := 0; i < config.MaxAttempts; i++ {
		attempts++
		
		err := fn()
		if err == nil {
			return nil
		}

		if config.RetryCondition != nil && !config.RetryCondition(err) {
			return &RetryError{
				Attempts:      attempts,
				LastErr:       err,
				TotalDuration: time.Since(startTime),
			}
		}

		if i == config.MaxAttempts - 1 {
			return &RetryError{
				Attempts:      attempts,
				LastErr:       err,
				TotalDuration: time.Since(startTime),
			}
		}

		// 添加抖动（0.5-1.5倍）
		jitter := delay / 2
		randomFactor := 0.5 + float64(time.Now().UnixNano()%1000) / 1000.0
		actualDelay := delay + time.Duration(float64(jitter) * randomFactor)
		
		select {
		case <-ctx.Done():
			return &RetryError{
				Attempts:      attempts,
				LastErr:       fmt.Errorf("retry timeout: %w", ctx.Err()),
				TotalDuration: time.Since(startTime),
			}
		case <-time.After(actualDelay):
			delay = time.Duration(float64(delay) * config.Multiplier)
			if delay > config.MaxDelay {
				delay = config.MaxDelay
			}
		}
	}

	return nil
}

// IsRetryError 判断错误是否为重试错误
// 参数：
//   - err: 错误对象
// 返回值：
//   - bool: 是否为重试错误
func IsRetryError(err error) bool {
	_, ok := err.(*RetryError)
	return ok
}

// GetRetryAttempts 获取重试次数
// 参数：
//   - err: 错误对象
// 返回值：
//   - int: 重试次数（如果不是重试错误返回0）
func GetRetryAttempts(err error) int {
	if retryErr, ok := err.(*RetryError); ok {
		return retryErr.Attempts
	}
	return 0
}

// CalculateBackoffDelay 计算指数退避延迟时间
// 功能说明：
//   - 根据重试次数计算延迟时间
//   - 用于自定义重试策略
//
// 参数：
//   - attempt: 当前重试次数（从1开始）
//   - initialDelay: 初始延迟时间
//   - maxDelay: 最大延迟时间
//   - multiplier: 延迟时间乘数
//
// 返回值：
//   - time.Duration: 计算后的延迟时间
//
// 使用示例：
//   delay := CalculateBackoffDelay(3, 100*time.Millisecond, 5*time.Second, 2.0)
//   // 返回: 400ms (100 * 2^2)
func CalculateBackoffDelay(attempt int, initialDelay time.Duration, maxDelay time.Duration, multiplier float64) time.Duration {
	delay := float64(initialDelay) * math.Pow(multiplier, float64(attempt-1))
	result := time.Duration(delay)
	if result > maxDelay {
		result = maxDelay
	}
	return result
}