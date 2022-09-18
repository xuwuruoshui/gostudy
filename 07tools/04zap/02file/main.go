package main

import "go.uber.org/zap"

func NewLogger()(*zap.Logger,error){
	config := zap.NewProductionConfig()
	config.OutputPaths = append(config.OutputPaths,"./test.log")
	return config.Build()
}

func main(){
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Info("test")
	sugar.Errorf("err....")
}
