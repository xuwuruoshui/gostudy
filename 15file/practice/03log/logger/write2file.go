package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// 写入文件
func Write2File(msg string) {
	path := "./log/"
	name := time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(path+name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("日志文件创建错误:", err)
	}
	writer := bufio.NewWriter(file)
	writer.Write([]byte(msg + "\n"))
	writer.Flush()
}