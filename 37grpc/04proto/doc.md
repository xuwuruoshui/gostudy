# Protobuf

```shell
# 1.安装
go get github.com/golang/protobuf/proto
go get google.golang.org/grpc
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

# 4.命名规范
- message: 驼峰命名，其内部字段小写加下划线
- service: 驼峰命名
- rpc方法: 驼峰命名

# 5.字段规则
required: 必填
optional: 可填可不填
repeated: 数组/切片,可以是多个值
```
