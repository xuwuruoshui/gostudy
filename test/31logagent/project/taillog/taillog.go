package taillog

import (
	"context"
	"github.com/hpcloud/tail"
	"gostudy/31logagent/project/kafka"
	"log"
	"time"
)

type TailTask struct {
	path string
	topic string
	instance *tail.Tail
	// 关gorutine
	ctx context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx,cancel := context.WithCancel(context.Background())
	tailObj =  &TailTask{
		path: path,
		topic: topic,
		ctx: ctx,
		cancelFunc: cancel,
	}
	tailObj.init()
	return
}

func (t *TailTask)init(){
	config := tail.Config{
		ReOpen:    true, // 重新打开新创建的文件
		Follow:    true, // 追加后按行返回
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false, //true:如果文件不存在则失败,false:反之亦然
		Poll:      true,  //轮询文件更改,不去通知
	}

	var err error
	t.instance,err = tail.TailFile(t.path, config)
	if err!=nil{
		panic(err)
	}

	go t.run()
}

func (t *TailTask)run(){
	for{
		select {
		// cancel后done就接到信号了
		case <-t.ctx.Done():
			log.Printf("tailtask 结束:%v_%v\n",t.path,t.topic)
			return
		// 从日志中读取按行读取
		case line:=<-t.instance.Lines:
			// 函数调用函数,需要等待消息发送到kafka后,代码才能往后继续执行,效率低
			//kafka.SendToKafka(t.topic,line.Text)
			log.Println(t.topic,":",line.Text)
			kafka.SendtoChan(t.topic,line.Text)
		default:
			time.Sleep(time.Millisecond*50)
		}
	}
}
