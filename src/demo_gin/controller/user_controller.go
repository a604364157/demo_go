package controller

import (
	"demo_gin/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c *UserController) Router(engine *gin.Engine) {
	engine.GET("/user", c.Get)
}

func (c *UserController) Get(context *gin.Context) {
	userName := context.Query("userName")
	user := new(models.User)
	user.UserName = userName
	users := models.QueryUser(*user)
	context.JSON(200, map[string]interface{}{
		"message": users,
	})
}
