package controllers

import (
	"demo_beego/commom/utils"
	"demo_beego/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

//处理注册
func (c *RegisterController) Post() {
	//获取表单信息
	username := c.GetString("username")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	hasBlank := username == "" || password == "" || repassword == ""
	if hasBlank {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "传入参数存在空值"}
		c.ServeJSON()
		return
	}
	if password != repassword {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "两次密码输入不一致"}
		c.ServeJSON()
		return
	}
	logs.GetBeeLogger().Info(username, password, repassword)
	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	user := new(models.User)
	user.UserName = username
	users := models.QueryUsers(*user)
	if len(users) > 0 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名已经存在"}
		c.ServeJSON()
		return
	}
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	logs.GetBeeLogger().Info("md5后：", password)
	user2 := models.User{0, username, username, password, 0, time.Now().Unix()}
	err := models.InsertUser(user2)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	c.ServeJSON()
}
