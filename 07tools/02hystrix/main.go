package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

/**
* @creator: xuwuruoshui
* @date: 2022-02-14 22:21:14
* @content: 熔断
 */

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

}
func handler(w http.ResponseWriter, r *http.Request) {

	// 配置和Do,方法名互相对应
	// 超时、最大并发数、有5个请求才开启熔断、错误百分比为25
	hystrix.ConfigureCommand("getUser", hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  5,
		RequestVolumeThreshold: 5,
		ErrorPercentThreshold:  25,
		SleepWindow:            int(1000),
	})
	// 获取断路器状态
	c, _, _ := hystrix.GetCircuit("getUser")
	output := make(chan string, 1)

	// 异步，协程
	errors := hystrix.Go("getUser", func() error {
		// talk to other services
		output <- GetUser()
		return nil
	}, func(err error) error {

		// 获取一个func的错误
		if err != nil {
			fmt.Println(err)
			output <- `{"error":"` + err.Error() + `"}`
		}
		return nil
	})
	fmt.Println(c.IsOpen())

	select {
	case out := <-output:
		w.Write([]byte(out))
	case err := <-errors:
		w.Write([]byte(err.Error()))
	}

	// 同步api
	// err := hystrix.Do("getUser", func() error {
	// 	w.Write([]byte(GetUser()))
	// 	return nil
	// }, nil)
	// if err != nil {
	// 	w.Write([]byte(`{"error":"timeout"}`))
	// }
}

func GetUser() string {
	randNum := rand.Intn(10)
	if randNum < 6 {
		time.Sleep(time.Second * 3)
	}
	return `{"msg":"haha"}`
}
