package controllers

import (
	"demo_beego/models"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type DeleteArticleController struct {
	BaseController
}

//点击删除后重定向到首页
func (c *DeleteArticleController) Get() {
	artID, _ := c.GetInt("id")
	fmt.Println("删除 id:", artID)

	b := models.DeleteArticle(artID)
	if !b {
		logs.GetBeeLogger().Info("删除失败")
	}
	c.Redirect("/", 302)
}
