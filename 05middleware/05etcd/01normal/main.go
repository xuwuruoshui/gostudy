package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

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

	ctx := context.Background()
	//putData(err, client)

	getData(ctx, err, client)
}

func putData(ctx context.Context,err error, client *clientv3.Client) {
	// 2. 添加 put
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	_, err = client.Put(ctx, "name", "ming")
	cancel()

	if err != nil {
		fmt.Println("put to etcd failed: ", err)
		return
	}
	fmt.Println("put to etcd success")
}

func getData(ctx context.Context, err error, client *clientv3.Client) {
	// 3.获取 get
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, "name")
	cancle()
	if err != nil {
		fmt.Println("get etcd failed: ", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Println("key:", string(ev.Key), " value:", string(ev.Value))
	}
}
