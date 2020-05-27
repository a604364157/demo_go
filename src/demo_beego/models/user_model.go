package models

import (
	"demo_beego/commom/constants"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int
	UserName   string `orm:"size(64)"`
	NickName   string `orm:"size(64)"`
	PassWord   string `orm:"size(64)"`
	Status     int
	CreateTime int64
}

func init() {
	orm.RunSyncdb(constants.DEFAULT, false, true)
}

func QueryUserById(id int) User {
	user := new(User)
	user.Id = id
	err := getDb().Read(user)
	if err != nil {
		logs.GetBeeLogger().Info("查询失败")
	}
	return *user
}

func QueryUserByName(name string) User {
	user := new(User)
	user.UserName = name
	err := getDb().Read(user, "UserName")
	if err != nil {
		logs.GetBeeLogger().Info("查询失败")
	}
	return *user
}

func QueryUsers(user User) []User {
	qt := getDb().QueryTable(&user)
	if user.Id != 0 {
		qt = qt.Filter("id", user.Id)
	}
	if user.UserName != "" {
		qt = qt.Filter("user_name", user.UserName)
	}
	if user.NickName != "" {
		qt = qt.Filter("nick_name", user.NickName)
	}
	if user.PassWord != "" {
		qt = qt.Filter("PassWord", user.PassWord)
	}
	if user.Id != 0 {
		qt = qt.Filter("status", user.Status)
	}
	if user.Id != 0 {
		qt = qt.Filter("create_time", user.CreateTime)
	}
	var users []User
	_, err := qt.All(&users)
	if err != nil {
		logs.GetBeeLogger().Info("查询失败")
	}
	return users
}

func InsertUser(user User) error {
	_, err := getDb().Insert(&user)
	if err != nil {
		logs.GetBeeLogger().Info("写入用户数据失败")
	}
	return err
}

func UpdateUser(user User) int64 {
	if user.Id == 0 {
		logs.GetBeeLogger().Info("传入model对象没有id")
		return 0
	}
	num, err := getDb().Update(&user)
	if err != nil {
		logs.GetBeeLogger().Info("修改用户数据失败")
	}
	return num
}

func DeleteUser(id int) bool {
	user := new(User)
	user.Id = id
	num, err := getDb().Delete(user)
	if err != nil {
		logs.GetBeeLogger().Info("删除用户数据失败")
	}
	return num > 0
}
