package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

func main() {
	config := sarama.NewConfig()
	// ack 三种策略
	// NoResponse 直接ack
	// WaitForLocal 当前broker确认后ack
	// WaitForAll 所有broker确认后才ack
	config.Producer.RequiredAcks = sarama.WaitForAll

	// 三种分区访问策略
	// hash NewHashPartitioner key%partion
	// 随机 NewRandomPartitioner
	// 轮询 NewRoundRobinPartitioner
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	// kafka是否返回success才能操作下一条
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"192.168.0.132:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     "test",
		Partition: int32(1),
		//Key:       sarama.StringEncoder("test"),
	}

	// 键入发送信息
	ireader := bufio.NewReader(os.Stdin)
	for{
		data, _, err := ireader.ReadLine()
		if err != nil {
			break
		}
		_,err=fmt.Scanf("%s",&data)
		if err != nil {
			break
		}
		msg.Value = sarama.StringEncoder(data)
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("send message field")
		}
		fmt.Printf("partition:%d offset:%d\n",partition,offset)
	}
}
