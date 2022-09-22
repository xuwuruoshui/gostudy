package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

func main(){
	// 1.连接
	dsn := "root:root@tcp(192.168.0.132:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	fmt.Println("连接成功")


}
