package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gostudy/37gorm/08update/model"
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

}

func add(){
	var products []*model.Product
	p1 := &model.Product{
		Price: 100,
		Code: "aaa",
	}

	p2 := &model.Product{
		Price: 200,
		Code: "fff",
	}
	products = append(products,p1,p2)

	db.CreateInBatches(products,2)
}


func main(){
	// add()

	// 删除
	//var p model.Product
	//db.First(&p,1)
	//db.Delete(&p)

	//db.Where("price",200).Delete(&model.Product{})

	// 查询被软删除的数据
	//var ps []model.Product
	//db.Unscoped().Where("price",200).Find(&ps)
	//fmt.Println(ps)

	// 永久删除
	//db.Unscoped().Delete()
}
