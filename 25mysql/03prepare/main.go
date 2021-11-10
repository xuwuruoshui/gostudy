package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-10 10:11:15
* @content: 预处理
* 优势:1.防止SQL注入
*			2. 优化服务器重复执行SQL,使服务器提前编译,一次编译多次执行
*/

type user struct{
	id int
	age int
	name string
}

var db *sql.DB

func initDB() (err error){
	host := "root:root@tcp(192.168.2.152:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db,err = sql.Open("mysql",host)
	if err!=nil{
		log.Panic("连接数据库失败",err)
	}

	err = db.Ping()
	if err!=nil{
		log.Panic("连接数据库失败",err)
	}
	return err
}

func selectPrepare(){
	sqlStr := "select id,name,age from user where id>?"
	stmt,err := db.Prepare(sqlStr)
	if err!=nil{
		log.Println("预处理失败")
	}
	defer stmt.Close()

	rows,err := stmt.Query(0)
	if err!=nil{
		log.Println("查询失败:",err)
	}

	for rows.Next(){
		var u user
		err = rows.Scan(&u.id,&u.name,&u.age)
		if err!=nil{
			log.Println("赋值失败:",err)
		}
		log.Println(u.id,u.age,u.age)
	}
}

func insertPrepare(){
	sqlStr := "insert into user(name,age) value(?,?)"
	stmt,err := db.Prepare(sqlStr)
	if err!=nil{
		log.Println("预编译失败:",err)
	}
	defer stmt.Close()

	result,err := stmt.Exec("hah",12)
	if err!=nil{
		log.Println("插入失败:",err)
	}

	count,_:=result.RowsAffected()
	log.Println("影响行数:",count)

	id,_ := result.LastInsertId()
	log.Println("id:",id)
}

func main(){
	initDB()

	// 1.查询预处理
	selectPrepare()

	// 2.插入预处理
	insertPrepare()
}

