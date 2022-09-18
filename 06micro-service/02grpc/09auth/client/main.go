package main

import (
	"02grpc/09auth/proto/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)



func MyInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
	//  1.metadata设置token
	md := metadata.New(map[string]string{"token": "pass"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	t1 := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err!=nil{
		panic(err)
	}
	take := time.Since(t1)
	fmt.Printf("客户端执行时间:%d\n",take.Milliseconds())
	return err
}

type MyCredentials struct {

}

func (m *MyCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error){
	return map[string]string{"token": "pass"},nil
}

// 这个应该是是否使用ssl、tls吧
func(m *MyCredentials) RequireTransportSecurity() bool{
	return false
}

func main(){

	// 方法1 metadata+interceptor
	//opt := grpc.WithUnaryInterceptor(MyInterceptor)

	// 方法2 WithPerRPCCredentials
	opt := grpc.WithPerRPCCredentials(&MyCredentials{})
	conn, err := grpc.Dial("127.0.0.1:9096", grpc.WithInsecure(), opt)
	if err!=nil{
		panic(err)
	}
	defer conn.Close()



	client := pb.NewTodoServiceClient(conn)
	res, err := client.DoWork(context.Background(), &pb.TodoRequest{Msg: "I am so happy!!!"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.ResMsg)
}
