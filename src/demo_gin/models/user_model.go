package models

import "fmt"

type User struct {
	Id           int     `xorm:"pk autoincr" json:"id"`
	UserName     string  `xorm:"varchar(20)" json:"user_name"`
	Password     string  `xorm:"varchar(255)" json:"password"`
	Mobile       string  `xorm:"varchar(11)" json:"mobile"`
	Avatar       string  `xorm:"varchar(255)" json:"avatar"`
	Balance      float64 `xorm:"double" json:"balance"`
	IsActive     int8    `xorm:"tinyint" json:"is_active"`
	City         string  `xorm:"varchar(10)" json:"city"`
	RegisterTime int64   `xorm:"bigint" json:"register_time"`
}

func QueryUserById(id int) User {
	user := new(User)
	_, err := getDb().ID(id).Get(user)
	if err != nil {
		fmt.Println("查询用户信息失败")
	}
	return *user
}

func QueryUser(user User) map[int64]User {
	session := getDb().Cols("id", "user_name", "password", "mobile", "avatar", "balance", "is_active", "city", "register_time")
	if user.Id != 0 {
		session = session.Where("id = ?", user.Id)
	}
	if user.UserName != "" {
		session = session.Where("user_name = ?", user.UserName)
	}
	if user.Password != "" {
		session = session.Where("password = ?", user.Password)
	}
	if user.Mobile != "" {
		session = session.Where("mobile = ?", user.Mobile)
	}
	if user.Avatar != "" {
		session = session.Where("avatar = ?", user.Avatar)
	}
	if user.Balance != 0 {
		session = session.Where("balance = ?", user.Balance)
	}
	if user.IsActive != 0 {
		session = session.Where("is_active = ?", user.IsActive)
	}
	if user.City != "" {
		session = session.Where("city = ?", user.City)
	}
	users := make(map[int64]User)
	err := session.Find(&user)
	if err != nil {
		fmt.Println("查询用户信息失败")
	}
	return users
}
