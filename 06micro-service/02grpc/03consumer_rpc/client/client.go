package client

import (
	rpc "gostudy/36grpc/03consumer_rpc/data"
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client{
	return &Client{conn: conn}
}

// 实现通用的RPC客户端
// 传入访问的函数名
// fPtr指向的是函数原型
// var select fun xx(User)
// cli.callRPC("selectUser",&select)
func (c *Client)CallRPC(rpcName string,fPtr interface{}){
	// 通过反射，获取fPtr未初始化的函数原型
	fn := reflect.ValueOf(fPtr).Elem()
	// 需要另一个函数，作用是对第一个函数参数操作
	f := func(args []reflect.Value) []reflect.Value{
		// 处理参数
		inArgs := make([]interface{},0,len(args))

		for _,arg := range args{
			inArgs = append(inArgs,arg.Interface())
		}

		// 连接
		cliSession := rpc.NewsSession(c.conn)
		// 编码数据
		reqRPC := rpc.RPCData{Name: rpcName,Args: inArgs}
		b,err := rpc.Encode(reqRPC)
		if err!=nil{
			panic(err)
		}

		// 写数据
		err = cliSession.Write(b)
		if err!=nil{
			panic(nil)
		}

		// 服务端发过来的返回值，此时应该读取和解析
		respBytes,err := cliSession.Read()
		if err!=nil{
			panic(err)
		}

		// 解码
		respRPC,err := rpc.Decode(respBytes)
		if err!=nil{
			panic(err)
		}

		// 处理服务端返回的数据
		outArgs := make([]reflect.Value,0,len(respRPC.Args))
		for i,args := range respRPC.Args{
			// 必须进行nil转换
			if args==nil{
				// reflect.Zero() 返回一个持有类型type的零值的Value, type.Out()返回该类型的方法集中方法的数目
				outArgs = append(outArgs,reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs,reflect.ValueOf(args))
		}
		return outArgs
	}

	// 完成原型到函数调用的内部转换
	// 参数1是reflect.Type
	// 参数2 f是函数类型，是对于参数1 fn函数的操作
	// fn是定义,f是具体操作
	v := reflect.MakeFunc(fn.Type(),f)
	// 为函数fPtr赋值，过程
	fn.Set(v)
}
