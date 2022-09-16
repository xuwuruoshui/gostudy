# grpc/protobuf

## protobuf

默认值
- string ""
- double ""
- float 0
- int64 0
- int32 0
- uint32 0
- uint64 0
- bool 0
- bytes([]uint8) nil

```protobuf
syntax = "proto3";

// 与package包对应
option go_package = "./pb";

package pb;

service Study{
  rpc Study(BookRequest) returns (BookResponse);
}

message BookRequest{
  string name = 1;
}

message BookResponse{
  string msg = 1;
}
```

## 生产go

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

# grpc通常使用

server

```go
type BookInfo struct {
pb.UnimplementedStudyServer
}

func (b *BookInfo) Study(ct context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
fmt.Println(req.Name)
return &pb.BookResponse{Msg: "Server: Welcome to GRPC!!!"}, nil
}

func main() {

// 1.创建服务
server := grpc.NewServer()
pb.RegisterStudyServer(server, &BookInfo{})
listen, err := net.Listen("tcp", "0.0.0.0:9090")
if err != nil {
panic(err)
}
err = server.Serve(listen)
if err != nil {
panic(err)
}
}

```

client

```go
func main() {
conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
if err != nil {
panic(err)
}

client := pb.NewStudyClient(conn)
resp, err := client.Study(context.Background(), &pb.BookRequest{Name: "Client: I want to study go micro"})
if err != nil {
panic(err)
}
fmt.Println(resp.Msg)
}

```

# stream

三种流模式

food.protobuf

```protobuf
syntax = "proto3";

option go_package = "/pb";

service FoodService{
  rpc SayName(FoodStreamRequest) returns(stream FoodStreamResponse);// 服务端流模式
  rpc PostName(stream FoodStreamRequest) returns(FoodStreamResponse);// 客户端流模式
  rpc FullStream(stream FoodStreamRequest) returns(stream FoodStreamResponse);// 双向流模式
}

message FoodStreamRequest{
  string name = 1;
}

message FoodStreamResponse{
  string msg = 1;
}
```

server

```go

func main() {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewFoodServiceClient(conn)
	// 1. 服务端流模式
	//ServerStream(err, client)

	// 2. 客户端流模式
	//ClientStream(err, client)

	// 3. 双向模式
	fullClient, err := client.FullStream(context.Background())
	if err != nil {
		panic(err)
	}

	wg := FullStream(err, fullClient)

	wg.Wait()
}

func FullStream(err error, fullClient pb.FoodService_FullStreamClient) sync.WaitGroup {
	var wg sync.WaitGroup
	if err != nil {
		panic(err)
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			recv, err := fullClient.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(recv.Msg)
		}
	}()

	go func() {
		defer wg.Done()
		foods := []string{"东坡肘子", "回锅肉", "叫花鸡", "伤心凉粉", "担担面"}
		for _, item := range foods {
			err := fullClient.Send(&pb.FoodStreamRequest{Name: item})
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}()
	return wg
}

func ClientStream(err error, client pb.FoodServiceClient) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	clientPost, err := client.PostName(ctx)
	if err != nil {
		panic(err)
	}

	foods := []string{"东坡肘子", "回锅肉", "叫花鸡", "伤心凉粉", "担担面"}

	for _, item := range foods {
		err = clientPost.Send(&pb.FoodStreamRequest{Name: item})
		time.Sleep(time.Second)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func ServerStream(err error, client pb.FoodServiceClient) {
	res, err := client.SayName(context.Background(), &pb.FoodStreamRequest{Name: "奥里给"})
	if err != nil {
		panic(err)
	}

	for {
		recv, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv.Msg)
	}
}

```

client

```go
func main() {
conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
if err != nil {
panic(err)
}

client := pb.NewFoodServiceClient(conn)
// 1. 服务端流模式
//ServerStream(err, client)

// 2. 客户端流模式
//ClientStream(err, client)

// 3. 双向模式
fullClient, err := client.FullStream(context.Background())
if err != nil {
panic(err)
}

wg := FullStream(err, fullClient)

wg.Wait()
}

func FullStream(err error, fullClient pb.FoodService_FullStreamClient) sync.WaitGroup {
var wg sync.WaitGroup
if err != nil {
panic(err)
}
wg.Add(2)
go func() {
defer wg.Done()
for {
recv, err := fullClient.Recv()
if err != nil {
fmt.Println(err)
break
}
fmt.Println(recv.Msg)
}
}()

go func () {
defer wg.Done()
foods := []string{"东坡肘子", "回锅肉", "叫花鸡", "伤心凉粉", "担担面"}
for _, item := range foods {
err := fullClient.Send(&pb.FoodStreamRequest{Name: item})
if err != nil {
fmt.Println(err)
break
}
}
}()
return wg
}

func ClientStream(err error, client pb.FoodServiceClient) {
ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
defer cancelFunc()
clientPost, err := client.PostName(ctx)
if err != nil {
panic(err)
}

foods := []string{"东坡肘子", "回锅肉", "叫花鸡", "伤心凉粉", "担担面"}

for _, item := range foods {
err = clientPost.Send(&pb.FoodStreamRequest{Name: item})
time.Sleep(time.Second)
if err != nil {
fmt.Println(err)
break
}
}
}

func ServerStream(err error, client pb.FoodServiceClient) {
res, err := client.SayName(context.Background(), &pb.FoodStreamRequest{Name: "奥里给"})
if err != nil {
panic(err)
}

for {
recv, err := res.Recv()
if err != nil {
fmt.Println(err)
break
}
fmt.Println(recv.Msg)
}
}

```