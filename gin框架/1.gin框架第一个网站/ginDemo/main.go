package main

import "github.com/gin-gonic/gin"


// 首页
func index(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
		// c.String(200, "Hello, world!")
	}

func main()  {

	// 忽略gin的日志
	gin.SetMode(gin.ReleaseMode)

	// 初始化
	r := gin.Default()

	// 路由
	r.GET("/login", index)

	// 绑定端口
	r.Run("0.0.0.0:1234")
}
