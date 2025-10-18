package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	// 配置日志文件
	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)



	engine := gin.Default()

	// 自定义全局中间件处理错误
	engine.Use(func(c *gin.Context) {
		c.Next()

		// 检查是否有发生错误
		if len(c.Errors) > 0 {
			// 自定义错误处理
			c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
		}
	})

	engine.GET("/ping", func(c *gin.Context) {
		// 模拟处理过程中发生错误
		c.Error(gin.Error{Err: errors.New("处理过程中发生错误")})
	})

	engine.Run(":1234")
}
