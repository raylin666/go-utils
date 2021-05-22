package http

import (
	"github.com/gin-gonic/gin"
	auth "go-server/internal/api/http/auth/router"
)

// API 接口路由
func Router(engine *gin.Engine) {
	// 注册 鉴权服务 路由
	router := engine.Group("/auth")
	{
		auth.Router(router)
	}
}
