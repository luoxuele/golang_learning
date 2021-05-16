# 1. Go 语言的源码复用建立在包（package）基础之上。包通过 package, import, GOPATH 操作完成。
    1. 一个代码包可以看成一个目录
    2. 包可以嵌套（子目录）
    3. 包里面的所有.go文件都






# 1. 在Go1.5之前使用GOROOT和GOPATH这2个系统环境变量来决定包的位置，对于开发者主要使用GOPATH

    $GOPATH/src 第三方包
    $GOROOT/src 标准库的包 如fmt,math

# 2. 第三方包
    https://godoc.org
    目前多数包来自 Github
    官方包来自 golang.org/x/...

    github.com/astaxie/beego 成熟稳定的Web框架，包含更过的Web Framework特性 国内开发者 astaxie
    github.com/go-redis/redis 连接Redis
    github.com/gomodule/redigo/redis 连接Redis
    github.com/jinzhu/gorm 数据库ORM框架，类似Java领域的Hibernate或MyBatis 国内开发者
    github.com/sirupsen/logrus 日志框架，类似Java领域的log4j
    github.com/robfig/cron 定时任务，类似Java领域的Quartz
    github.com/gin-gonic/gin 精巧高效的Web框架，多用于提供Web服务


