package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitTable() {
	driver := beego.AppConfig.String("driver")
	userName := beego.AppConfig.String("userName")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbName := beego.AppConfig.String("dbName")
	dbConn := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"
	db1, err := sql.Open(driver, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
	}
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
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

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
