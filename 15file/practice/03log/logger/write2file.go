package logger

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"sync"
	"time"
)

var one sync.Once

// 文件
type FileLogger struct {
	Level  LogLevel
	Name   string
	Suffix string
	Path   string
	logMsg chan *FileMsg
	Size   int64
}

type FileMsg struct {
	Level  LogLevel
	Name   string
	Suffix string
	Path   string
	Msg    string
}

// 没有名字设置默认为""
func NewFileLogger(level string, name string, path string, size int64) *FileLogger {
	lv, err := parseLogLevel(level)
	if err != nil {
		panic(err)
	}
	channel := make(chan *FileMsg, 100)
	return &FileLogger{Level: lv, Name: name, Path: path, Size: size, logMsg: channel}
}

// 判断文件格式等信息
func (f FileLogger) write(currLevel LogLevel, lv string, fmtMsg string, param ...interface{}) {
	finalMsg := combinmsg(currLevel, lv, fmtMsg, param...)

	// 默认当前系统时间
	if f.Name == "" {
		f.Name = time.Now().Format("2006-01-02")
	}

	// 默认日志大小
	if f.Size == 0 {
		f.Size = 1024 * 5
	}
	f.Suffix = ".log"

	if finalMsg != "" {
		//f.Write2File(finalMsg)
		f.logMsg <- &FileMsg{Level: f.Level, Name: f.Name, Suffix: f.Suffix, Path: f.Path, Msg: finalMsg}

		// 日志级别为error以上,单独创建一个文件
		loglv, err := parseLogLevel(lv)
		if err != nil {
			panic(err)
		}
		if loglv >= ERROR {
			f.Suffix = ".err"
			f.logMsg <- &FileMsg{Level: f.Level, Name: f.Name, Suffix: f.Suffix, Path: f.Path, Msg: finalMsg}
			//f.Write2File(finalMsg)
		}
	}

	one.Do(func() {
		go f.Write2File()
	})
}

// 写入文件
func (f FileLogger) Write2File() {
	for {

		data := *<-f.logMsg
		path := data.Path
		name := data.Name + data.Suffix
		file, err := os.OpenFile(path+name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
		if err != nil {
			fmt.Println("日志文件创建错误:", err)
		}

		fileInfo, err := file.Stat()
		if err != nil {
			panic(err)
		}

		msg := data.Msg
		// 文件切割
		Newfile, ok := f.fileSplit(fileInfo, msg, file, path, name)
		if ok {
			file = Newfile
		}
		writer := bufio.NewWriter(file)
		writer.Write([]byte(msg + "\n"))
		writer.Flush()
		file.Close()

	}
}

// 文件切割
func (f FileLogger) fileSplit(fileInfo fs.FileInfo, msg string, file *os.File, path string, name string) (*os.File, bool) {
	// 文件大于等于1M进行分割
	fmt.Println(fileInfo.Size())

	if fileInfo.Size()+int64(len(msg)) > f.Size {
		// 1.关闭文件
		file.Close()

		// 2.重命名
		os.Rename(path+name, path+name+".bak"+time.Now().Format("20060102150405"))

		// 3.创建新文件
		file, err := os.OpenFile(path+name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
		if err != nil {
			fmt.Println("日志文件创建错误:", err)
		}
		return file, true
	}
	return nil, false
}

func (f FileLogger) Debug(fmtMsg string, param ...interface{}) {
	f.write(f.Level, "DEBUG", fmtMsg, param...)
}

func (f FileLogger) Info(fmtMsg string, param ...interface{}) {
	f.write(f.Level, "INFO", fmtMsg, param...)
}

func (f FileLogger) Warning(fmtMsg string, param ...interface{}) {
	f.write(f.Level, "WARNING", fmtMsg, param...)
}

func (f FileLogger) Error(fmtMsg string, param ...interface{}) {
	f.write(f.Level, "ERROR", fmtMsg, param...)
}

func (f FileLogger) Fatail(fmtMsg string, param ...interface{}) {
	f.write(f.Level, "FATAIL", fmtMsg, param...)
}
