package main

import (
	"02grpc/03messagenest/proto/pb"
)

func main(){


	response := pb.TestResponse{Name: "zhangsan"}
	// UserInfo分开写
	info := pb.UserInfo{Name: "test"}
	response.UserInfo = &info

	//OtherInfo嵌套写可能就是
	//otherInfo := pb.TestResponseOtherInfo{OtherMsg: "fff"}


}
