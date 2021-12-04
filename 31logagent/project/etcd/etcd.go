package etcd

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	cli *clientv3.Client
)

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
