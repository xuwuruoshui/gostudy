package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-28 21:22:49
* @content: sqlx包,取代database/sql
 */

var db *sqlx.DB

func init() {
	initDB()
}

func initDB() {
	dsn := "root:root@tcp(192.168.0.110:3306)/sql_test?charset=utf8mb4&parseTime=True"

	// MustConnet连接不成功旧panic,而Connect只会返回error
	db = sqlx.MustConnect("mysql", dsn)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
}

type user struct {
	Id   int
	Age  int
	Name string
}


// 查询
func queryRowDemo(){
	sqlStr := "select id,name,age from user1 where id = ?"
	var u user

	err := db.Get(&u,sqlStr,1)
	if err!=nil{
		panic(err)
	}

	fmt.Println(u)
}

func queryRowList(){
	sqlStr := "select id,name,age from user1"
	var uList []user

	// 这里居然还是要指针
	err := db.Select(&uList,sqlStr)
	if err!=nil{
		panic(err)
	}

	fmt.Println(uList)
}

func main() {
	//queryRowDemo()
	queryRowList()
}
