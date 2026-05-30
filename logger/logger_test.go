package logger

import (
	"errors"
	"testing"
)

// TestJSONLogger 测试 JSON 日志记录器功能
// 测试内容：
//   - 日志记录器初始化
//   - 自定义字段添加
//   - 错误日志记录
//   - 资源释放（Close）
func TestJSONLogger(t *testing.T) {
	// 创建日志记录器（现在返回两个值）
	logger, err := NewJSONLogger(
		WithField("defined_key", "defined_value"),
	)
	
	// 检查初始化错误（直接处理错误，无需调用 InitError）
	if err != nil {
		t.Fatalf("日志记录器初始化失败: %v", err)
	}
	
	// 确保资源释放
	defer logger.Sync()
	defer logger.Close()

	// 测试错误日志记录
	testErr := errors.New("pkg error")
	logger.Error("err occurs", WrapMeta(nil, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)
	logger.Error("err occurs", WrapMeta(testErr, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)
}

// BenchmarkJsonLogger 性能基准测试
// 测试日志记录器的性能表现
func BenchmarkJsonLogger(b *testing.B) {
	b.ResetTimer()
	
	// 创建日志记录器（现在返回两个值）
	logger, err := NewJSONLogger(
		WithField("defined_key", "defined_value"),
	)
	
	// 检查初始化错误
	if err != nil {
		b.Fatalf("日志记录器初始化失败: %v", err)
	}
	
	// 确保资源释放
	defer logger.Sync()
	defer logger.Close()
	
	// 性能测试循环
	for i := 0; i < b.N; i++ {
		logger.Info("benchmark test")
	}
}