package rpc

import (
	"bytes"
	"encoding/gob"
)

type RPCData struct {
	// 函数名
	Name string
	// 对应的参数
	Args []interface{}
}

func Encode(data RPCData)([]byte,error){

	// 创建一个字节数组编码器
	var buf bytes.Buffer
	bufEnc := gob.NewEncoder(&buf)

	// 对数据编码
	if err:=bufEnc.Encode(data);err!=nil{
		return nil,err
	}

	return buf.Bytes(),nil
}

func Decode(b []byte)(RPCData,error){
	buf := bytes.NewBuffer(b)
	// 创建一个字节数组编码器
	bufDec := gob.NewDecoder(buf)
	var data RPCData

	//解码
	if err:=bufDec.Decode(&data);err!=nil{
		return data,err
	}
	return data,nil
}
