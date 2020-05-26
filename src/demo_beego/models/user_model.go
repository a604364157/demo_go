package models

import (
	"demo_beego/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type User struct {
	Id         int
	Username   string
	Nickname   string
	Password   string
	Status     string
	Createtime int64
}

//写入
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,nickname,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

//通过条件查询单条
func QueryUser(where string) *User {
	sql := fmt.Sprintf("select id,username,password,nickname,status,createtime from users %s", where)
	logs.GetBeeLogger().Info(sql)
	row := utils.QueryRow(sql)
	user := new(User)
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Nickname, &user.Status, &user.Createtime)
	return user
}

//通过条件查询多条
func QueryUsers(where string) ([]User, error) {
	sql := fmt.Sprintf("select id,username,password,nickname,status,createtime from users %s", where)
	logs.GetBeeLogger().Info(sql)
	rows, err := utils.QueryRows(sql)
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		user := new(User)
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Nickname, &user.Status, &user.Createtime)
		users = append(users, *user)
	}
	return users, nil
}
