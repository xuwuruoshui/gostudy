# grpc/protobuf

## protobuf

### 默认值

- string ""
- double ""
- float 0
- int64 0
- int32 0
- uint32 0
- uint64 0
- bool 0
- bytes([]uint8) nil

### message

尽量将message写出来来定义

```protobuf
syntax = "proto3";

option go_package = "/pb";

package pb;

message TestRequest{
  int64 id = 1;
}

message TestResponse{
  string name = 1;
  UserInfo userInfo = 2;
}

message UserInfo{
  string name = 1;
}
```

### 常用类型

```protobuf
syntax = "proto3";
// 导入时间类型
import "google/protobuf/timestamp.proto";

option go_package = "/pb";

// 枚举
enum Week{
  Sunday = 0;
  Monday = 1;
  Tuesday = 2;
  Wednesday = 3;
  Thursday = 4;
  Friday = 5;
  Saturday = 6;
}

message TodoRequest{
  string todo = 1;
  Week week = 2;
  // map类型
  map<string, string> bookMap = 3;
  google.protobuf.Timestamp doneTime = 4;
}

message TodoResponse{
  bool done = 1;
}


```

## protobuf生成go

1. 编写文件

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

2. 安装protoc

```shell
# 下载
https://github.com/protocolbuffers/protobuf/releases

# 将文件中的protoc.exe放到go的bin目录下,include放到go的bin目录同级
- bin
  protoc.exe
- include
```

3. 安装protoc-gen-go

```shell
# 文档
# https://grpc.io/docs/languages/go/quickstart/
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go get -u google.golang.org/grpc
```

4. 生成go文件

```shell
# 文件目录
- proto
  - pb
  book.proto

# proto目录下执行,生成go代码
# 第一种
protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative book.proto
# 第二种
protoc -I=.  --go_out=. --go-grpc_out=. machine.proto  
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

# metadata
相当于http中的header

todo.proto
```protobuf
syntax="proto3";

option go_package="/pb";

service TodoService{
  rpc DoWork(TodoRequest) returns(TodoResponse);
}

message TodoRequest{
  string msg = 1;
}

message TodoResponse{
  string resMsg = 1;
}
```

server
```go
type Todo struct {
	pb.UnimplementedTodoServiceServer
}

func(t *Todo) DoWork(ctx context.Context,req *pb.TodoRequest) (*pb.TodoResponse, error){

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok{
		fmt.Println("metadata不存在")
		return nil, nil
	}
	for k,v := range md {
		fmt.Printf("%s:%s\n",k,v)
	}

	fmt.Println("服务端已接受到客户端消息: ",req.Msg)
	return &pb.TodoResponse{ResMsg: "你要学习: go微服务,grpc"},nil
}

func main(){




	server := grpc.NewServer()
	pb.RegisterTodoServiceServer(server,&Todo{})

	listen, err := net.Listen("tcp", ":9094")
	if err!=nil{
		panic(err)
	}
	err = server.Serve(listen)
	if err!=nil{
		panic(err)
	}
}

```

client
```go
func main(){
	conn, err := grpc.Dial("127.0.0.1:9094",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewTodoServiceClient(conn)

	// 必须是ascii码
	//md1 := metadata.New(map[string]string{
	//	"name": "metadata try",
	//})
	md1 := metadata.Pairs("name","microserver","key","value")

	ctx := metadata.NewOutgoingContext(context.Background(), md1)
	resp, err := client.DoWork(ctx, &pb.TodoRequest{Msg: "看一下要学什么"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.ResMsg)
}

```

# interceptor
server和client拦截器

server
```go
type Todo struct {
	pb.UnimplementedTodoServiceServer
}

func (t Todo) DoWork(ctx context.Context,req *pb.TodoRequest) (*pb.TodoResponse, error) {

	msg := req.GetMsg()
	fmt.Println(msg)
	return &pb.TodoResponse{ResMsg: "hahaha"},nil
}

func MyInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
	t1 := time.Now()
	resp, err = handler(ctx, req)
	take := time.Now().Sub(t1)
	fmt.Printf("执行时间:%d",take.Milliseconds())
	return
}

func main(){

	interceptor := grpc.UnaryInterceptor(MyInterceptor)
	server := grpc.NewServer(interceptor)

	pb.RegisterTodoServiceServer(server,&Todo{})

	listen, err := net.Listen("tcp", ":9095")
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
func MyInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
	t1 := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err!=nil{
		panic(err)
	}
	take := time.Since(t1)
	fmt.Printf("客户端执行时间:%d\n",take.Milliseconds())
	return err
}

func main(){
	opt := grpc.WithUnaryInterceptor(MyInterceptor)
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure(), opt)
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
```
