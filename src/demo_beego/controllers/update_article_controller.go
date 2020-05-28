package controllers

import (
	"demo_beego/models"
)

type UpdateArticleController struct {
	BaseController
}

func (c *UpdateArticleController) Get() {
	id, _ := c.GetInt("id")
	//获取id所对应的文章信息
	art := models.QueryArticleById(id)
	c.Data["Title"] = art.Title
	c.Data["Tags"] = art.Tags
	c.Data["Short"] = art.Short
	c.Data["Content"] = art.Content
	c.Data["Id"] = art.Id
	c.TplName = "write_article.html"
}

func (c *UpdateArticleController) Post() {
	id, _ := c.GetInt("id")
	//获取浏览器传输的数据，通过表单的name属性获取值
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")
	//获取session
	login := c.LoginUser
	user := login.(models.User)
	//实例化model，修改数据库
	art := models.Article{Id: id, Title: title, Author: user.UserName, Tags: tags, Short: short, Content: content}
	err := models.UpdateArticle(art)
	//返回数据给浏览器
	if err == nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "更新成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "更新失败"}
	}
	c.ServeJSON()
}
