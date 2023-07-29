package main

import (
	"github.com/Shopify/sarama"
	"log"
	"strconv"
	"sync"
	"time"
)
var wg sync.WaitGroup
func main(){


	wg.Add(2)
	go Produce("hello",10)
	go Consume("hello")
	wg.Wait()
}

func Produce(topic string, limit int) {
	config := sarama.NewConfig()
	//config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true					// 成功交付的消息将在success channel返回
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{"192.168.0.132:9092"}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.Close()
	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))

		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(str)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
		log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", partition, offset, str)
	}
	wg.Done()
}

func Consume(topic string) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"192.168.0.132:9092"}, config)
	if err != nil {
		log.Fatal("NewConsumer err: ", err)
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
	wg.Done()
}
