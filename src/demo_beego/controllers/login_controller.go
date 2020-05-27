package controllers

import (
	"demo_beego/commom/utils"
	"demo_beego/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	logs.GetBeeLogger().Info("入参:", username, password)
	user := models.QueryUserByName(username)
	if user.PassWord == utils.MD5(password) {
		c.SetSession("username", user)
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "登录成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名或密码错误"}
	}
	c.ServeJSON()
}
