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
	tcp := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbDb)
	var err interface{}
	engine, err = xorm.NewEngine(cfg.DbDriver, tcp)
	if err != nil {
		panic(err)
	}
}

//连接池也是DB操作对象
func GetDb() *xorm.Engine {
	return engine
}

//连接,用于带事务的操作,需手动关闭
func GetSession() *xorm.Session {
	return engine.NewSession()
}
