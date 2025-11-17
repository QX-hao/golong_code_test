package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// "fmt"
)

	// fmt.Printf("\033[31mthis is 红色\n\033[0m")
	// fmt.Printf("\033[32mthis is 绿色\n\033[0m")
	// fmt.Printf("\033[33mthis is 黄色\n\033[0m")
	// fmt.Printf("\033[34mthis is 蓝色\n\033[0m")
	// fmt.Printf("\033[35mthis is 紫色\n\033[0m")
	// fmt.Printf("\033[36mthis is 青色\n\033[0m")

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorReset  = "\033[0m"
)

func EncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.DebugLevel:
		enc.AppendString(colorBlue + "DEBUG" + colorReset)
	case zapcore.InfoLevel:
		enc.AppendString(colorGreen + "INFO" + colorReset)
	case zapcore.WarnLevel:
		enc.AppendString(colorYellow + "WARN" + colorReset)
	case zapcore.ErrorLevel:
		enc.AppendString(colorRed + "ERROR" + colorReset)
	case zapcore.DPanicLevel:
		enc.AppendString(colorRed + "DPANIC" + colorReset)
	case zapcore.PanicLevel:
		enc.AppendString(colorRed + "PANIC" + colorReset)
	case zapcore.FatalLevel:
		enc.AppendString(colorRed + "FATAL" + colorReset)
	default:
		enc.AppendString(level.String()) // 默认
	}
}

func main()  {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = EncodeLevel
	logger, _ := cfg.Build()
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
	logger.DPanic("dpanic message")
	logger.Panic("panic message")
	logger.Fatal("fatal message")

}