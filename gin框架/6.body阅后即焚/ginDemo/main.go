package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"bytes"
	"fmt"
)

func main()  {
	
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()

	engine.GET("/body", func(c *gin.Context) {
		// 读取body后，body就会被清空
		body, _ := io.ReadAll(c.Request.Body)
			c.JSON(200, gin.H{
				"body": string(body),
			})

		fmt.Print(string(body))
		
		// 这里就拿不到body的name信息了，因为body已经被读取了
		name := c.PostForm("name")
		c.JSON(200, gin.H{
			"name": name,
		})

		fmt.Println("请求头：",c.Request.Header)

		// 重新设置body后，就可以从body中读取name信息了
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		// 这里就可以拿到body的name信息了
		name = c.PostForm("name")
		c.JSON(200, gin.H{
			"name": name,
		})


	})

	engine.Run(":1234")
}
