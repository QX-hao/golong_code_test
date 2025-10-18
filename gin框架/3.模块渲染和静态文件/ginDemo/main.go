package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	// 模块渲染  
	// LoadHTMLGlob -- 载多个文件
	// LoadHTMLFiles -- 精确加载指定的 HTML 模板文件
	engine.LoadHTMLGlob("templates/*")
	// engine.LoadHTMLFiles("templates/index.html")


	// 静态文件
	// Static -- 静态文件服务
	// 参数1：URL 前缀
	// 参数2：文件系统路径
	engine.Static("/static", "./static")

	// 提供单个静态文件
	// StaticFile -- 提供单个静态文件服务
	// 参数1：URL 路径
	// 参数2：文件系统路径
	// engine.StaticFile("/favicon.ico", "./static/favicon.ico")


	engine.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "QX-hao的第一个 Gin 项目",
		})
	})

	engine.Run(":1234")

}