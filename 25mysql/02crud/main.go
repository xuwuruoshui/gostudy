package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-10 09:34:31
* @content:
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

func fetchOne() {
	sqlStr := "select id,age,name from user where id = ?"
	var u user
	err := db.QueryRow(sqlStr,1).Scan(&u.id,&u.age,&u.name)
	if err!=nil{
		log.Println("数据库查询失败",err)
	}
	log.Println(u.id,u.age,u.name)
}

func fetchList(){
	sqlStr := "select id,age,name from user where id > ?"
	rows,err := db.Query(sqlStr,0)
	if err!=nil{
		log.Println("查询失败:",err)
	}


	// 遍历结果集
	for rows.Next(){
		var u user
		err := rows.Scan(&u.id,&u.age,&u.name)
		if err!=nil{
			log.Println("结果集赋值失败",err)
		}
		log.Println(u.id,u.age,u.name)
	}
}

func insert(){
	sqlStr := "insert into user(name,age) values(?,?)"
	result,err := db.Exec(sqlStr,"小明",10)
	if err!=nil{
		log.Println("插入失败:",err)
	}
	count,_ := result.RowsAffected()
	log.Println("影响行数:",count)

	id,_ := result.LastInsertId()
	log.Println("反显Id:",id)
}

func update(){
	sqlStr := "update user set name=?,age=? where id = ?"
	rows,err := db.Exec(sqlStr,"小芳",20,4)
	if err!=nil{
		log.Println("更新失败:",err)
	}

	count,_ := rows.RowsAffected()
	log.Println("影响行数:",count)
}

func delete(){
	sqlStr := "delete from user where id = ?"
	rows,err := db.Exec(sqlStr,4)
	if err !=nil{
		log.Println("删除失败",err)
	}

	count,_ := rows.RowsAffected()
	log.Println("影响行数",count)
}

func main(){
	initDB()

	// 1.单条查询
	//fetchOne()

	// 2.多条查询
	//fetchList()

	// 3.插入数据
	//insert()
	
	// 4.更新数据
	//update()

	// 5.删除数据
	delete()
}



