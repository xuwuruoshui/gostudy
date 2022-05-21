package subscript

import (
	"fmt"
	"github.com/streadway/amqp"
)

/**
* @creator: xuwuruoshui
* @date: 2022-03-19 16:51:34
* @content: rabbitmq订阅模式
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

func PublicEx(exchange, types, routingKey, body string) (err error) {

	// 创建channel
	channel, err := Conn.Channel()
	defer channel.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建交换机
	err = channel.ExchangeDeclare(exchange, types, true, false, false, false, nil)
	if err != nil {
		return
	}

	channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	return
}

type CallBack func(callBack string)

func ComsumerEx(exchange, types, routingKey string, callBack CallBack) (err error) {

	// 创建channel
	channel, err := Conn.Channel()
	defer channel.Close()
	if err != nil {
		return
	}

	// 创建交换机
	err = channel.ExchangeDeclare(exchange, types, true, false, false, false, nil)
	if err != nil {
		return
	}

	// 创建队列
	q, err := channel.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 绑定队列与交换机
	err = channel.QueueBind(q.Name, routingKey, exchange, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 消费
	msg, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		select {
		case d := <-msg:
			callBack(string(d.Body))
			d.Ack(false)
		}
	}
	return
}
