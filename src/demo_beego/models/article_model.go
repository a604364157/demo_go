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

/*
gt	greater的缩写，表示大于的意思	>
gte	greater than or equal的缩写，即大于等于	>=
lt	less than的缩写，表示小于	<
lte	less thanor equal的缩写，小于等于	<=
in	等同与sql语言中的in	in
exact	等于	=
contains	包含，一般用于字符类型，如包含某某字符	like '%查询内容%'
startswith	以…起始，一般用于字符类型，如从什么字符开始	like '开始字符%'
endswith	以…结束 ，一般用于字符类型，如以什么字符结束	like '%结束字符'
isnull	表示改字段不能为空

加i表示忽略大小写 入 icontains
*/

func QueryArticlePage(art Article, num, size int) []Article {
	qt := getDb().QueryTable(&art)
	if art.Id != 0 {
		qt = qt.Filter("id__exact", art.Id)
	}
	if art.Title != "" {
		qt = qt.Filter("title__exact", art.Title)
	}
	if art.Author != "" {
		qt = qt.Filter("author__exact", art.Author)
	}
	if art.Tags != "" {
		qt = qt.Filter("tags__icontains", art.Tags)
	}
	if art.Short != "" {
		qt = qt.Filter("short__exact", art.Short)
	}
	qt = qt.Limit(size, (num-1)*size)
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
