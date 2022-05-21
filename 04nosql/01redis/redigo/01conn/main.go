package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2022-03-19 16:51:34
* @content: redisgo包连接
 */

var Pool *redis.Pool

func init() {

	Pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		Wait:        true, //超过最大连接数时，是等待还是报错
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "192.168.0.110:6379")
		},
	}
}

func main() {
	c := Pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池

	_, err := c.Do("set", "username", 10)
	if err != nil {
		fmt.Println("redis操作失败: ", err)
	}
	// 设置过期时间
	c.Do("expire", "username", 10000)
	r, err := redis.Int64(c.Do("get", "username"))
	if err != nil {
		fmt.Println("redis操作失败: ", err)
	}
	fmt.Println(r)
	// 获取过期时间
	time.Sleep(time.Second * 2)
	ttl, err := c.Do("ttl", "username")
	fmt.Printf("%d", ttl)
}
