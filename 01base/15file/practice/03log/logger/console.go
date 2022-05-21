package logger

import "fmt"

// 终端
type ConsoleLogger struct {
	Level LogLevel
}

func NewConsoleLogger(level string) *ConsoleLogger {
	lv, err := parseLogLevel(level)
	if err != nil {
		panic(err)
	}
	return &ConsoleLogger{Level: lv}
}

func console(currLevel LogLevel, lv string, fmtMsg string, param ...interface{}) {
	finalMsg := combinmsg(currLevel, lv, fmtMsg, param...)
	if finalMsg != "" {
		fmt.Println(finalMsg)
	}
}

func (l ConsoleLogger) Debug(fmtMsg string, param ...interface{}) {
	console(l.Level, "DEBUG", fmtMsg, param...)
}

func (l ConsoleLogger) Info(fmtMsg string, param ...interface{}) {
	console(l.Level, "INFO", fmtMsg, param...)
}

func (l ConsoleLogger) Warning(fmtMsg string, param ...interface{}) {
	console(l.Level, "WARNING", fmtMsg, param...)
}

func (l ConsoleLogger) Error(fmtMsg string, param ...interface{}) {
	console(l.Level, "ERROR", fmtMsg, param...)
}

func (l ConsoleLogger) Fatail(fmtMsg string, param ...interface{}) {
	console(l.Level, "FATAIL", fmtMsg, param...)
}
