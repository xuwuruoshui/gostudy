# protobuf
## 编写一个简单的protobuf
```protobuf
syntax = "proto3";

// 与package包对应
option go_package="./pb";

package pb;

service Study{
  rpc Study(BookRequest) returns (BookRequest);
}

message BookRequest{
  string name=1;
}

message BookResponse{
  string msg=1;
}
```

## go使用protobuf
1. 安装protoc
```shell
# 下载
https://github.com/protocolbuffers/protobuf/releases

# 将文件中的protoc.exe放到go的bin目录下
protoc.exe
```
2. 安装protoc-gen-go
```shell
# 文档
# https://grpc.io/docs/languages/go/quickstart/
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go get -u google.golang.org/grpc
```

3. 生产go文件
```shell
# 文件目录
- proto
  - pb
  book.proto

# proto目录下执行,生成go代码
protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative book.proto

# 下载依赖包
go mod tidy
```