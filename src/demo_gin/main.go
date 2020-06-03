package main

import (
	"demo_gin/config"
	"demo_gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	//读取配置
	cfg := config.GetConfig()
	//实例app
	app := gin.Default()
	//注册路由
	routers.RegisterRouter(app)
	//启动app
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}
