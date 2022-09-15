package main

import (
	"02grpc/01protobuf/proto/pb"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	Name string
}

func main() {
	// GRPC
	request := pb.BookRequest{Name: "Go Test"}
	data, err := proto.Marshal(&request)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))

	// Json
	data2, err := json.Marshal(&Book{
		Name: "Go Test",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(data2)
	fmt.Println(string(data2))

	//Go Test
	//{"Name":"Go Test"}
	// grpc体积比json小

}
