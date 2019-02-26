package database

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     int64
	Name   string
	Age    int
	Height int
}

func TestOrm() {
	o := orm.NewOrm()

	var users []User
	o.Raw("select * from user").QueryRows(&users)

	fmt.Printf("users: %v\n", users)
}
