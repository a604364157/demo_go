package models

type User struct {
	Id         int
	Username   string `orm:"size(64)"`
	Nickname   string `orm:"size(64)"`
	Password   string `orm:"size(64)"`
	Status     int
	Createtime int64
}
