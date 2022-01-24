package rpc

import (
	"encoding/binary"
	"io"
	"net"
)

// 会话
type Session struct {
	conn net.Conn
}

func NewsSession(conn net.Conn)*Session{
	return &Session{conn: conn}
}

// 写入数据，需要考虑站包
func (s *Session)Write(data []byte) (err error){
	// 定义写数据的格式
	// 4字节头 + 任意长度文件(最大也就2G 1<<31-1)
	buf := make([]byte,4+len(data))

	// 4字节头
	binary.BigEndian.PutUint32(buf[:4],uint32(len(data)))

	// 任意长度文件(最大也就2G 1<<31-1)
	copy(buf[4:],data)
	_,err = s.conn.Write(buf)
	if err!=nil{
		return
	}
	return
}

// 读数据
func (s *Session)Read()(data []byte,err error){
	// 读取头部文件
	header := make([]byte,4)

	// 按长度读取消息
	_,err = io.ReadFull(s.conn,header)
	if err!=nil{
		return
	}

	// 读取数据
	dataLen := binary.BigEndian.Uint32(header)
	data = make([]byte,dataLen)
	_,err = io.ReadFull(s.conn,data)
	if err!=nil{
		return
	}
	return
}