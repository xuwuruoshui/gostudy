package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gostudy/37gorm/03crud/model"
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

func main(){
	// 1.插入
	//db.Create(&model.Product{Code: "A927", Price: 200})

	// 2.查询
	var p model.Product
	db.First(&p, "code=?", "A111")
	fmt.Println(p)

	// 3.更新
	// 3.1单个字段更新
	//db.Model(&p).Update("price", 300)
	//db.First("code=?","A927")
	//fmt.Println(p.Price)


	// 3.2多个字段更新,零值不更新
	//db.Model(&p).Updates(model.Product{
	//	Code:"A111",
	//	Price: 0,
	//})
	//db.First(&p,2)

	// 3.3 map的形式可以设置0值,
	// 还有一种使用3.2的方式但字段类型为 sql.NullString 或者其他的
	//db.Model(&p).Updates(map[string]interface{}{
	//	"Code":"A122",
	//	"Price": 0,
	//})
	//db.First(&p,2)
	//fmt.Println(p.Code)
	//fmt.Println(p.Price)

	// 4.逻辑删除
	var p2 model.Product
	p2.ID=1
	db.Delete(&p2,1)
}
