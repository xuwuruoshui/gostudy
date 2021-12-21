# Protobuf

```shell
# 1.安装
go get github.com/golang/protobuf/protoc-gen-go
# 下载 https://github.com/protocolbuffers/protobuf/releases window版本
将proctoc.exe 放入bin目录

# 2.编写proto文件
# 自定义导出包名，写法如下:
syntax = "proto3";
package proto;
option go_package=./你的包名

message Sender{
  string name =1;
}

message Reciver{
  string name=1;
  int32 age =2;
  bool isMarried=3;
  repeated string hobby=4;
}

service TestService{
  rpc Search(Sender) returns(Reciver){};
}

# 3.编译
# -I输入地址
# --goout=输出位置
# 最后一个参数输入文件名字
protoc -I ./ --go_out=../ test.proto
# 含rpc的编译
protoc -I=. --go_out=plugins=grpc:../ test.proto
```
