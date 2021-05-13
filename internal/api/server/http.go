package server

import (
	"fmt"
	"go-server/config"
	"go-server/internal/api/router"
	"net/http"
	"time"
)

// 创建服务
func NewHttpServer(r *router.Router) {
	// 启动服务
	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", config.Get().Http.Host, config.Get().Http.Port),
		Handler:        r.New(),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server startup err: %v", err))
	}
}
