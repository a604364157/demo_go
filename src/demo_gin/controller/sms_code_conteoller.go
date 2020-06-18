package controller

import (
	"demo_gin/common/http"
	"demo_gin/common/utils"
	"demo_gin/models"
	"github.com/gin-gonic/gin"
	"time"
)

type SmsCodeController struct {
}

func (c *SmsCodeController) Router(engine *gin.Engine) {
	engine.GET("/api/sms/code", c.send)
}

// 发送验证码
func (c *SmsCodeController) send(context *gin.Context) {
	// 验证用户名格式
	userName := context.Query("userName")
	if !utils.RegexpEmail(userName) {
		http.ERROR(context, "用户名格式错误,请输入正确的邮箱格式")
	}
	// 生成数字验证码
	code := utils.RandString(6)
	// 写入数据库
	models.InsertSmsCode(models.SmsCode{Email: userName, Code: code, CreateTime: time.Now().Unix()})
}
