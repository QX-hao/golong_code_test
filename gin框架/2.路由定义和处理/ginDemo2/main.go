package main

import (
	"ginDemo2/user"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()

	// 参数化路由
	engine.GET("/user/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "GET 请求成功",
			"data": gin.H{
				"username": username,
			},
		})
	})

	// 路由组
	v1 := engine.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.String(200, "用户列表")
		})
		v1.POST("/users", func(c *gin.Context) {
			c.String(200, "创建用户")
		})
		v1.PUT("/users/:username", func(c *gin.Context) {
			username := c.Param("username")
			var user user.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(400, gin.H{
					"code":  400,
					"msg":   "PUT 请求参数错误",
					"error": err.Error(),
				})
				return
			}

			// 更新用户信息
			user.Username = username
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "PUT 请求成功",
				"data": gin.H{
					"username": user.Username,
				},
			})
		})
		v1.DELETE("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.String(200, "删除用户ID: %s", id)
		})
	}

	engine.Run(":1234")
}
