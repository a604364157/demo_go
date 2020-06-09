package routers

import (
	"demo_gin/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.SmsCodeController).Router(router)
	new(controller.UserController).Router(router)
}
