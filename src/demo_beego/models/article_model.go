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

func CountArticle() int {
	count, err := getDb().QueryTable(new(Article)).Count()
	if err != nil {
		logs.GetBeeLogger().Info("查询失败")
	}
	return int(count)
}

func QueryArticlePage(art Article, num, size int) []Article {
	qt := getDb().QueryTable(&art)
	if art.Id != 0 {
		qt.Filter("id", art.Id)
	}
	if art.Title != "" {
		qt.Filter("title", art.Title)
	}
	if art.Author != "" {
		qt.Filter("author", art.Author)
	}
	if art.Tags != "" {
		qt.Filter("tags", art.Tags)
	}
	if art.Short != "" {
		qt.Filter("short", art.Short)
	}
	qt.Limit((num-1)*size, num*size)
	var arts []Article
	_, err := qt.All(&arts)
	if err != nil {
		logs.GetBeeLogger().Info("查询失败")
	}
	return arts
}

func QueryArticleById(id int) Article {
	art := new(Article)
	art.Id = id
	err := getDb().Read(art)
	if err != nil {
		logs.GetBeeLogger().Info("查询失败")
	}
	return *art
}

func InsertArticle(article Article) error {
	_, err := getDb().Insert(&article)
	if err != nil {
		logs.GetBeeLogger().Info("写入文章数据失败")
	}
	return err
}

func UpdateArticle(article Article) error {
	_, err := getDb().Update(&article)
	if err != nil {
		logs.GetBeeLogger().Info("更新文章数据失败")
	}
	return err
}

func DeleteArticle(id int) bool {
	art := new(Article)
	art.Id = id
	num, err := getDb().Delete(art)
	if err != nil {
		logs.GetBeeLogger().Info("删除文章数据失败")
	}
	return num > 0
}
