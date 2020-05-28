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
    
    sessionon = true
    
###三,创建并初始化数据库
    在配置的数据库中建立数据库blog
    配置正确数据库后,启动工程
    会自动初始化相关表
    相关逻辑见db.go
    
###四,讲一个坑
    项目的配置应该规矩go项目的目录规则
    demo_beego为主工程的程序包
    主工程其他的为其依赖包
    demo_beego应放在主工程的src目录下
    项目的根目录为[主工程/src/]
    
###五,引入markdown
    go get github.com/russross/blackfriday
    go get github.com/PuerkitoBio/goquery
    go get github.com/sourcegraph/syntaxhighlight

