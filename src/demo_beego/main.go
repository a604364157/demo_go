package main

import (
	_ "demo_beego/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.WebConfig.Session.SessionOn = true
	//开启session,可以在配置里加sessionon = true
	beego.Run("localhost:8080")
}

//beego框架

/*
程序先找到routers包,根据路径映射到相应的controller
controllers包下的MainController只有一个GET方法go
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
数据和模板,然后响应

1:项目结构
conf包:配置文件
controllers包: 控制器层
models包: 数据层
routers包: 路由层
static包: 静态资源(html,js,css等)
views包: tpl模板文件

2:路由配置
	1,固定路由 beego.Router("/",controller);
	2,正则路由
	3,自定义路由

3:静态文件的设置
	beego.SetStaticPath("/down1","download1")

本工程实战开发一个博客项目

*/
