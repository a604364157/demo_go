package controller

import (
	"demo_gin/common/utils"
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
		utils.ERROR(context, "用户名或密码错误")
	}

	utils.SUCCESS(context, "登录成功")
}

// 实现登录功能:邮箱和验证码
func (c *LoginController) LoginByCode(context *gin.Context) {
	utils.SUCCESS(context, "登录成功")
}
