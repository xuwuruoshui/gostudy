package logger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAIL
)

type Logger interface {
	Debug(fmtMsg string, param ...interface{})
	Info(fmtMsg string, param ...interface{})
	Warning(fmtMsg string, param ...interface{})
	Error(fmtMsg string, param ...interface{})
	Fatail(fmtMsg string, param ...interface{})
}

// string转日志级别
func parseLogLevel(level string) (LogLevel, error) {
	lv := strings.ToUpper(level)

	switch lv {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAIL":
		return FATAIL, nil
	default:
		err := errors.New("无法解析日志级别")
		return UNKNOWN, err
	}
}

// 组合信息
func combinmsg(currLv LogLevel, lvMsg string, fmtMsg string, param ...interface{}) (finalMsg string) {

	lv, err := parseLogLevel(lvMsg)
	if err != nil {
		panic(err)
	}
	if currLv <= lv {
		timeMsg := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
		level := " [" + lvMsg + "] "
		pc, file, line, ok := runtime.Caller(3)
		var position string
		if ok {
			fileName := runtime.FuncForPC(pc).Name()
			basePath := path.Base(file)
			position = " [" + fileName + ":" + basePath + ":" + strconv.Itoa(line) + "] "
		}

		fmtMsg := fmt.Sprintf(fmtMsg, param...)
		finalMsg = timeMsg + level + fmtMsg + position
	}
	return finalMsg
}
