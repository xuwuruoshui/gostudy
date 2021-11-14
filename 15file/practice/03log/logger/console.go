package logger

import (
	"fmt"
)

func (l Logger) Debug(msg string) {
	console(l.Level,"DEBUG",msg)
}

func (l Logger) Info(msg string) {
	console(l.Level,"INFO",msg)
}

func (l Logger) Warning(msg string){
	console(l.Level,"WARNING",msg)
}

func (l Logger) Error(msg string) {
	console(l.Level,"ERROR",msg)
}

func (l Logger) Fatail(msg string) {
	console(l.Level,"FATAIL",msg)
}

func console(currLevel LogLevel,lv string,msg string){
	finalMsg := combinmsg(currLevel,lv,msg)
	if finalMsg!=""{
		fmt.Println(finalMsg)
	}
}