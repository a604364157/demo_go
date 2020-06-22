package controller

import (
	"demo_gin/common/http"
	"demo_gin/common/session"
	"demo_gin/common/utils"
	"demo_gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type LoginController struct {
}

func (c *LoginController) Router(engine *gin.Engine) {
	engine.GET("/api/login/ca", c.Captcha)
	engine.GET("/api/login/pwd", c.LoginByPwd)
	engine.GET("/api/login/code", c.LoginByCode)
}

// 生成图片验证码
func (c *LoginController) Captcha(context *gin.Context) {
	ca := new(utils.Captcha)
	err := utils.GenerateCaptcha(ca)
	if err != nil {
		http.ERROR(context, "获取验证码失败,请稍后再试")
	}
	http.SUCCESS(context, ca)
}

// 实现登录功能:用户名密码
func (c *LoginController) LoginByPwd(context *gin.Context) {
	// 校验验证码
	c.VerifyCa(context)
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
	c.VerifyEmail(context)
	// 验证通过后,判断是否为注册用户
	userName := context.Query("userName")
	users := models.QueryUser(models.User{UserName: userName})
	if len(users) > 0 {
		session.SetSession(context, userName, users[0])
		http.SUCCESS(context, "登录成功")
	} else {
		// 非注册用户让用户跳转信息完善页面(信息完善后再存入session信息)
		http.SUCCESS(context, "前往注册")
	}
}

func (c *LoginController) register(context *gin.Context) {
	body := context.Request.GetBody
	fmt.Println(body())
}

// 验证图形验证码
func (c *LoginController) VerifyCa(context *gin.Context) {
	caId := context.Query("captchaId")
	caValue := context.Query("captchaValue")
	if caId == "" || caValue == "" {
		http.ERROR(context, "请输入验证码")
	}
	// 校验验证码
	ca := new(utils.Captcha)
	ca.Id = caId
	ca.VerifyValue = caValue
	v := utils.VerifyCaptcha(ca)
	if !v {
		http.ERROR(context, "验证码错误")
	}
}

// 验证邮箱验证码
func (c *LoginController) VerifyEmail(context *gin.Context) {
	userName := context.Query("userName")
	code := context.Query("code")
	// 验证用户名格式
	if !utils.RegexpEmail(userName) {
		http.ERROR(context, "用户名格式错误,请输入正确的邮箱格式")
	}
	// 通过用户名查询验证信息
	email := models.SmsCode{Email: userName}
	smsCodes, _ := models.QuerySmsCode(email)
	if len(smsCodes) < 1 {
		http.ERROR(context, "验证码错误")
	}
	smsCode := smsCodes[0]
	// 验证有效时间
	t := smsCode.CreateTime
	now := time.Now().Unix()
	if now-t > 2*60*1000 {
		http.ERROR(context, "验证码已过有效期, 请重新获取")
	}
	// 验证验证码
	if !strings.EqualFold(code, smsCode.Code) {
		http.ERROR(context, "验证码错误")
	}
}
