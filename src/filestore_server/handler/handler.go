package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传的文件
		data, err := ioutil.ReadFile("E:\\Workspace\\demo_go\\src\\filestore_server\\static\\view\\index.html")
		if err != nil {
			_, _ = io.WriteString(w, err.Error())
			return
		}
		_, _ = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收文件并存储
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data, err:%s\n", err.Error())
			return
		}
		defer file.Close()

		newFile, err := os.Create("\\tmp\\" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create file, err:%s\n", err.Error())
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file, err:%s\n", err.Error())
			return
		}
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Upload finished!")
}
