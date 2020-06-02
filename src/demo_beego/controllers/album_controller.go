package controllers

import "demo_beego/models"

type AlbumController struct {
	BaseController
}

func (c *AlbumController) Get() {
	albums := models.QueryAllAlbums()
	c.Data["Album"] = albums
	c.TplName = "album.html"
}
