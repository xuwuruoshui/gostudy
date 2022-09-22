package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gostudy/37gorm/02generate/model"
	"log"
	"os"
	"time"
)

// logger配置,自动迁移

func main(){

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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: log})
	if err!=nil{
		panic(err)
	}
	fmt.Println("连接成功")

	// 2.生成表
	err = db.AutoMigrate(&model.Product{})
	if err!=nil{
		panic(err)
	}
}
