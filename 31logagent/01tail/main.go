package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

/**
* @creator: xuwuruoshui
* @date: 2021-12-01 16:56:28
* @content: github.com/hpcloud/tail使用,主要用于读日志
* tip:包报错可能是由于在gopath下
 */

// tail demo
func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true, // 重新打开新创建的文件
		Follow:    true, // 追加后按行返回
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false, //true:如果文件不存在则失败,false:反之亦然
		Poll:      true,  //轮询文件更改,不去通知
	}

	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}

	for {
		msg, ok := <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}

}
