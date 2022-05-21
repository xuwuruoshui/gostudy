package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-29 16:38:13
* @content: Redis连接
 */

var rdb *redis.Client

func init(){
	err := initRDB()	
	if err!=nil{
		panic(err)
	}
	fmt.Println("连接成功")
}

func initRDB() error{
	rdb = redis.NewClient(&redis.Options{
		Addr: "192.168.0.110:6379",
		Password: "",
		DB: 0,
	})

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	// V8新版要传context
	_,err := rdb.Ping(ctx).Result()
	return err
}



func main(){
}