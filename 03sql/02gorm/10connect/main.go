package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gostudy/37gorm/10connect/model"
	"log"
	"os"
	"time"
)

var db *gorm.DB
func init(){
	// logger配置
	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:      true,         // 禁用彩色打印
		},
	)

	// 1.连接
	dsn := "root:root@tcp(192.168.0.132:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: log})
	if err!=nil{
		panic(err)
	}
	fmt.Println("连接成功")

	db.AutoMigrate(&model.Company{},&model.Employer{})
}

func create(){
	c1 := model.Company{
		Name: "Google",
	}

	e1 := model.Employer{
		Name:"zhangsan",
		Company: c1,
	}

	db.Create(&e1)
}

func main(){
	//var e1 model.Employer
	// 使用preload可以预加载
	//db.Preload("Company").First(&e1)
	//fmt.Println(e1.Name)
	//fmt.Println(e1.Company.ID)
	//fmt.Println(e1.Company.Name)

	// 连接查询, 一条sql语句
	var e2 model.Employer
	db.Joins("Company").First(&e2)
	fmt.Println(e2.Name)
	fmt.Println(e2.Company.ID)
	fmt.Println(e2.Company.Name)

}
