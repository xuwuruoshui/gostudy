package logger

import (
	"errors"
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

type Logger struct {
	Level LogLevel
}

func NewLogger(level string) Logger {
	lv, err := parseLogLevel(level)
	if err != nil {
		panic(err)
	}
	return Logger{Level: lv}
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
func combinmsg(currLv LogLevel, lvMsg string, msg string) (finalMsg string) {

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
			fileName :=  runtime.FuncForPC(pc).Name()
			basePath := path.Base(file)
			position = " [" + fileName + ":" + basePath + ":" + strconv.Itoa(line) + "] "
		}

		finalMsg = timeMsg + level + msg + position
	}
	return finalMsg
}
