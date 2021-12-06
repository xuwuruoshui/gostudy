package taillog

import (
	"fmt"
	"gostudy/31logagent/project/etcd"
	"time"
)

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntries []*etcd.LogEntry
	tskMap map[string]*TailTask
	newConfChan chan[]*etcd.LogEntry
}

func Init(logEntries []*etcd.LogEntry){
	tskMgr = &tailLogMgr{
		logEntries: logEntries,
		tskMap: make(map[string]*TailTask,16),
		newConfChan:  make(chan[]*etcd.LogEntry),
	}
	for _, logEntry := range tskMgr.logEntries {
		tailTask := NewTailTask(logEntry.Path,logEntry.Topic)
		tskMgr.tskMap[logEntry.Topic+"_"+logEntry.Path] = tailTask
	}

	go tskMgr.run()
}

func (t *tailLogMgr)run(){
	for{
		select{
		case newConf:=<-t.newConfChan:

			// 1.判断是否存在topic,不存在就新增
			for _, conf := range newConf {
				_,ok := t.tskMap[conf.Topic+"_"+conf.Path]
				if ok{
					continue
				}else{
					tailTask := NewTailTask(conf.Path,conf.Topic)
					tskMgr.tskMap[conf.Topic+"_"+conf.Path] = tailTask
				}
			}

			// 2. 配置删除
			for _, c1 := range t.logEntries {
				isDel := true
				for _, c2 := range newConf {
					if c1.Topic==c2.Topic && c1.Path==c2.Path {
						isDel = false
						continue
					}
				}
				if isDel{
					// 关闭
					t.tskMap[c1.Topic+"_"+c1.Path].cancelFunc()
				}
			}
			// 3. 配置变更
			fmt.Println("新配置:",newConf)
		default:
			time.Sleep(50*time.Millisecond)
		}
	}
}

// 一个函数，向外暴露tskMgr的newConfChan
func GetNewConfChan()chan<-[]*etcd.LogEntry{
	return tskMgr.newConfChan
}