package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

var (
	client *elastic.Client
	ch     chan *LogData
)

type LogData struct {
	Topic string `json:"topic"`
	Msg   string `json:msg`
}

func InitEs(address string, maxSize int64) (err error) {
	// 关闭嗅探集群
	client, err = elastic.NewClient(elastic.SetURL(address), elastic.SetSniff(false))
	if err != nil {
		return err
	}
	fmt.Println("连接成功")

	ch = make(chan *LogData, maxSize)

	go SendMsgToEs()
	return nil
}

func SendMsgToChan(data *LogData) {
	ch <- data
}

func SendMsgToEs() {
	for {
		select {
		case msg := <-ch:
			{
				indexRes, err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
				if err != nil {
					panic(err)
				}
				fmt.Println(indexRes.Index, indexRes.Id, indexRes.Result, indexRes.SeqNo)
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

