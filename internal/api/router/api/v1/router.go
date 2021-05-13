package v1

import (
	"github.com/gin-gonic/gin"
	"go-server/internal/api/context"
	v1 "go-server/internal/api/controller/api/v1"
)

func Router(routerGroup *gin.RouterGroup) {
	auth(routerGroup)
}

func auth(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	router := routerGroup.Group("/auth")
	{
		router.POST("/token", context.ContextHandler(v1.GetTokenAuth))
	}
	return router
}

