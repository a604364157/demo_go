package controllers

type HomeController struct {
	BaseController
}

/**
 * 请求：http://localhost:8080/
 * 请求类型：Get
 * 请求描述：
 */
func (c *HomeController) Get() {
	c.TplName = "home.html"
}
