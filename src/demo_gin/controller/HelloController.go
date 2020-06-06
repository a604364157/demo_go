package controller

import (
	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func (c *HelloController) Router(engine *gin.Engine) {
	engine.GET("/", c.Hello)
}

func (c *HelloController) Hello(context *gin.Context) {
	context.JSON(200, map[string]interface{}{
		"message": "hello world",
	})
}
