package models

/*
MySQL： github.com/go-sql-driver/mysql
PostgreSQL：github.com/lib/pq
Sqlite3：github.com/mattn/go-sqlite3
*/

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	log := logs.GetBeeLogger()
	//获取配置
	driver := beego.AppConfig.String("driver")
	userName := beego.AppConfig.String("userName")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbName := beego.AppConfig.String("dbName")
	//注册驱动(mysql)
	orm.RegisterDriver(driver, orm.DRMySQL)
	//dbConn := "user:pwd@tcp(host:port)/db?charset=utf8"
	dbConn := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"

	err := orm.RegisterDataBase("default", driver, dbConn)
	if err != nil {
		log.Error("连接数据库时,出错了!")
		return
	}
	log.Info("初始化数据库成功")
}
