package router

import (
	"github.com/gin-gonic/gin"
	api "go-server/internal/api/http/api/v1/router"
)

// API 接口路由
func Router(engine *gin.Engine) {
	// 注册 v1 路由
	router := engine.Group("/api/v1")
	{
		api.Auth(router)
	}
}
