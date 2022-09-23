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
	// 不写Where,会报ErrMissingWhereClause,防止表中数据全改
	// db.Model(&User{}).Where("name=?","fgfg").Update("age",10)

	// 另一种更新方式,结构体中有值
	// 当然结构体中写条件，再写个where也没问题
	// db.Model(&User{Id: 7}).Update("email","qwe@fff.com")

	// updates 使用结构体
	//更新
	//u3 := User{}
	//db.First(&u3,9)
	//db.Model(&u3).Updates(User{Name: "哈哈哈",Age: 17})

	// 使用map
	//u4 := User{Id:9}
	//db.Model(&u4).Updates(map[string]interface{}{
	//	"name":"fff",
	//	"age":12,
	//})

	// 只更新select中的字段
	// 用于查询出的某条数据，只更新某些字段
	//u5 := User{}
	//db.First(&u5,9)
	//db.Model(&u5).Select("name").Updates(map[string]interface{}{
	//	"Name":"嘟嘟嘟",
	//	"age": "19",
	//})

	// 忽略name值更新
	u6 := User{}
	db.First(&u6,9)
	db.Model(&u6).Omit("name").Updates(map[string]interface{}{
		"Name":"嘟嘟嘟",
		"age": "19",
	})

}
