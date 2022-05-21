package main

import (
	helloworld "04proto/proto"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type JsonHelloWorld struct {
	Name  string
	Age   int32
	Hobby []string
}

func main() {
	req := helloworld.HelloRequest{Name: "zhangsan", Age: 12, Hobby: []string{"sing", "swim", "game"}}
	// protobuf原理，varint
	marshal, _ := proto.Marshal(&req)
	//fmt.Println(string(marshal))
	fmt.Println(len(marshal))

	jsonReq := JsonHelloWorld{Name: "zhangsan", Age: 12, Hobby: []string{"sing", "swim", "game"}}
	jsonmarshal, _ := json.Marshal(&jsonReq)
	fmt.Println(len(jsonmarshal))
	// 可以看出protobuf更省空间

	// 反序列化解出来也是没问题的
	newReq := &helloworld.HelloRequest{}
	proto.Unmarshal(marshal, newReq)
	fmt.Println(newReq)
}
