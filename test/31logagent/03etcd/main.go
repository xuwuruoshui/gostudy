package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/hello/v3"
	"time"
)

var (
	cli *clientv3.Client
)

func initEtcd() (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints: []string{"192.168.0.110:20000", "192.168.0.110:20002", "192.168.0.110:20004"},
		// watch
		//Endpoints:   []string{"192.168.0.110:20001", "192.168.0.110:20003", "192.168.0.110:20005"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	fmt.Println("connect success")
	return nil
}

func init() {
	err := initEtcd()
	if err != nil {
		fmt.Println("init etcd failed,err:", err)
	}
}

// 普通的操作，put、get
func operater() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	//putRes, err := cli.Put(ctx, "/logagent/conf", `[{"path":"f:/data/nginx.log","topic":"nginx_topic"},{"path":"f:/data/redis.log","topic":"redis_topic"},{"path":"f:/data/mysql.log","topic":"mysql_topic"}]`)
	//putRes, err := cli.Put(ctx, "/logagent/conf", `[{"path":"f:/data/redis.log","topic":"redis_topic"},{"path":"f:/data/mysql.log","topic":"mysql_topic"}]`)
	putRes, err := cli.Put(ctx, "/logagent/192.168.0.105/conf", `[{"path":"f:/data/redis.log","topic":"redis_topic"}]`)
	if err != nil {
		panic(err)
	}
	fmt.Println(putRes)

	getRes, err := cli.Get(ctx, "/logagent/conf")
	cancel()
	if err != nil {
		panic(err)
	}

	for _, v := range getRes.Kvs {
		fmt.Println(string(v.Key), string(v.Value))
	}

}

// 哨兵
func watch() {
	//fmt.Println("start watching")
	ctx,cancle := context.WithCancel(context.Background())
	channel := cli.Watch(ctx, "x")
	defer cancle()
	for k := range channel {
		fmt.Println(k)
		for _, value := range k.Events {
			fmt.Println(string(value.Kv.Key), string(value.Kv.Value))
		}
	}

}
func main() {
		//watch()
	operater()
}
