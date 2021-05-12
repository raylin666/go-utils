package main

import (
	"github.com/raylin666/go-gin-api/initx"
	"go-server/config"
	"go-server/internal/model"
	"go-server/internal/server"
	"sync"
)

func init()  {
	var YmlEnvFileName = ".env.yml"

	initx.InitApplication(&initx.Config{
		YmlEnvFileName: YmlEnvFileName,
	})

	// 追加配置项
	config.InitConfig(YmlEnvFileName)

	// Model 初始化
	model.InitModel()
}

func main()  {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		// 创建 HTTP 服务
		server.NewHttpServer()
	}()

	go func() {
		defer wg.Done()
		// 创建 gRPC 服务
		server.NewGrpcServer()
	}()

	wg.Wait()
}
