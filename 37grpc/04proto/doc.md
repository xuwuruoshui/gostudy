# Protobuf

```shell
# 1.安装
# protoc命令
go get -u google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
# 下载 https://github.com/protocolbuffers/protobuf/releases window版本
将proctoc.exe 放入bin目录

# 2.编写proto文件
# 自定义导出包名，写法如下:
syntax = "proto3";
package proto;
option go_package="./helloworld";

message HelloRequest{
  string name = 1;
  int32  age =2;
  repeated string hobby =3;
}

message Response{
  string replay = 1;
}

service Hello{
  rpc Hello(HelloRequest) returns (Response);
}


# 3.编译
# -I输入地址
# --goout=输出位置
# 最后一个参数输入文件名字
protoc --proto_path=. --go_out=. --go_opt=paths=source_relative helloworld.proto
# 含rpc的编译
protoc -I . helloworld.proto --go_out=plugins=grpc:.
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto
# 4.命名规范
- message: 驼峰命名，其内部字段小写加下划线
- service: 驼峰命名
- rpc方法: 驼峰命名

# 5.字段规则
required: 必填
optional: 可填可不填
repeated: 数组/切片,可以是多个值
```
