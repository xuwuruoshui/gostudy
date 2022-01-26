package test

import (
	"fmt"
	rpc "gostudy/36grpc/03consumer_rpc/data"
	"net"
	"sync"
	"testing"
)

func TestSession_ReadWriter(t *testing.T) {
	addr := "127.0.0.1:8000"
	my_data := "test"

	var wg sync.WaitGroup
	wg.Add(2)

	// 写数据的协程
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		s := rpc.NewsSession(conn)
		err = s.Write([]byte(my_data))
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 读数据协程
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)

		if err != nil {
			t.Fatal(err)
		}
		s := rpc.NewsSession(conn)

		data, err := s.Read()
		if err != nil {
			t.Fatal(err)
		}
		//if my_data!=string(data){
		//	t.Fatal("数据不一致")
		//}
		fmt.Println(string(data))

	}()

	wg.Wait()
}
