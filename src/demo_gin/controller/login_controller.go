package controller

import (
	"demo_gin/common/http"
	"demo_gin/common/session"
	"demo_gin/common/utils"
	"demo_gin/models"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

func (c *LoginController) Router(engine *gin.Engine) {
	engine.GET("/api/login/pwd", c.LoginByPwd)
	engine.GET("/api/login/code", c.LoginByCode)
}

// 实现登录功能:用户名密码
func (c *LoginController) LoginByPwd(context *gin.Context) {
	userName := context.Query("userName")
	password := context.Query("password")
	if userName == "" || password == "" {
		http.ERROR(context, "请输入完整用户名或密码")
	}
	user := models.User{UserName: userName, Password: utils.MD5(password)}
	users := models.QueryUser(user)
	if len(users) < 1 {
		http.ERROR(context, "用户名或密码错误")
	}
	if len(users) > 1 {
		http.ERROR(context, "用户数据异常,请联系管理员调整")
	}
	// 写入session
	session.SetSession(context, userName, users[0])
	http.SUCCESS(context, "登录成功")
}

// 实现登录功能:邮箱和验证码
func (c *LoginController) LoginByCode(context *gin.Context) {
	http.SUCCESS(context, "登录成功")
}
