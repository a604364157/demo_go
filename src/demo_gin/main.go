package main

import (
	"demo_gin/common/session"
	"demo_gin/config"
	"demo_gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 读取配置
	cfg := config.GetConfig()
	// 实例app
	app := gin.Default()
	// 注册路由
	routers.RegisterRouter(app)
	// 初始化session
	session.InitSession(app)
	// 启动app
	err := app.Run(cfg.AppHost + ":" + cfg.AppPort)
	if err != nil {
		panic(err)
	}
}
