package main

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2022-03-19 16:51:34
* @content: rabbitmq持久化
 */

var Conn *amqp.Connection

func init() {
	conn, err := amqp.Dial("amqp://root:root@192.168.0.110:5672/")
	if err != nil {
		fmt.Println("连接失败: ", err)
		return
	}
	Conn = conn
}

// 发布消息
func Publish(exchange, queueName, body string) (err error) {
	// 1.创建Channel
	channel, err := Conn.Channel()
	if err != nil {
		fmt.Println("创建channel失败: ", err)
		return
	}
	defer channel.Close()

	// 2.创建队列
	// durable持久化
	que, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		fmt.Println("创建queue失败: ", err)
		return
	}

	// 3.发送消息
	// DeliveryMode: amqp.Persistent 持久化
	channel.Publish(exchange, que.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	return
}

type CallBack func(msg string)

// 接受消息
func Consumer(exchange, queueName string, ctx context.Context, callback CallBack) (err error) {
	// 1.创建Channel
	channel, err := Conn.Channel()
	if err != nil {
		fmt.Println("创建channel失败: ", err)
		return
	}
	defer channel.Close()

	// 2.创建队列
	que, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		fmt.Println("创建queue失败: ", err)
		return
	}

	// autoAck: 自动/手动ack, 手动调用ack,待业务执行完后ack
	msgs, err := channel.Consume(que.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println("消费失败: ", err)
		return
	}
	fmt.Println("等待消息来临............")
	// 读取消息
	for {
		select {
		case d := <-msgs:
			callback(string(d.Body))
			d.Ack(false)
		case <-ctx.Done():
			goto End
		}
	}
End:
	return
}

var wg sync.WaitGroup

func main() {
	defer Conn.Close()

	ctx, cancelFuc := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			Publish("", "rabbitmq-test", "hello rabbitmq"+string(i))
			time.Sleep(time.Second)
		}
		wg.Done()
		cancelFuc()
	}()

	Consumer("", "rabbitmq-test", ctx, func(msg string) {
		fmt.Println("msg is :", msg)
	})
	wg.Wait()
}
