package main

import (
	"gopkg.in/ini.v1"
	"gostudy/31logagent/project/config"
	"gostudy/31logagent/project/kafka"
	"gostudy/31logagent/project/taillog"
	"log"
	"time"
)

var (
	cfg *config.LogConfig = new(config.LogConfig)
)

func init() {
	// 1.加载配置文件
	err := ini.MapTo(cfg, "./config/conf.ini")
	if err!=nil{
		log.Panicln(err)
	}

	// 2.连接kafka
	err = kafka.Init([]string{cfg.KafkaConfig.Address})
	if err != nil {
		log.Panicln("kafka初始化错误,err:", err)
	}
	log.Println("kafka初始化成功")

	// 3.读取文件中的日志
	err = taillog.Init(cfg.TailLogConfig.Path)
	if err != nil {
		log.Panicln("tail初始化错误,err:", err)
	}
	log.Println("sarama初始化成功")
}

func run() {
	for  {
		select {
		case line :=<-taillog.ReadLog():
			kafka.SendToKafka(cfg.KafkaConfig.Topic,line.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}

func main() {
	run()
}
