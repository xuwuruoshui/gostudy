package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var wg sync.WaitGroup
func main()  {
	//config := sarama.NewConfig()
	// 消费者组
	// Rebalance
	// BalanceStrategyRange	消费分区重排
	// BalanceStrategyRoundRobin 新的分区轮询分配
	// BalanceStrategySticky  不变
	//newConfig.Consumer.Group.Rebalance =sarama.BalanceStrategyRange
	consumer, err := sarama.NewConsumer([]string{"192.168.0.132:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionlist, err := consumer.Partitions("test")
	if err != nil {
		panic(err)
	}

	for partition := range partitionlist {
		pt, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pt.Close()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for message := range pt.Messages() {
				fmt.Printf("topic:%s partition:%d offset:%d key:%s value:%s\n",message.Topic,message.Partition,message.Offset,message.Key,message.Value)
			}
		}(pt)
	}
	wg.Wait()
}


