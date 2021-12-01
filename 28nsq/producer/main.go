package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/nsqio/go-nsq"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-29 21:22:12
* @content: nsq一个消息队列。生产者
 */

var producer *nsq.Producer

func init(){
	err := initProduct("192.168.0.110:4150")
	if err!=nil{
		log.Panic(err)
	}
}

func initProduct(str string) (err error){
	config := nsq.NewConfig()
	producer,err = nsq.NewProducer(str,config)
	return err
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	for{
		data,err := reader.ReadString('\n')
		if err!=nil{
			log.Println(err)
		}
		data = strings.TrimSpace(data)
		if strings.ToUpper(data)=="Q"{
			break
		}

		// 向topic_demo中推送数据
		err = producer.Publish("topic_demo",[]byte(data))
		if err!=nil{
			log.Println(err)
		}
	}
	

}