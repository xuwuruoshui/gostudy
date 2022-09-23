package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	Id uint
	Name string
	Email *string
	Age uint8
	Birthday *time.Time
	MemberNumber sql.NullString
	ActivatedAt sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

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
	users := []User{
		{Name: "A"},
		{Name: "B"},
		{Name: "C"},
		{Name: "D"},
	}
	// 方法1 批量添加
	// db.Create(users)

	// 方法2
	//db.CreateInBatches(users,len(users))


	for _, v := range users {
		fmt.Println(v.Id)
	}

	// 方法3
	// map只添加Name值，其他值默认为null
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "A"},
		{"Name": "B"},
		{"Name": "C"},
		{"Name": "D"},
	})


}
