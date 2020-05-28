package utils

import (
	"bytes"
	"crypto/md5"
	"demo_beego/commom/constants"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
	"io"
	"time"
)

const SOLT = "demo_beego"

func MD5(str string) string {
	w := md5.New()
	io.WriteString(w, str+SOLT)
	md5Srt := fmt.Sprintf("%x", w.Sum(nil))
	return md5Srt
}

func TimeStampToData(timesTamp int64) string {
	datetime := time.Unix(timesTamp, 0).Format(constants.FormatDate)
	return datetime
}

func MarkdownToHtml(content string) template.HTML {
	markdown := blackfriday.MarkdownCommon([]byte(content))
	//获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	/**
	  对document进程查询，选择器和css的语法一样
	  第一个参数：i是查询到的第几个元素
	  第二个参数：selection就是查询到的元素
	*/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}
