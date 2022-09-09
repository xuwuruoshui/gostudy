package main

import (
	"05etcd/05register/discover"
	"log"
	"time"
)

func main() {
	var endpoints = []string{"192.168.0.132:2379"}
	ser := discover.NewServiceDiscovery(endpoints)
	defer ser.Close()
	ser.WatchService("node")
	for {
		select {
		case <-time.Tick(1 * time.Second):
			log.Println(ser.GetServices())
		}
	}
}
