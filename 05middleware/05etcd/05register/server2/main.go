package main

import (
	"05etcd/05register/register"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)



func main(){
	// 注册
	var endpoints = []string{"192.168.0.132:2379"}
	ser, err := register.NewServiceRegister(endpoints, "node"+time.Now().String(), "127.0.0.1:8082", 5)
	if err!=nil{
		log.Fatalln(err)
	}

	go ser.ListenLeaseRespChan()

	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK,"ok")
	})
	engine.Run(":8082")
}
