package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

/**
* @creator: xuwuruoshui
* @date: 2021-12-01 19:53:30
* @content: github.com/Shopify/sarama,基于sarama第三方库开发的kafka hello
* 主要用于写日志,根据教程里写的好像需要gcc,可能以前学c安装过gcc,所以这儿自己就好了
 */

var producer sarama.SyncProducer
func initProducer(){
	config := sarama.NewConfig()

	// 发送完数据需要leader和follower确认
	config.Producer.RequiredAcks = sarama.WaitForAll

	//新选一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	// 成功交付的消息 将在success channel返回n
	config.Producer.Return.Successes = true


	// 连接kafka
	var err error
	producer, err = sarama.NewSyncProducer([]string{"192.168.0.110:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
}
func sendMsg(){
	// 构造一一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("abcd")
	// 发送消息
	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

var consumer sarama.Consumer
func initConsumer() {
	var err error
	consumer,err = sarama.NewConsumer([]string{"192.168.0.110:9092"}, nil)
	if err!=nil{
		fmt.Println("消费初始化错误,err:",err)
		return
	}
	fmt.Println("消费者连接成功")
}
func getMsg(){
		partitionList,err := consumer.Partitions("redis_topic")
		if err!=nil{
			fmt.Println("分区列表获取失败,err:",err)
			return
		}
		fmt.Println("消费者初始化成功:",partitionList)
		for partition := range partitionList {
			// 对每个分区创建一个分区消费者
			pc,err := consumer.ConsumePartition("redis_topic",int32(partition),sarama.OffsetNewest)
			if err!=nil{
				fmt.Println("消费者获取失败,err:",err)
				return
			}
			defer pc.AsyncClose()
			// 异步从每个分区消费信息
			go func(partitionConsumer sarama.PartitionConsumer) {
				for msg := range pc.Messages() {
					fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
				}
			}(pc)
		}
	select {

		}
}
func main() {
	// 发送消息
	initProducer()
	sendMsg()

	// 接受消息
	//initConsumer()
	//getMsg()
}
