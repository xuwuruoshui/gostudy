package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"gostudy/31logagent/06logtransfer/es"
)

var consumer sarama.Consumer



func InitConsumer(address []string)(err error){
	consumer,err = sarama.NewConsumer(address, nil)
	if err!=nil{
		return err
	}
	return nil
}

func GetMsg(topic string)(err error){
	partitionList,err := consumer.Partitions(topic)
	if err!=nil{
		return err
	}
	fmt.Println("消费者初始化成功:",partitionList)
	for partition := range partitionList {
		// 对每个分区创建一个分区消费者
		pc,err := consumer.ConsumePartition(topic,int32(partition),sarama.OffsetNewest)
		if err!=nil{
			return err
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
				// 发送消息给ES
				es.SendMsgToChan(&es.LogData{Topic: topic,Msg: string(msg.Value)})
			}
		}(pc)
	}
	select {

	}

	return nil
}
