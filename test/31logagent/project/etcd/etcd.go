package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/hello/v3"
	"time"
)

var (
	cli *clientv3.Client
)

type LogEntry struct{
	// 日志存放的路径
	Path string `json:"path"`
	// 日志要发往kafka的topic
	Topic string `json:"topic"`
}

func Init(address []string,timeout int)(err error){
	cli, err = clientv3.New(clientv3.Config{
		Endpoints: address,
		// watch
		//Endpoints:   []string{"192.168.0.110:20001", "192.168.0.110:20003", "192.168.0.110:20005"},
		DialTimeout: time.Duration(timeout) * time.Second,
	})
	if err != nil {
		return err
	}
	fmt.Println("connect success")
	return nil
}

func GetConf(key string) (logEntries []*LogEntry,err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	getReq, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		panic(err)
	}
	for _, v := range getReq.Kvs {
		err = json.Unmarshal(v.Value, &logEntries)
		if err!=nil{
			return logEntries,err
		}
	}
	return logEntries,nil
}


func WatchConf(key string,newConfCh chan<-[]*LogEntry){
	channel := cli.Watch(context.Background(),key)

	// 从通道尝试取值
	for watchRes := range channel {
		for _, value := range watchRes.Events {
			fmt.Println(string(value.Kv.Key), string(value.Kv.Value))
			// 1.先判断类型
			var newConf []*LogEntry
			if value.Type!=clientv3.EventTypeDelete{
				// 如果不是删除操作
				err:=json.Unmarshal(value.Kv.Value,&newConf)
				if err!=nil{
					fmt.Println("json转换异常:",err)
					continue
				}
			}

			newConfCh<-newConf
		}
	}
}