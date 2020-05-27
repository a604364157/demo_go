package controllers

import "demo_beego/models"

type HomeController struct {
	BaseController
}

/**
 * 请求：http://localhost:8080/
 * 请求类型：Get
 * 请求描述：
 */
func (c *HomeController) Get() {
	page, _ := c.GetInt("page")
	if page < 1 {
		page = 1
	}
	article := models.Article{}
	articles := models.QueryArticlePage(article, page, 10)
	c.Data["PageCode"] = 1
	c.Data["HasFooter"] = true
	c.Data["Content"] = models.MakeHomeBlocks(articles, c.IsLogin)
	c.TplName = "home.html"
}
