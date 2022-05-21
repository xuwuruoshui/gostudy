package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-28 19:50:06
* @content:
 */

var db *sql.DB

func init() {
	err := initDB()
	if err != nil {
		panic(fmt.Sprintf("连接失败:%v", err))
	}
}
func initDB() (err error) {
	host := "root:root@tcp(192.168.0.110:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", host)
	if err != nil {
		log.Panic("连接数据库失败:", err)
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
	}

	// 第一个操作
	sqlStr1 := `insert into user1(name,age) value(?,?)`
	_, err = tx.Exec(sqlStr1, "小芳1", 12)
	if err != nil {
		tx.Rollback()
		panic("插入失败")
	}

	// 第二个操作
	sqlStr2 := `update user1 set name='en' where id=?`
	result, err := tx.Exec(sqlStr2, 4)
	if err != nil {
		tx.Rollback()
		panic("更新失败")
	}
	row, _ := result.RowsAffected()
	if row <= 0 {
		tx.Rollback()
		panic("未更新")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		panic("未知异常")
	}
}
