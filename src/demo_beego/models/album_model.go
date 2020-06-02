package models

import "github.com/astaxie/beego/logs"

type Album struct {
	Id         int
	FilePath   string `orm:"size(255)"`
	FileName   string `orm:"size(64)"`
	Status     int
	CreateTime int64
}

func QueryAllAlbums() []Album {
	album := new(Album)
	var albums []Album
	_, err := getDb().QueryTable(album).All(&albums)
	if err != nil {
		logs.GetBeeLogger().Info("查询文件数据失败")
	}
	return albums
}

func InsertAlbum(album Album) error {
	_, err := getDb().Insert(&album)
	if err != nil {
		logs.GetBeeLogger().Info("写入文件数据失败")
	}
	return err
}
