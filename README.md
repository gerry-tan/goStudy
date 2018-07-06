# Beego with Swagger UI

## 安装
* 安装 Go 和 Git 环境
* 执行以下命令安装 beego 和 bee 工具：\
  $ go get -u github.com/astaxie/beego \
  $ go get -u github.com/beego/bee

## 创建项目
* 通过jdbc创建Restful风格API项目 \
  $ go get github.com/Go-SQL-Driver/MySQL \
  $ go install github.com/Go-SQL-Driver/MySQL \
  $ bee api test -conn="root:123456@tcp(127.0.0.1:3306)/test"
* API文档自动化 \
  $ cd $GOPATH/test \
  $ bee generate docs 
* swagger 安装\
  $ bee run test -gendoc=true -downdoc=true

## 配置
* conf/app.conf
  * 访问端口配置\
    httpport = 8080
  * 开启进程内监控\
    EnableAdmin = true\
    AdminHttpAddr = 0.0.0.0\
    AdminHttpPort = 8088

* routers/router.go 
  * 修改访问路径 \
    ns := beego.NewNamespace("/api", \
        beego.NSNamespace("/test", \
			beego.NSInclude( \
				&controllers.TestController{}, \
			), \
		),\
	)

## 启动程序
* $ bee run beegoTest -gendoc=true

## Web页面
* swagger页面：
  http://localhost:8080/swagger/
* manager页面：
  http://localhost:8088/qps
* API访问地址：
  http://localhost:8080/v1/test
