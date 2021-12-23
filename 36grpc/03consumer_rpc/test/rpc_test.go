package test

import (
	"encoding/gob"
	"fmt"
	"gostudy/36grpc/03consumer_rpc/client"
	rpc "gostudy/36grpc/03consumer_rpc/server"
	"net"
	"testing"
)

type User struct {
	Name string
	Age int
}

func queryUser(uid int)(User,error){

	user := make(map[int]User)
	// 假数据
	user[0] = User{"zs",20}
	user[1] = User{"ls",21}
	user[2] = User{"ww",22}

	// 模拟查询用户
	if u,ok := user[uid];ok{
		return u,nil
	}

	return User{},fmt.Errorf("%d err",uid)
}


func TestRPC(t *testing.T){

	// 编码中有一个字段是interface{}时,要注册一下
	gob.Register(User{})
	addr := "127.0.0.1:8000"
	// 创建服务端
	srv := rpc.NewServer(addr)
	// 将服务端方法，注册一下
	srv.Register("queryUser",queryUser)
	// 服务端等待调用
	go srv.Run()

	// 客户端端获取连接
	conn,err := net.Dial("tcp",addr)
	if err!=nil{
		fmt.Println("err")
	}
	// 创建客户端对象
	cli := client.NewClient(conn)
	// 需要声明函数原型
	var query func(int)(User,error)
	cli.CallRPC("queryUser",&query)
	// 得到查询结果
	u,err := query(0)
	if err!=nil{
		fmt.Println("err")
	}
	fmt.Println(u)
}
