package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gostudy/31logagent/06logtransfer/conf"
	"gostudy/31logagent/06logtransfer/es"
	"gostudy/31logagent/06logtransfer/kafka"
	"log"
)


func main(){
 	// 1. 获取配置文件
	var logTransfer conf.LogTransfer
	err := ini.MapTo(&logTransfer, "./conf/conf.ini")
	if err!=nil{
		panic(err)
	}
	fmt.Println(logTransfer)

	// 2. 初始化es
	err = es.InitEs(logTransfer.EsConf.Address,logTransfer.EsConf.MaxSize)
	if err!=nil{
		panic(err)
	}

	// 3. 初始化kafka
	err = kafka.InitConsumer(logTransfer.KafkaConf.Address)
	if err!=nil{
		log.Panicln("kafka消费者初始化错误,err:",err)
	}
	err = kafka.GetMsg(logTransfer.KafkaConf.Topic)
	if err!=nil{
		log.Panicln("kafka消费者分区异常,err:",err)
	}



}
