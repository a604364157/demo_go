package controllers

import (
	"demo_beego/commom/utils"
	"demo_beego/models"
	"strconv"
)

type ShowArticleController struct {
	BaseController
}

func (c *ShowArticleController) Get() {
	param := c.Ctx.Input.Param(":id")
	//字符转int
	id, _ := strconv.Atoi(param)
	art := models.QueryArticleById(id)
	c.Data["Title"] = art.Title
	c.Data["Content"] = utils.MarkdownToHtml(art.Content)
	c.TplName = "show_article.html"
}
