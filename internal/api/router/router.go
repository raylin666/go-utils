package router

import (
	"github.com/gin-gonic/gin"
	"go-server/internal/constant"
	"go-server/internal/environment"
)

type Router struct {
	Before func(*gin.Engine)
	After  func(*gin.Engine)
}

// 创建路由
func (r *Router) New() *gin.Engine {
	var currentEnvironment = gin.ReleaseMode
	switch environment.GetEnvironment().Value() {
	case constant.EnvironmentPre, constant.EnvironmentProd:
		currentEnvironment = gin.ReleaseMode
	case constant.EnvironmentDev:
		currentEnvironment = gin.DebugMode
	case constant.EnvironmentTest:
		currentEnvironment = gin.TestMode
	default:
	}

	gin.SetMode(currentEnvironment)

	engine := gin.New()

	// 可加载中间件配置
	if r.Before != nil {
		r.Before(engine)
	}

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// 可加载路由配置
	if r.After != nil {
		r.After(engine)
	}

	return engine
}

