package main

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
	"goStudy/database"
	"time"
)

func testTime() {
	now := time.Now()
	fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	sencond := now.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, sencond)

	timestamp := now.Unix()
	fmt.Printf("timestamp is %d\n", timestamp)
}

func testTimeStamp(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	sencond := timeObj.Second()
	fmt.Printf("current timestamp is %d\n", timestamp)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, sencond)
}

func testTicker() {
	ticker := time.Tick(time.Second * 10)
	for range ticker {
		testTimeStamp(time.Now().Unix())
	}

}

func testFormat() {
	now := time.Now()
	format := now.Format("2006-01-02 15:04:05")
	fmt.Println(format)

	format2 := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(format2)

}

type User struct {
	Id     int64
	Name   string
	Age    int
	Height int
}

var o orm.Ormer

func main() {
	//testTime()
	//testTimeStamp(time.Now().Unix())
	//testTicker()
	//testFormat()

	dns := "root:12345678@tcp(localhost:3306)/test?charset=utf8&parseTime=true&loc=Local"
	orm.RegisterDataBase("default", "mysql", dns)

	o = orm.NewOrm()
	o.Using("default")

	var users []User
	o.Raw("select * from user").QueryRows(&users)

	fmt.Printf("%v\n", users)

	database.TestOrm()
}
