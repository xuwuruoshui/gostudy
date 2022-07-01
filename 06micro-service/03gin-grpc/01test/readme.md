# 使用
```shell
# 1.安装
# protoc命令
go get -u google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
# 下载 https://github.com/protocolbuffers/protobuf/releases window版本
将proctoc.exe 放入bin目录

cd 01test
# 编译普通entity文件和grpc文件
protoc --go_out=pb/ --go_opt=paths=source_relative --go-grpc_out=pb/ --go-grpc_opt=paths=source_relative protos/user.proto
```
