package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

/**
* @creator: xuwuruoshui
* @date: 2021-12-01 19:53:30
* @content: github.com/Shopify/sarama,主要用于写日志,根据教程里写的好像需要gcc,可能以前学c安装过gcc,所以这儿自己就好了
 */

// 基于sarama第三方方库开发的kafka client
func main() {
	config := sarama.NewConfig()

	// 发送完数据需要leader和follower确认
	config.Producer.RequiredAcks = sarama.WaitForAll

	//新选一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner 

	// 成功交付的消息 将在success channel返回n
	config.Producer.Return.Successes = true

	// 构造一一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("qaqaqa")

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.0.110:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
