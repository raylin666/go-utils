package router

import (
	"github.com/gin-gonic/gin"
	"go-server/internal/api/context"
	"go-server/internal/api/http/api/v1/controller"
)

// 鉴权服务
func Auth(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	router := routerGroup.Group("/auth")
	{
		router.POST("/token", context.ContextHandler(controller.GetTokenAuth))
		router.POST("/token/verify", context.ContextHandler(controller.VerifyTokenAuth))
	}
	return router
}

