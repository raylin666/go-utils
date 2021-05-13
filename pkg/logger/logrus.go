package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"go-server/config"
	"go-server/internal/constant"
	"go-server/pkg/utils"
	"log"
	"path"
	"time"
)

var (
	WriteMaps map[string]*Logger
)

type H logrus.Fields

func (data H) Fields() logrus.Fields {
	return logrus.Fields(data)
}

type Logger struct {
	// 日志实例
	Instance *logrus.Logger
	// 文件名称(文件写入时存在值)
	FileName string
	// 日志级别
	Level logrus.Level
	// 日志格式
	Format logrus.Formatter
}

func InitLogger() {
	var logger Logger

	WriteMaps = make(map[string]*Logger)

	// 创建文件夹
	utils.CreateDirectory(config.Get().Logs.Path)

	// 注册日志写实例
	registerWriteMaps(map[string]*Logger{
		constant.LogApp:     logger.instance(constant.LogApp),
		constant.LogDb:      logger.instance(constant.LogDb),
		constant.LogCache:   logger.instance(constant.LogCache),
		constant.LogRequest: logger.instance(constant.LogRequest),
		constant.LogSql:     logger.instance(constant.LogSql),
	})
}

// 注意: 该注册方法必须在服务启动前调用, 否则会有问题
func registerWriteMaps(maps map[string]*Logger) map[string]*Logger {
	for filename, logger := range maps {
		WriteMaps[filename] = logger
	}
	return WriteMaps
}

// 获取打印日志实例
func New() *logrus.Logger {
	return logrus.StandardLogger()
}

// 获取日志写入文件实例
func NewWrite(filename string) *logrus.Logger {
	var (
		logger *Logger
		ok     bool
	)

	if logger, ok = WriteMaps[filename]; !ok {
		return New()
	}

	return logger.Instance
}

func (logger Logger) instance(filename string) *Logger {
	logger.FileName = filename
	return logger.create()
}

// 创建 Logger 实例 初始化配置
func (logger Logger) create() *Logger {
	l := logrus.New()

	// 设置日志级别
	if logger.Level == 0 {
		l.SetLevel(logrus.DebugLevel)
	} else {
		l.SetLevel(logger.Level)
	}

	// 设置日志格式
	if logger.Format == nil {
		l.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: constant.TimestampFormat,
		})
	} else {
		l.SetFormatter(logger.Format)
	}

	file := path.Join(config.Get().Logs.Path, logger.FileName)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		file+"-%Y-%m-%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(file),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err == nil {
		l.SetOutput(logWriter)
	} else {
		log.Printf("(%s) failed to create rotatelogs: %s", logger.FileName, err)
	}

	logger.Instance = l

	return &logger
}
