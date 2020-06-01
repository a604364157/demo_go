package controllers

import "demo_beego/models"

type TagsController struct {
	BaseController
}

func (c *TagsController) Get() {
	tags := models.QueryArticleTags()
	c.Data["Tags"] = models.HandleTags(tags)
	c.TplName = "tags.html"
}
