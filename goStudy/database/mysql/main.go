package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

type UserInfo struct {
	Id   int64          `db:"id"`
	Name sql.NullString `db:"name"`
	Age  int            `db:"age"`
}

var db *sql.DB

func init() {
	mysqlUrl := "root:Cango@1219@tcp(10.42.0.121:3306)/test?charset=utf8&loc=Asia%2FShanghai"
	var e error
	db, e = sql.Open("mysql", mysqlUrl)
	if e != nil {
		fmt.Println("connect mysql db failed, ", e)
	}
}

func queryRow() {
	sqlStr := "select * from userInfo where id = ?"
	row := db.QueryRow(sqlStr, 2)

	var user UserInfo
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("id: %v, name: %v, age: %v\n", user.Id, user.Name, user.Age)
}

func queryMultiRow() {
	sqlStr := "select * from userInfo where id > ?"
	rows, _ := db.Query(sqlStr, 0)

	//重点
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	for rows.Next() {
		var user UserInfo
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("id: %v, name: %v, age: %v\n", user.Id, user.Name, user.Age)
	}
}

func insertRow() {
	sqlStr := "insert into userInfo(name, age) values(?, ?)"
	result, e := db.Exec(sqlStr, "Jame", 34)
	if e != nil {
		fmt.Println(e)
		return
	}

	lastInsertId, _ := result.LastInsertId()
	fmt.Println("lastInsertId is: ", lastInsertId)
}

func prepareQuery() {
	sqlStr := "select * from userInfo where id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	row := stmt.QueryRow(1)
	var user UserInfo

	_ = row.Scan(&user.Id, &user.Name, &user.Age)

	fmt.Printf("id: %v, name: %v, age: %v\n", user.Id, user.Name, user.Age)
}

func trans() {
	//事物操作
	tx, e := db.Begin()
	if e != nil {
		fmt.Println(e)
	}

	sqlStr := "update userInfo set age = 18 where id = ?"
	_, e = tx.Exec(sqlStr, 1)
	if e != nil {
		tx.Rollback()
		fmt.Println(e)
		return
	}

	sqlStr = "update userInfo set ages = 43 where id = ?"
	_, e = tx.Exec(sqlStr, 10)
	if e != nil {
		tx.Rollback()
		fmt.Println(e)
		return
	}

	tx.Commit()

}

func main() {
	//queryRow()

	//queryMultiRow()

	//insertRow()

	//prepareQuery()

	trans()
	db.Close()
}
