package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var(
	// 生产者
	client sarama.SyncProducer
)

// 向kafka中写日志
func Init(address []string)(err error){
	config := sarama.NewConfig()
	// 发送完数据需要leader和follower确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	//新选一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功交付的消息 将在success channel返回n
	config.Producer.Return.Successes = true


	// 连接kafka
	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		return err
	}
	return nil
}

//向kafka中发送消息,主题 内容
func SendToKafka(topic, data string){
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}