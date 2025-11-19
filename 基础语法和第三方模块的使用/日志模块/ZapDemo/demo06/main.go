package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)


// 方法一
func funOne()  {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	// 文件日志
	file, _ := os.OpenFile("app1.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	filecore := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		file,
		cfg.Level,
	)

	// 控制台日志
	consolecore := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		cfg.Level,
	)
	core := zapcore.NewTee(filecore, consolecore)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("info1日志")
}


// 方法二
func funTwo()  {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	// 文件日志
	file, _ := os.OpenFile("app2.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(file),
	)

	// 控制台日志
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		multiWriteSyncer,
		cfg.Level,
	)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("info2日志")
}

func main()  {
	funOne()
	funTwo()
}