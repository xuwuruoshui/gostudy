package deadletter

import (
	"fmt"

	"github.com/streadway/amqp"
)

/**
* @creator: xuwuruoshui
* @date: 2022-03-19 16:51:34
* @content: rabbitmq路由模式
 */

var Conn *amqp.Connection

func init() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		fmt.Println("连接失败: ", err)
		return
	}
	Conn = conn
}

func PublicEx(exchange, types, routingKey, body string) (err error) {

	// 创建channel
	channel, err := Conn.Channel()
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

func PublicDlx(exchangeA, body string) (err error) {

	// 创建channel
	channel, err := Conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	channel.Publish(exchangeA, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	return
}

type CallBack func(callBack string)

func ComsumerDlx(exchangeA, queueAName, exchangeB, queueBName string, ttl int, callBack CallBack) (err error) {

	// 创建channel
	channel, err := Conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建A交换机、A队列, 并绑定
	err = channel.ExchangeDeclare(exchangeA, "fanout", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	queueA, err := channel.QueueDeclare(queueAName, true, false, false, false, amqp.Table{
		"x-message-ttl": ttl,
		// 绑定某个交换机
		"x-dead-letter-exchange": exchangeB,
		// 绑定某个队列
		// "x-dead-letter-queue":""
		// 绑定路由
		// "x-dead-letter-routing-key":""
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = channel.QueueBind(queueA.Name, "", exchangeA, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建B交换机、B队列, 并绑定
	err = channel.ExchangeDeclare(exchangeB, "fanout", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	queueB, err := channel.QueueDeclare(queueBName, true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = channel.QueueBind(queueB.Name, "", exchangeB, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 两个交换机和队列绑定后，最终消费的还是queueB中的内容
	msg, err := channel.Consume(queueB.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wating for message.....")
	for {
		select {
		case d := <-msg:
			callBack(string(d.Body))
			d.Ack(false)
		}
	}

	return
}
