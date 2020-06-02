package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["weChat"] = "微信：18215685065"
	c.Data["qq"] = "QQ：604364157"
	c.Data["tel"] = "Tel：18215685065"
	c.TplName = "aboutMe.html"
}
