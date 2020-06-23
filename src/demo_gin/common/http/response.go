package http

import "github.com/gin-gonic/gin"

func SUCCESS(context *gin.Context, msg interface{}) {
	context.JSON(200, map[string]interface{}{
		"code": "0",
		"msg":  msg,
	})
}

func ERROR(context *gin.Context, msg interface{}) {
	context.JSON(200, map[string]interface{}{
		"code": "1",
		"msg":  msg,
	})
}

func RESPONSE(context *gin.Context, code int, msg interface{}) {
	context.JSON(code, msg)
}
