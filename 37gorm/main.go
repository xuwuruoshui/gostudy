package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

func main(){
	dsn := "root:root@tcp(192.168.0.110:3306)/sql_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	db.AutoMigrate(&User{})
	u := User{2,"haha","女","唱歌"}
	db.Create(&u)
	result1 := db.Create(&u)
	if result1.RowsAffected>0{
		log.Print("haha")
	}
}
