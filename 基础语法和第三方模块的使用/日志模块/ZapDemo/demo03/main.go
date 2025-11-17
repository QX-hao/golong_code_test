package main

import (
	"go.uber.org/zap"
	// "fmt"
)

func main()  {
	logger, _ := zap.NewDevelopment()

	// *zap.Logger的功能并不多
	// 使用 *zap.SugarLogger 进行格式化输出

	// 类似sprintf的格式化输出
	// logger.Info(fmt.Sprintf("这是一条 %s 日志", "Info"))
	sugar := logger.Sugar()
	sugar.Infof("这是一条 %s 日志", "Info")
	sugar.Debugf("这是一条 %s 日志", "Debug")
	sugar.Warnf("这是一条 %s 日志", "Warn")
	sugar.Errorf("这是一条 %s 日志", "Error")

}