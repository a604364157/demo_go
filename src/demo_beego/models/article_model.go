package models

import "github.com/astaxie/beego/logs"

type Article struct {
	Id         int
	Title      string `orm:"size(30)"`
	Author     string `orm:"size(20)"`
	Tags       string `orm:"size(30)"`
	Short      string `orm:"size(255)"`
	Content    string `orm:"size(5000)"`
	CreateTime int64
}

func InsertArticle(article Article) error {
	_, err := getDb().Insert(&article)
	if err != nil {
		logs.GetBeeLogger().Info("写入文章数据失败")
	}
	return err
}
