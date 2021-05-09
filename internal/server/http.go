package server

import (
	"github.com/gin-gonic/gin"
	"github.com/raylin666/go-gin-api/middleware"
	"github.com/raylin666/go-gin-api/router"
	"github.com/raylin666/go-gin-api/server"
)

// 创建 HTTP 服务
func NewHttpServer()  {
	r := &router.Router{
		Before: func(engine *gin.Engine) {
			engine.Use(middleware.RequestLoggerWrite())
		},
		After: func(engine *gin.Engine) {
			// 路由配置 ...
		},
	}

	server.New(r)
}
