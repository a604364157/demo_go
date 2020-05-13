package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
)

type Client struct {
	conn net.Conn
	name string
	addr string
}

var (
	clientsMap = make(map[string]Client)
)

func ShandleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

func main() {
	//建立服务端监听
	listener, e := net.Listen("tcp", "127.0.0.1:8888")
	ShandleError(e, "创建链接失败")
	defer func() {
		for _, client := range clientsMap {
			_, _ = client.conn.Write([]byte("all：服务器维护了，大家洗洗睡吧！"))
		}
		_ = listener.Close()
	}()

	for {
		//循环接入所有链接
		conn, e := listener.Accept()
		ShandleError(e, "listener.Accept")
		clientAddr := conn.RemoteAddr()
		//接收并保存昵称
		buffer := make([]byte, 1024)
		var clientName string
		for {
			n, err := conn.Read(buffer)
			ShandleError(err, "conn.Read(buffer)")
			if n > 0 {
				clientName = string(buffer[:n])
				break
			}
		}
		fmt.Println(clientName + "上线了")

		//将每一个用户存入map
		client := Client{conn, clientName, clientAddr.String()}
		clientsMap[clientName] = client

		//给在线用户发送上线通知
		for _, client := range clientsMap {
			_, _ = client.conn.Write([]byte(clientName + "上线了"))
		}

		//在单独的线程中于用户聊天
		go ioWithClient(client)
	}
}

func ioWithClient(client Client) {
	buffer := make([]byte, 1024)
	for {
		n, err := client.conn.Read(buffer)
		if err != io.EOF {
			ShandleError(err, "conn.Read")
		}
		if n > 0 {
			msg := string(buffer[:n])
			fmt.Printf("%s:%s\n", client.name, msg)
			writeMsgToLog(msg, client)

			strs := strings.Split(msg, "#")
			if len(strs) > 1 {
				//要发送的目标昵称
				targetName := strs[0]
				targetMsg := strs[1]
				//使用昵称定位客户端的conn
				if targetName == "all" {
					//群体发送
					for _, c := range clientsMap {
						_, _ = c.conn.Write([]byte(client.name + ":" + targetMsg))
					}
				} else {
					//单独消息
					for key, c := range clientsMap {
						if key == targetName {
							_, _ = c.conn.Write([]byte(client.name + ":" + targetMsg))
							//日志
							go writeMsgToLog(client.name+":"+targetMsg, c)
							break
						}
					}
				}
			} else {
				//客户端断开
				if msg == "exit" {
					//从在线用户移出并发送下线通知
					for name, c := range clientsMap {
						if c == client {
							delete(clientsMap, name)
						} else {
							_, _ = c.conn.Write([]byte(name + "下线了"))
						}
					}
				} else if strings.Index(msg, "log@") == 0 {
					filterName := strings.Split(msg, "@")[1]
					//向客户端发送聊天日志
					go sendLogToClient(client, filterName)
				} else {
					_, _ = client.conn.Write([]byte("已读：" + msg))
				}
			}
		}
	}
}

func sendLogToClient(client Client, filterName string) {
	//读取日志
	logBytes, e := ioutil.ReadFile("E:\\logs\\" + client.name + ".log")
	ShandleError(e, "ioutil.ReadFile")

	if filterName != "all" {
		//个人聊天记录
		logStr := string(logBytes)
		targetStr := ""
		lineSlice := strings.Split(logStr, "\n")
		for _, lineStr := range lineSlice {
			if len(lineStr) > 20 {
				contentStr := lineStr[20:]
				if strings.Index(contentStr, filterName+"#") == 0 || strings.Index(contentStr, filterName+":") == 0 {
					targetStr += lineStr + "\n"
				}
			}
		}
		_, _ = client.conn.Write([]byte(targetStr))
	} else {
		_, _ = client.conn.Write(logBytes)
	}
}

func writeMsgToLog(msg string, client Client) {
	file, e := os.OpenFile("E:\\logs\\"+client.name+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	ShandleError(e, "os.OpenFile")
	defer func() {
		_ = file.Close()
	}()
	logMsg := fmt.Sprintln(time.Now().Format("2006-01-02 15:04:05"), msg+"\n")
	_, _ = file.Write([]byte(logMsg))
}
