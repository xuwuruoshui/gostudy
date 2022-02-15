package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

/**
* @creator: xuwuruoshui
* @date: 2022-02-14 22:21:14
* @content: zap、lumberjack使用
 */

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func init() {
	// 直接使用
	//logger, _ = zap.NewProduction()
	//sugarLogger = logger.Sugar()

	// 客制化
	encode := func() zapcore.Encoder {
		encoderConfig := zap.NewProductionEncoderConfig()
		//转换时间格式、日志级别大写
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		// 普通logger
		return zapcore.NewConsoleEncoder(encoderConfig)
		// json形式
		//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}
	writerSync := func() zapcore.WriteSyncer {
		// 控制台打印
		//return zapcore.AddSync(os.Stderr)

		//file, _ := os.Create("./04zap/test.log")
		//引入lumberjack，文件切割
		//MaxSize：日志文件的最大大小（以MB为单位）
		//MaxBackups：保留旧文件的最大个数
		//MaxAges：保留旧文件的最大天数
		//Compress：是否压缩/归档旧文件
		lumberjacker := &lumberjack.Logger{
			Filename:   "./04zap/test.log",
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   false,
		}
		return zapcore.AddSync(lumberjacker)
	}
	core := zapcore.NewCore(encode(), writerSync(), zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func main() {
	defer logger.Sync()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")

}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		//logger.Error("Error fetching url...", zap.String("url", url), zap.Error(err))
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		//logger.Info("Success...", zap.String("status", resp.Status), zap.String("url", url))
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
