package main

import (
	"github.com/gin-gonic/gin"
	"go-server/config"
	"go-server/internal/api/middleware"
	"go-server/internal/api/router"
	"go-server/internal/api/server"
	"go-server/internal/environment"
	"go-server/internal/model"
	"go-server/pkg/cache"
	"go-server/pkg/database"
	"go-server/pkg/logger"
)

func init()  {
	// 初始化加载配置文件
	config.InitAutoloadConfig(".env.yml")
	// 初始化环境
	environment.InitEnvironment()
	// 日志初始化
	logger.InitLogger()
	// 数据库初始化
	database.InitDatabase()
	// 初始化数据库模型
	model.InitModel()
	// 缓存初始化
	cache.InitRedis()
}

func main()  {
	r := &router.Router{
		Before: func(engine *gin.Engine) {
			engine.Use(middleware.RequestLoggerWrite())
		},
		After: func(engine *gin.Engine) {
			router.RouterApi(engine)
		},
	}

	server.NewHttpServer(r)
}
