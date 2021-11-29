package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-29 21:41:44
* @content: 消费者
 */

// MyHandler一个消费者类型
type MyHandler struct {
	Title string
}

func (m *MyHandler) HandleMessage(message *nsq.Message) (err error) {
	log.Println("收到消息:", m.Title, message.NSQDAddress, string(message.Body))
	return
}

func initConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Println("创建消费者失败,err:", err)
	}

	consumer := &MyHandler{Title: "张三"}

	c.AddHandler(consumer)

	err = c.ConnectToNSQLookupd(address)
	return err
}

func main() {
	err := initConsumer("topic_demo", "first", "192.168.0.110:4161")
	if err != nil {
		log.Panic("消费者初始化失败,err:", err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}
