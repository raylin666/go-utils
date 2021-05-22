package router

import (
	"github.com/gin-gonic/gin"
	"go-server/internal/api/context"
	"go-server/internal/api/http/auth/controller"
)

// 鉴权服务
func Router(router *gin.RouterGroup) *gin.RouterGroup {
	{
		router.POST("/token", context.ContextHandler(controller.GetTokenAuth))
		router.POST("/token/verify", context.ContextHandler(controller.VerifyTokenAuth))
		router.POST("/token/refresh", context.ContextHandler(controller.RefreshTokenAuth))
		router.POST("/token/delete", context.ContextHandler(controller.DeleteTokenAuth))
	}
	return router
}

