package main

import (
	"gostudy/15file/practice/03log/logger"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-13 11:56:15
* @content:
 */

var log logger.Logger

func main() {
	// log := logger.NewConsoleLogger("debug")

	// log.Debug("这是一个Debug日志")
	// log.Info("这是一个Info日志,id:%d name:%s",125,"张三")
	// log.Warning("这是一个Warning日志")
	// log.Error("这是一个Error日志")

	log = logger.NewFileLogger("info","","./log/",1024)
	// fileLog.Name="haha"
	for {
		log.Debug("这是一个Debug日志")
		log.Info("这是一个Info日志,id:%d name:%s", 125, "张三")
		log.Warning("这是一个Warning日志")
		log.Error("这是一个Error日志")
		time.Sleep(time.Second*2)
	}

}
