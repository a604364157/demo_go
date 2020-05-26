package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	//获取配置
	driver := beego.AppConfig.String("driver")
	userName := beego.AppConfig.String("userName")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbName := beego.AppConfig.String("dbName")
	//注册驱动(mysql)
	err := orm.RegisterDriver(driver, orm.DRMySQL)
	//dbConn := "user:pwd@tcp(host:port)/db?charset=utf8"
	dbConn := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"
	err = orm.RegisterDataBase("default", driver, dbConn, 30)
	if err != nil {
		logs.GetBeeLogger().Error("连接数据库时,出错了!")
		return
	}
	logs.GetBeeLogger().Info("初始化数据库成功")
}
