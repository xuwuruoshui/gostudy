package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gostudy/37gorm/11connect2/model"
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

	db.AutoMigrate(&model.Company{},&model.Employer{},&model.Creditcard{})
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


	// 一对多
	// 方法1
	//var e1 model.Employer
	//db.First(&e1,1)
	//
	//for i := 0; i < 5; i++ {
	//	u := uuid.New()
	//	c := model.Creditcard{
	//		Number: u.String(),
	//		EmployerId: e1.ID,
	//	}
	//	db.Create(&c)
	//}
	//
	//// 预先查询某个字段
	//db.Preload("Creditcards").First(&e1)
	//for _, creditcard := range e1.Creditcards {
	//	fmt.Println(creditcard.Number)
	//}

	var e2 model.Employer
	db.First(&e2,1)
	var creditCards []*model.Creditcard
	db.Model(&e2).Association("Creditcards").Find(&creditCards)
	for _, creditcard := range creditCards {
		fmt.Println(creditcard.Number)
	}
}
