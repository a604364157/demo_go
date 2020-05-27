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
	//获取session
	login := c.LoginUser
	//因为我们在LoginUser存的是models.User,这里可以直接转
	user := login.(models.User)
	art := models.Article{Title: title, Author: user.UserName, Tags: tags, Short: short, Content: content, CreateTime: time.Now().Unix()}
	err := models.InsertArticle(art)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "保存失败"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "ok"}
	}
	c.ServeJSON()
}
