package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-09 09:22:39
* @content:
 */

var db *sql.DB

func initDB() (err error) {
	host := "root:root@tcp(192.168.2.152:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", host)
	if err != nil {
		log.Panic("连接数据库失败:", err)
	}

	// 数据库最大连接数目,0<n<SetMaxIdleConns: SetMaxIdleConns=SetMaxOpenConns,n<=0: SetMaxIdleConns不受限制
	db.SetMaxOpenConns(20)
	// 数据库最大闲置数目
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	sqlStr := "select id,name,age from user where id=?"
	var u user
	err = db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println(u.id, u.name, u.age)

}

type user struct {
	id   int
	name string
	age  int
}
