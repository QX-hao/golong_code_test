package main

import (
	"github.com/gin-gonic/gin"
	// "io"
	// "bytes"
	"fmt"
)

func main() {

	engine := gin.Default()

	// 使用bind绑定参数
	engine.POST("/Bind", func(c *gin.Context) {
		type bind_user struct {
			// Name string
			// Age  int
			Name string `form:"name" json:"name" binding:"required"`
			Age  int    `form:"age" json:"age" binding:"required,number,max=3"`
		}

		var user bind_user

		// ShouldBind 会根据请求的 Content-Type 自动选择绑定器
		// bind返回值为error类型

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 打印
		c.JSON(200, gin.H{
			"user": user,
		})
		fmt.Println("请求体的Content-Type:", c.ContentType())
	})

	// get请求绑定参数
	engine.GET("/Bind/:name/:age", func(c *gin.Context) {
		type bind_user struct {
			// Name string
			// Age  int
			Name string `uri:"name" json:"name" binding:"required"`
			Age  int    `uri:"age" json:"age" binding:"required"`
		}

		var user bind_user

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 打印
		c.JSON(200, gin.H{
			"user": user,
		})
		fmt.Println("请求体的Content-Type:", c.ContentType())
	})
	engine.Run(":1234")
}
