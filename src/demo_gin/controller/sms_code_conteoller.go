package controller

import (
	"demo_gin/common/http"
	"demo_gin/common/utils"
	"demo_gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type SmsCodeController struct {
}

func (c *SmsCodeController) Router(engine *gin.Engine) {
	engine.GET("/api/sms/code", c.send)
}

// 发送邮件验证码
func (c *SmsCodeController) send(context *gin.Context) {
	// 验证用户名格式
	userName := context.Query("userName")
	if !utils.RegexpEmail(userName) {
		http.ERROR(context, "用户名格式错误,请输入正确的邮箱格式")
		return
	}
	// 生成数字验证码
	code := utils.RandString(6)
	// 通过用户名查询历史验证信息
	email := models.SmsCode{Email: userName}
	smsCodes, err := models.QuerySmsCode(email)
	if len(smsCodes) > 0 {
		// 修改数据
		smsCode := smsCodes[0]
		smsCode.Code = code
		smsCode.CreateTime = time.Now().Unix()
		err = models.UpdateSmsCode(smsCode)
		if err != nil {
			http.ERROR(context, "系统异常,请稍后重试或联系管理员")
			return
		}
	} else {
		// 写入数据库
		err = models.InsertSmsCode(models.SmsCode{Email: userName, Code: code, CreateTime: time.Now().Unix()})
		if err != nil {
			http.ERROR(context, "系统异常,请稍后重试或联系管理员")
			return
		}
	}
	// 发送验证码
	users := []string{userName}
	err = utils.SendMail(users, "demo_gin登录验证", fmt.Sprintf("你好: 你的登录验证码为[%s], 验证码有效期为2分钟!如非本人操作请忽略本消息。", code))
	if err != nil {
		http.ERROR(context, "系统异常,请稍后重试或联系管理员")
		return
	}
}
