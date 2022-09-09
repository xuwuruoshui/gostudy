package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)
// watch获取数据的变化

func main() {
	// 1. 连接
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.0.132:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect faild: ", err)
		return
	}
	fmt.Println("connect to etcd success: ")
	defer client.Close()

	watcher := client.Watch(context.Background(), "name")
	for response:= range watcher {
		for _, ev := range response.Events {
			fmt.Println("Type:",ev.Type," Key:",string(ev.Kv.Key)," Value:",string(ev.Kv.Value))
		}
	}
}
