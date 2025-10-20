package main

import (
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "index页面",
	})
}

func M1(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "中间件1请求部分",
	})

	// 阻止后续中间件和处理函数执行
	c.Abort()

	// c.Next()
	c.JSON(200, gin.H{
		"msg": "中间件1响应部分",
	})
}

func M2(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "中间件2请求部分",
	})
	c.Next()
	c.JSON(200, gin.H{
		"msg": "中间件2响应部分",
	})
}


func GM1(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "GM1中间件1请求部分",
	})

	// Abort()阻止后续中间件和处理函数执行
	// c.Abort()

	// Next()继续执行后续中间件和处理函数
	// c.Next()

	// Set()设置键值对，后续中间件和处理函数可以通过Get()获取
	c.Set("GM1", "GM1中间件1响应部分")

	c.Next()
	c.JSON(200, gin.H{
		"msg": "GM2中间件1响应部分",
	})
}

func GM2(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "GM2中间件2请求部分",
	})
	
	// Get()获取之前设置的键值对
	GM1Msg, _ := c.Get("GM1")
	c.JSON(200, gin.H{
		"msg": GM1Msg.(string),
	})

	c.Next()

	// Get()获取之前设置的键值对
	GM2Msg, _ := c.Get("GM1")
	c.JSON(200, gin.H{
		"msg": GM2Msg.(string),
	})
	
	c.JSON(200, gin.H{
		"msg": "GM2中间件2响应部分",
	})
}

func main() {
	engine := gin.Default()

	// 局部中间件
	engine.GET("/M", M1, M2, index)

	// 组内全局中间件
	g := engine.Group("/g_api")
	{
		g.Use(GM1, GM2)
		g.GET("/GM", index)
	}
	
	engine.Run(":1234")
}