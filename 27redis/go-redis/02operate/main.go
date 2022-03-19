package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-29 18:12:33
* @content: redis增删改查
 */

var rdb *redis.Client
var ctx context.Context
var cancel context.CancelFunc

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

	ctx,cancel = context.WithTimeout(context.Background(),5*time.Second)

	// V8新版要传context
	_,err := rdb.Ping(ctx).Result()
	return err
}

func getSet(){
	err := rdb.Set(ctx,"score",100,time.Second*2).Err()
	if err!=nil{
		log.Panic("set core faild,err:",err)
	}

	result,err := rdb.Get(ctx,"score").Result()
	if err!=nil{
		log.Panic("get core faild,err:",err)
	}

	fmt.Println(result)

	result,err = rdb.Get(ctx,"name").Result()
	if err == redis.Nil{
		log.Println("name is nil,err:",err)
	}else if err!=nil{
		log.Panic("name is nil,err:",err)
	}else{
		fmt.Println(result)
	}
}

func zgetSet(){
	key := "language_rank"
	languages := []*redis.Z{
		{Score: 90.0,Member: "Golang"},
		{Score: 98.0,Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}

	// 添加元素
	num,err := rdb.ZAdd(ctx,key,languages...).Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println(num)

	// 将元素值变大
	newScore,err := rdb.ZIncrBy(ctx,key,10.0,"Golang").Result()
	if err!=nil{
		panic(err)
	}
	fmt.Println("Golang's score is:",newScore)

	// 取分数最高的三个,0、1、2
	ret,err := rdb.ZRevRangeWithScores(ctx,key,0,2).Result()
	if err!=nil{
		panic(err)
	}
	for _, v := range ret {
		fmt.Println(v)
	}

	fmt.Println("==========================")
	// 取95-100的
	op := redis.ZRangeBy{
		Min:"98",
		Max:"100",
	}
	ret,err = rdb.ZRangeByScoreWithScores(ctx,key,&op).Result()
	if err!=nil{
		panic(err)
	}
	for _, v := range ret {
		fmt.Println(v)
	}

}

func main(){
	// getSet()
	zgetSet()

}