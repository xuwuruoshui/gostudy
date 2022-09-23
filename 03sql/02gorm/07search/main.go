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
	// 1.查一条
	//u1 := User{}
	//r := db.First(&u1)
	//fmt.Println(u1.Id)
	//fmt.Println(r.Error)
	//b := errors.Is(r.Error, gorm.ErrRecordNotFound)
	//if b{
	//	fmt.Println("查无此人")
	//}else{
	//	fmt.Println("他是:",u1.Name)
	//}

	// 2.查一条
	//u2 := User{}
	//db.Take(&u2)
	//fmt.Println(u2.Id,u2.Name)


	// 3.主键查询,主键写成字符串也能查
	//u3 := User{}
	//// 这里传入100和10, 100没有的话就继续查后面的,知道有一个查到就结束返回
	//tx := db.First(&u3, 100,10)
	//fmt.Println(u3.Name)
	//fmt.Println(tx)

	// 4.查多个
	//var users []*User
	//db.Find(&users, []int{7, 8, 9})
	//for _,v := range users {
	//	fmt.Println(v)
	//}

	// 5.条件查询
	// 5.1 等于
	//u5 := User{}
	//db.Where("name=?","fgfg").First(&u5)
	//fmt.Println(u5.Id)

	// 5.2 等于, 无法查询零值
	//u6 := User{
	//	Name: "fgfg",
	//	Age: 0,
	//}
	//db.Where(u6).First(&u6)
	//fmt.Println(u6.Id)

	// 5.2 查询多条 使用map可以查询零值
	//var users []*User
	//db.Where(map[string]interface{}{"Name":"fgfg","Age":0}).Find(&users)
	//for _, item := range users {
	//	fmt.Println(item.Id)
	//}

	// 5.4 不等于
	//var users []*User
	//db.Where("name <> ?","fgfg").Find(&users)
	//for _, v := range users {
	//	fmt.Println(v.Id)
	//}
}
