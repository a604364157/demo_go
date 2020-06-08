package models

import (
	"demo_gin/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func init() {
	cfg := config.GetConfig()
	tcp := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbDb)
	var err interface{}
	engine, err = xorm.NewEngine(cfg.DbDriver, tcp)
	engine.ShowSQL(true)
	if err != nil {
		panic(err)
	}
	syncTable()
}

//注册模型,同步生成表
func syncTable() {
	engine.Sync2(new(SmsCode))
	engine.Sync2(new(User))
}

//连接池也是DB操作对象
func getDb() *xorm.Engine {
	return engine
}

//连接,用于带事务的操作,需手动关闭
func getSession() *xorm.Session {
	return engine.NewSession()
}
