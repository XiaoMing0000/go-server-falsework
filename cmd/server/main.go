package main

import (
	"fmt"
	"log"
	"main/internal/app"
	"main/internal/router"

	"github.com/gin-gonic/gin"
)

// 参考文档：https://gin-gonic.com/zh-cn/docs/routing/querystring-param/

func main() {
	logger := app.Logger()

	// 初始化配置
	app.GetConfig()

	// 启动服务
	engine := gin.New()

	// 配置路由
	router.Router(engine.Group(""))

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	var addr = fmt.Sprintf("%s:%d", app.GetConfig().Server.Host, app.GetConfig().Server.Port)
	if err := engine.Run(addr); err != nil {
		logger.Error(fmt.Sprintf("failed to run server: %+v", err))
		log.Fatal(err)
	}
	logger.Info(fmt.Sprintf("Start server on %s\n", addr))

}
