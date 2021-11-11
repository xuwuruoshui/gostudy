package service

import (
	"database/sql"
	"gostudy/25mysql/04practice/logreg/entity"
	"gostudy/25mysql/04practice/logreg/vo"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	host := "root:root@tcp(192.168.0.110:3306)/sql_test?charset=utf8mb4&parseTime=True"
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

func FetchOne(id int) entity.User {
	initDB()

	sqlStr := "select id,username,age from user where id = ?"
	stmt, err := db.Prepare(sqlStr)

	if err != nil {
		log.Panic("预处理失败:", err)
	}

	defer stmt.Close()

	var u entity.User
	err = stmt.QueryRow(id).Scan(&u.Id, &u.Username, &u.Age)
	if err != nil {
		log.Panic("查询失败:", err)
	}

	return u
}

func FetchOneByUsernamePasswd(u entity.User) (vo.UserVo, bool) {
	initDB()

	sqlStr := "select id,username from user where username=? and password=?"
	stmt, err := db.Prepare(sqlStr)

	if err != nil {
		log.Panic("预处理失败:", err)
		return vo.UserVo{}, false
	}

	var uservo vo.UserVo
	err = stmt.QueryRow(u.Username, u.Password).Scan(&uservo.Id, &uservo.Username)
	if err != nil {
		log.Println("查询失败:", err)
		return vo.UserVo{}, false
	}

	return uservo, true
}

func AddUser(u entity.User) (int64,bool) {
	_, ok := FetchOneByUsernamePasswd(u)
	if ok {
		log.Println("已经存在用户名")
		return 0,false
	}

	initDB()
	sqlStr := "insert into user(username,age,password) values(?,?,?)"
	stmt, err := db.Prepare(sqlStr)

	if err != nil {
		log.Panic("预处理失败:", err)
	}

	result, err := stmt.Exec(&u.Username, &u.Age, &u.Password)
	if err != nil{
		log.Println("插入失败:", err)
		return 0,false
	}

	id,_ := result.LastInsertId()

	return id,true

}
