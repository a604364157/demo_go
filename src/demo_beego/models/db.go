package models

import (
	"demo_beego/commom/constants"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//go语言只要加载文件,就会执行文件的init函数
func init() {
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
	err = orm.RegisterDataBase(constants.DEFAULT, driver, dbConn, 30)
	if err != nil {
		logs.GetBeeLogger().Error("连接数据库时,出错了!")
		return
	}
	RegisterModels()
	orm.RunSyncdb(constants.DEFAULT, false, true)
	logs.GetBeeLogger().Info("初始化数据库成功")
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		orm.Debug = true
	}
}

//注册model对象
func RegisterModels() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Article))
	orm.RegisterModel(new(Album))
}

func getDb() orm.Ormer {
	return orm.NewOrm()
}
