package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

// 测试最大并发量
func TestMaxConcurrentRequests(t *testing.T) {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go MaxConcurrentRequests()
	}
	wg.Wait()

	time.Sleep(time.Second * 1)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go MaxConcurrentRequests()
	}
	wg.Wait()
}

func MaxConcurrentRequests() {
	defer wg.Done()
	url := "http://192.168.0.107:8000/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
