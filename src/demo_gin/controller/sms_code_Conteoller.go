package controller

import (
	"demo_gin/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SmsCodeController struct {
}

func (c *SmsCodeController) Router(engine *gin.Engine) {
	engine.GET("/api/sms/code", c.Get)
}

func (c *SmsCodeController) Get(context *gin.Context) {
	idStr := context.Query("id")
	id, _ := strconv.Atoi(idStr)
	smsCode, _ := models.QuerySmsCodeById(id)
	context.JSON(200, map[string]interface{}{
		"message": smsCode,
	})
}
