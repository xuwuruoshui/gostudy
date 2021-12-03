package taillog

import (
	"github.com/hpcloud/tail"
)
// 读取日志模块到终端

var (
	tailObj *tail.Tail
	tailLog *tail.Line
)

// 初始化日志模块
func Init(fileName string)(err error){
	config := tail.Config{
		ReOpen:    true, // 重新打开新创建的文件
		Follow:    true, // 追加后按行返回
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false, //true:如果文件不存在则失败,false:反之亦然
		Poll:      true,  //轮询文件更改,不去通知
	}

	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		return err
	}

	return nil
}

// 读取日志模块
func ReadLog() <-chan *tail.Line {
	return tailObj.Lines
}
