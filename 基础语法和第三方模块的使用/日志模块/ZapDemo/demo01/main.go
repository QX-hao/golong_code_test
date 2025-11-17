package main


import (
	"go.uber.org/zap"
)

// dev模式下，日志格式是text格式，并且warn和error会有栈信息
func dev()  {
	logger, _ := zap.NewDevelopment()
	logger.Info("dev this is info",
		zap.String("name", "dev"),
		zap.Int("age", 18),
		zap.Bool("isMale", true),
	)
	logger.Info("dev this is info",)
	logger.Warn("dev this is warn")
	logger.Error("dev this is error")

}

// example模式下，格式是json，并且字段只有level和msg
func test() {
	logger := zap.NewExample()
	logger.Info("exam this is info")
	logger.Warn("exam this is warn")
	logger.Error("exam this is error")

}

// prod模式下，格式也是json，多一个时间和函数位置字段，生产环境上用json格式的日志更方便排查
func prod() {
	logger, _ := zap.NewProduction()
	logger.Info("prod this is info",
		zap.String("name", "prod"),
		zap.Int("age", 18),
		zap.Bool("isMale", true),
	)

	logger.Warn("prod this is warn")
	logger.Error("prod this is error")
}

func main()  {
	dev()
	test()
	prod()
}