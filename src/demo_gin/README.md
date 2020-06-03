# demo_beego

#项目初始化配置

###一,依赖安装
    gin go get -u github.com/gin-gonic/gin
    
    xormplus go get -u github.com/xormplus/xorm(是xorm的增强,个人感觉比xorm好用一点)
    
    gomail go get -u gopkg.in/gomail.v2
    
    因为网络问题,很多依赖down不下来
    直接运行main函数,会出现很多找不到包
    然后根据提示路径,下载依赖包到对应路径即可
    

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
