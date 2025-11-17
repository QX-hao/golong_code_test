package main

import (
	"fmt"
	"strings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main()  {
	cfg := zap.NewDevelopmentConfig()
	// 设置日志级别为 DebugLevel 以上
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	// 构建 logger
	logger, _ := cfg.Build()
	logger.Debug("这是一条 Debug 日志")  // debug级别的会被过滤掉
	logger.Info("这是一条 Info 日志")
	logger.Warn("这是一条 Warn 日志")
	logger.Error("这是一条 Error 日志")


	fmt.Println(strings.Repeat("-", 20))
	// 默认的时间要么是带了时区的，要么就是时间戳，不太美观
	// 自定义时间格式
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logger, _ = cfg.Build()
	logger.Debug("这是一条 Debug 日志")  
	logger.Info("这是一条 Info 日志")
	logger.Warn("这是一条 Warn 日志")

}