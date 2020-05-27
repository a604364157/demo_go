package controllers

import (
	"demo_beego/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

func (c *AddArticleController) Get() {
	c.TplName = "write_article.html"
}

func (c *AddArticleController) Post() {
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")
	user := c.LoginUser
	art := models.Article{Title: title, Author: user.UserName, Tags: tags, Short: short, Content: content, CreateTime: time.Now().Unix()}
	err := models.InsertArticle(art)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "ok"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "保存失败"}
	}
	c.ServeJSON()
}
