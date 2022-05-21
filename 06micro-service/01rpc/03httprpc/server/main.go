package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
* @creator: xuwuruoshui
* @date: 2022-01-24 16:22:30
* @content: http rpc server
 */

type HelloRpc struct{}

func (h *HelloRpc)Add(slice []int,result *int)error{
	for i:=0;i<len(slice);i++{
		*result +=slice[i]
	}
	return nil
}

func main(){

	rpc.Register(&HelloRpc{})

	http.HandleFunc("/jsonrpc",func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer: w,
			ReadCloser: r.Body,
		}
		//jsonrpc.ServeConn(conn)
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":8000",nil)
	
}