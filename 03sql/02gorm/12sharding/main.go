package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gostudy/37gorm/12sharding/model"
	"log"
	"os"
	"strconv"
	"time"
)

var db *gorm.DB

func init() {
	// logger配置
	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	// 1.连接
	dsn := "root:123456@tcp(192.168.2.39:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: log})
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功")

	db.AutoMigrate(&model.User{})
}

func main() {
	for i := 0; i < 10; i++ {
		u := uuid.New().ID()
		id := strconv.Itoa(int(u))
		user := model.User{Id: id, Name: "hahah" + strconv.Itoa(i)}
		data, _ := strconv.Atoi(id[len(id)-1:])

		tx := db.Scopes(model.TableOfIdLastYear("user", data, time.Now())).Create(&user)
		if tx.RowsAffected == 0 {
			fmt.Println("failed")
		}
	}

}
