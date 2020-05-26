package utils

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func Init() {
	if db != nil {
		return
	}
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

	//err := orm.RegisterDataBase("default", driver, dbConn)
	db1, err := sql.Open(driver, dbConn)
	if err != nil {
		log.Error("连接数据库时,出错了!")
		return
	}
	log.Info("初始化数据库成功")
	db = db1
}

func InitTable() {
	//创建user表
	CreateTableWithUser()
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		nickname VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	Init()
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//查询单行
func QueryRow(sql string) *sql.Row {
	Init()
	return db.QueryRow(sql)
}

func QueryRows(sql string) (*sql.Rows, error) {
	Init()
	return db.Query(sql)
}
