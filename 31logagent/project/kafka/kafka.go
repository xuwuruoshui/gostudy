package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type logData struct {
	topic string
	data  string
}

var (
	// 生产者
	client      sarama.SyncProducer
	logDataChan chan *logData
)

// 向kafka中写日志
func Init(address []string, maxSize int) (err error) {
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

	// 初始化logDataChan
	logDataChan = make(chan *logData, maxSize)
	go sendToKafka()
	return nil
}


// 向kafka中发送消息,主题 内容
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)

			// 发送消息
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(50*time.Millisecond)
		}

	}

}


// 消息过多时，可以缓冲一下
func SendtoChan(topic, data string) {
	logDataChan <- &logData{
		topic: topic,
		data:  data,
	}
}