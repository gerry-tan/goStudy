package main

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jmoiron/sqlx"
)

type UserInfo struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

var db *sqlx.DB

func init() {
	mysqlUrl := "root:Cango@1219@tcp(10.42.0.121:3306)/test?charset=utf8&loc=Asia%2FShanghai"
	var e error
	db, e = sqlx.Open("mysql", mysqlUrl)
	if e != nil {
		fmt.Printf("connect mysql db failed, error: %v", e)
	}
}

func queryRow() {
	sqlStr := "select * from userInfo where id = ?"

	var user UserInfo
	_ = db.Get(&user, sqlStr, 1)

	fmt.Printf("user: %#v", user)
}

func queryMultiRow() {
	sqlStr := "select * from userInfo where id > ?"

	var users []UserInfo

	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("users: %v\n", users)
}

func updateMulti() {
	sqlStr := "update userInfo set age=? where id = ?"
	result, err := db.Exec(sqlStr, 99, 2)
	if err != nil {
		fmt.Printf("update failed, error: %v\n", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("affected rows failed, error: %v\n", err)
	}

	fmt.Printf("affect rows: %v\n", count)

}

func main() {
	//queryRow()

	//queryMultiRow()

	updateMulti()

	db.Close()
}
