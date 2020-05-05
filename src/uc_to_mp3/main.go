package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	cachePath := "C:\\\\Users\\Administrator\\AppData\\Local\\Netease\\CloudMusic\\Cache\\Cache"
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("没有传入相应参数,第一个参数为:目标目录,第二个参数为:生成MP3路径")
	}
	cachePath = args[0]
	targetPath := args[1]
	fmt.Printf("cachePath:%s  targetPath:%s\n", cachePath, targetPath)
	if err := os.MkdirAll(targetPath, 0777); err != nil {
		log.Fatal(err)
	}
	files, _ := ioutil.ReadDir(cachePath)
	ch := make(chan string, 20)
	startTime := time.Now()
	count := 0
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".uc") {
			ch <- file.Name()
			go func(fileName string) {
				fmt.Println("parse:", fileName)
				err := decodeFile(cachePath, fileName, targetPath, fileName+".mp3")
				count++
				defer func() {
					fmt.Println("parse end:", <-ch, "  err:", err)
				}()
			}(file.Name())
		}
	}
	fmt.Println("all end", time.Since(startTime), "count:", count)
}

func decodeFile(dirName, fileName, newDirName, newName string) error {
	bytes, err := ioutil.ReadFile(dirName + "\\" + fileName)
	if err != nil {
		return err
	}
	for i := 0; i < len(bytes); i++ {
		bytes[i] ^= 0xa3
	}
	return ioutil.WriteFile(newDirName+"\\"+newName, bytes, 0777)
}
