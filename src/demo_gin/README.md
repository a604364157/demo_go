# demo_beego

#项目初始化配置

###一,依赖安装
    gin go get -u github.com/gin-gonic/gin
    
    xormplus go get -u github.com/xormplus/xorm(是xorm的增强,个人感觉比xorm好用一点)
    
    gomail go get -u gopkg.in/gomail.v2
    
    直接使用代理https://goproxy.io
    安装1.13版本以上的golang环境,开启module和代理
    设置环境变量
    GO111MODULE=on
    GOPROXY=https://goproxy.io
    开发工具也需要在设置里开启module并设置代理
    使用 go mod init [项目目录名]
    初始化module后,项目下会出现一个go.mod文件
    该文件是用来管理项目依赖的
    请按该教程设置,本项目不再说明使用的依赖
    

###二,配置文件config/app.json(数据同步至config/config.go内的结构体,请自行添加相关配置)
```json
{
  "app_name": "demo_gin",
  "app_mode": "debug",
  "app_host": "localhost",
  "app_port": "8080",
  "mail_user": "xxxxxxx@163.com",
  "mail_password": "xxxxxxxxxxxxxx",
  "mail_host": "smtp.163.com",
  "mail_port": "465"
}
```
