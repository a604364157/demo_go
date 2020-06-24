package controller

import (
	"demo_gin/common/http"
	"demo_gin/common/session"
	"demo_gin/common/utils"
	"demo_gin/models"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

type UserController struct {
}

func (c *UserController) Router(engine *gin.Engine) {
	engine.GET("/api/user", c.Get)
	engine.POST("/api/user/register", c.Post)
}

// 通过用户名查询用户信息
func (c *UserController) Get(context *gin.Context) {
	userName := context.Query("userName")
	user := new(models.User)
	user.UserName = userName
	users := models.QueryUser(*user)
	if len(users) < 1 {
		http.ERROR(context, "用户名不存在")
		return
	}
	user = &users[0]
	// 用户信息不展示密码
	user.Password = ""
	http.SUCCESS(context, user)
}

// 用户注册
func (c *UserController) Post(context *gin.Context) {
	// 获取body参数
	var user models.User
	err := utils.JsonParse(context.Request.Body, &user)
	if err != nil {
		http.RESPONSE(context, 400, "参数序列化失败")
		return
	}
	if user.UserName == "" {
		http.ERROR(context, "用户名不能为空")
		return
	}
	if !utils.RegexpEmail(user.UserName) {
		http.ERROR(context, "用户名格式错误,请输入正确的邮箱格式")
		return
	}
	if user.Password == "" {
		http.ERROR(context, "用户密码不能为空")
		return
	}
	// 验证用户名是否重复
	users := models.QueryUser(user)
	if len(users) >= 1 {
		http.ERROR(context, "用户数据异常,请联系管理员调整")
		return
	}
	flag := false
	// 设置一个特殊标识,用于判断是否设置session
	if user.IsActive == 10 {
		flag = true
	}
	// 加密密码
	user.Password = utils.MD5(user.Password)
	// 设置默认值
	user.IsActive = 1
	user.RegisterTime = time.Now().Unix()
	// 新增用户信息(其他的字段暂时为空算了)
	err = models.InsertUser(user)
	if err != nil {
		http.ERROR(context, "注册失败")
		return
	}
	// 设置session(用于短信验证码登录时自动注册+登录)
	if flag {
		session.SetSession(context, user.UserName, user)
	}
}

// 上传头像
func (c *UserController) uploadAvatar(context *gin.Context) {
	//1、获取上传的文件
	userName := context.Request.PostFormValue("userName")
	file, header, err := context.Request.FormFile("avatar")
	if err != nil {
		http.ERROR(context, "获取上传的头像文件失败")
		return
	}
	//2、将文件保存到本地
	fileFullPath := "./uploadfile/" + header.Filename
	out, err := os.Create(fileFullPath)
	defer out.Close()
	_, err = io.Copy(out, file)
	// 用户名不为空则走已登录流程
	if userName != "" {
		user := session.GetSession(context, userName).(models.User)
		if user.UserName == "" {
			http.ERROR(context, "你未登录或登录已失效")
			return
		}
	}
}
