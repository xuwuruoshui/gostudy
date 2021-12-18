package utils

import (
	"math/rand"
	"time"
)

// 随机值
func RandRange(min,max int64)int64{
	return rand.Int63n(max-min)+min
}

// 获取当前时间，发送最后一条数据的时间
func Millisecond()int64{
	return time.Now().UnixNano()/int64(time.Millisecond)
}
