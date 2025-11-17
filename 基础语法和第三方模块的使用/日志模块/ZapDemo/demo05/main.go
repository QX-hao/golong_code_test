package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// "fmt"
)

func main() {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)
	logger := zap.New(core, zap.AddCaller())
}
