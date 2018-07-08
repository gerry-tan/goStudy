package main

import (
	_ "test/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 设置最大空闲连接
	maxIdle, err := beego.AppConfig.Int("mysql.maxIdle")
	if err != nil {
		println(err)
	}
	// 设置最大数据库连接 (go >= 1.2)
	maxConn, error2 := beego.AppConfig.Int("mysql.maxConn")
	if error2 != nil {
		println(error2)
	}
	datasource := beego.AppConfig.String("mysql.default.datasource")
	orm.RegisterDataBase("default", "mysql", datasource, maxIdle, maxConn)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	o := orm.NewOrm()
	o.Using("default") //使用default数据库

	beego.Run()
}

