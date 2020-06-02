package controllers

import (
	"demo_beego/models"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UploadController struct {
	BaseController
}

func (c *UploadController) Post() {
	file, header, err := c.GetFile("upload")
	if err != nil {
		c.responseErr(err)
		return
	}
	now := time.Now()
	fileType := "unKnow"
	//限制文件类型为图片
	fileExt := filepath.Ext(header.Filename)
	fileExt = strings.ToLower(fileExt)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" || fileExt == ".bmp" {
		fileType = "img"
	}
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//创建目录,设置权限777
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		c.responseErr(err)
		return
	}
	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, header.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		c.responseErr(err)
		return
	}
	//拷贝文件
	_, err = io.Copy(desFile, file)
	if err != nil {
		c.responseErr(err)
		return
	}
	if fileType == "img" {
		album := models.Album{FilePath: filePathStr, FileName: fileName, CreateTime: timeStamp}
		models.InsertAlbum(album)
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "message": "上传成功"}
	c.ServeJSON()
}

func (c *UploadController) responseErr(err error) {
	c.Data["json"] = map[string]interface{}{"code": 1, "message": err}
	c.ServeJSON()
}
