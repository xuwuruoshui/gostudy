package main

import (
	"go.uber.org/zap"
	"time"
)

func main(){
	//zap.NewDevelopment()
	pro, err := zap.NewProduction()
	if err!=nil{
		panic(err)
	}

	defer pro.Sync()
	sugar := pro.Sugar()
	sugar.Info("出错了:","time:",time.Now().Unix())
	sugar.Infof("出错了:%s","奇怪的错误")

	// 性能更高
	pro.Info("6666",zap.String("err","value"),zap.Int("err",123))

}

