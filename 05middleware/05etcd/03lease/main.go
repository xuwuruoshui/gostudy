package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// 租约

func main()  {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.0.132:2379"},
		DialTimeout: time.Second * 5,
	})

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("connect to etcd success.")
	defer client.Close()

	ctx := context.TODO()
	// 存活时间为5秒
	resp, err := client.Grant(ctx, 5)
	if err!=nil{
		log.Fatal(err)
	}

	_, err = client.Put(ctx, "/user", "zhangsan", clientv3.WithLease(resp.ID))
	if err!=nil{
		log.Fatal(err)
	}

	for  {
		time.Sleep(time.Second)
		data, err := client.Get(ctx, "/user")
		if err!=nil{
			fmt.Println("get data err:",err)
			return
		}
		for k, v := range data.Kvs {
			fmt.Println("data:",k,v)
		}

	}

}
