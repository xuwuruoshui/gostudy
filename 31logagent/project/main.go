package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gostudy/31logagent/project/config"
	"gostudy/31logagent/project/etcd"
	"gostudy/31logagent/project/kafka"
	"gostudy/31logagent/project/taillog"
	"gostudy/31logagent/project/utils"
	"log"
	"sync"
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
	err = kafka.Init([]string{cfg.KafkaConfig.Address},cfg.KafkaConfig.MaxSize)
	if err != nil {
		log.Panicln("kafka初始化错误,err:", err)
	}
	log.Println("kafka初始化成功")

	// 3.初始化etcd
	err = etcd.Init(cfg.EtcdConfig.Address,cfg.EtcdConfig.Timeout)
	if err != nil {
		log.Panicln("etcd初始化错误,err:", err)
	}
	log.Println("etcd初始化成功")

	// 4 从etcd中获取日志收集信息
	key := fmt.Sprintf(cfg.EtcdConfig.Key,utils.GetLocalIP())
	logEntries,err := etcd.GetConf(key)
	if err != nil {
		log.Panicln("从etcd中获取日志收集信息失败,err:", err)
	}
	log.Println("从etcd中获取日志收集信息成功,logEntries:")

	// 5.将配置发送给taillog_mgr,往kafka中发送日志消息
	taillog.Init(logEntries)
	log.Println("sarama初始化成功")

	// 6.开启一个watch去实时监视日志收集项的变化(有变化,则将Chan放入通道中)
	newConfChan := taillog.GetNewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(key,newConfChan)
	wg.Wait()

}


func main() {

}
