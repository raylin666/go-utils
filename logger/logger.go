// Package logger 提供基于 zap 的高性能 JSON 日志记录功能。
// 支持日志轮转、多级别输出、自定义字段等功能。
// 注意：日志文件权限已设置为 0640，目录权限 0750，符合安全最佳实践。
package logger

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志相关常量配置
const (
	// DefaultLevel 默认日志级别为 Info
	DefaultLevel = zapcore.InfoLevel

	// DefaultTimeLayout 默认时间格式为 RFC3339
	DefaultTimeLayout = time.RFC3339

	// 日志字段键名常量
	AppKey         = "app"         // 应用名称键
	EnvironmentKey = "environment" // 环境标识键
	TimeKey        = "time"        // 时间键
	LevelKey       = "level"       // 日志级别键
	NameKey        = "name"        // 日志器名称键
	CallerKey      = "caller"      // 调用者信息键
	MessageKey     = "message"     // 日志消息键
	StacktraceKey  = "stacktrace"  // 堆栈追踪键
)

// Logger 包装 zap.Logger，提供资源管理能力
// 主要功能：
//   - 继承 zap.Logger 的所有日志记录方法
//   - 提供 Close() 方法用于关闭文件句柄，释放资源
//   - 支持日志文件轮转和压缩
type Logger struct {
	*zap.Logger
	file  io.Writer
	level zap.AtomicLevel
}

func (l *Logger) Close() error {
	if l.file != nil {
		if closer, ok := l.file.(io.Closer); ok {
			return closer.Close()
		}
	}
	return nil
}

func (l *Logger) SetLevel(level zapcore.Level) {
	if l.level != (zap.AtomicLevel{}) {
		l.level.SetLevel(level)
	}
}

func (l *Logger) GetLevel() zapcore.Level {
	if l.level != (zap.AtomicLevel{}) {
		return l.level.Level()
	}
	return zapcore.InfoLevel
}

func (l *Logger) Level() zapcore.Level {
	return l.GetLevel()
}

func SetDebugLevel(l *Logger) {
	l.SetLevel(zapcore.DebugLevel)
}

func SetInfoLevel(l *Logger) {
	l.SetLevel(zapcore.InfoLevel)
}

func SetWarnLevel(l *Logger) {
	l.SetLevel(zapcore.WarnLevel)
}

func SetErrorLevel(l *Logger) {
	l.SetLevel(zapcore.ErrorLevel)
}

// NewJSONLogger 创建 JSON 格式的高性能日志记录器
// 功能特性：
//   - JSON 格式输出，便于日志分析系统解析
//   - 支持多输出目标（控制台 + 文件）
//   - 支持日志级别过滤
//   - 支持自定义字段和时间格式
//   - 支持日志文件轮转和压缩
//
// 参数：
//   - opts: 选项函数列表，用于配置日志行为
//
// 返回值：
//   - *Logger: 成功创建的日志记录器
//   - error: 初始化失败时的错误（如文件创建失败、权限不足等）
//
// 使用示例：
//
//	logger, err := NewJSONLogger(
//	    WithInfoLevel(),
//	    WithField("app", "my-service"),
//	    WithPathFileRotation("/var/log/app.log", PathFileRotationOption{
//	        MaxSize:    100,  // MB
//	        MaxBackups: 10,
//	        MaxAge:     30,   // days
//	        LocalTime:  true,
//	        Compress:   true,
//	    }),
//	)
//	if err != nil {
//	    panic(err)  // 初始化失败必须处理
//	}
//	defer logger.Close()
func NewJSONLogger(opts ...Option) (*Logger, error) {
	opt := &option{
		level:        DefaultLevel,
		fields:       make(map[string]string),
		levelEncoder: zapcore.LowercaseLevelEncoder,
		timeLayout:   DefaultTimeLayout,
	}

	for _, f := range opts {
		f(opt)
	}

	if opt.initErr != nil {
		return nil, opt.initErr
	}

	atomicLevel := zap.NewAtomicLevelAt(opt.level)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       TimeKey,
		LevelKey:      LevelKey,
		NameKey:       NameKey,
		CallerKey:     CallerKey,
		MessageKey:    MessageKey,
		StacktraceKey: StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   opt.levelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(opt.timeLayout))
		},
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= atomicLevel.Level() && lvl < zapcore.ErrorLevel
	})

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= atomicLevel.Level() && lvl >= zapcore.ErrorLevel
	})

	stdout := zapcore.Lock(os.Stdout)
	stderr := zapcore.Lock(os.Stderr)

	core := zapcore.NewTee()

	if !opt.disableConsole {
		core = zapcore.NewTee(
			zapcore.NewCore(jsonEncoder,
				zapcore.NewMultiWriteSyncer(stdout),
				lowPriority,
			),
			zapcore.NewCore(jsonEncoder,
				zapcore.NewMultiWriteSyncer(stderr),
				highPriority,
			),
		)
	}

	if opt.file != nil {
		core = zapcore.NewTee(core,
			zapcore.NewCore(jsonEncoder,
				zapcore.AddSync(opt.file),
				atomicLevel,
			),
		)
	}

	logger := zap.New(core,
		zap.AddCaller(),
		zap.ErrorOutput(stderr),
	)

	for key, value := range opt.fields {
		logger = logger.WithOptions(zap.Fields(zapcore.Field{
			Key:    key,
			Type:   zapcore.StringType,
			String: value,
		}))
	}

	return &Logger{
		Logger: logger,
		file:   opt.file,
		level:  atomicLevel,
	}, nil
}
