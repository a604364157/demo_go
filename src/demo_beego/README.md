# demo_beego

#项目初始化配置

###一,依赖安装
    beego:  go get github.com/astaxie/beego
    
    bee:    go get github.com/beego/bee
    
    MySQL： go get github.com/go-sql-driver/mysql
    
###二,创建配置文件
    conf/app.conf
    
    appname = demo_beego
    httpport = 8080
    runmode = dev
    
    #mysql配置
    driver = mysql
    userName = xxx
    password = xxx
    host = xxx
    port = xxx
    dbname = blog
    
###三,创建并初始化数据库
    在配置的数据库中建立数据库blog
    工程已提供utils/util.go工具类
    运行InitTable函数,即可初始化相关表和初始数据
    后续开发会逐渐补充初始化脚本
    


