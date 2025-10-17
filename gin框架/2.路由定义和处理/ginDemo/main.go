package main

import (
	"github.com/gin-gonic/gin"
	"ginDemo/user"
)

func main()  {
	
	engine := gin.Default()
	
	// GET 请求处理
	engine.GET("/get_test", func(c *gin.Context){
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "GET 请求成功",
		})
	})
	
	// POST 请求处理
	engine.POST("/post_test", func(c *gin.Context){
		var user user.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg": "POST 请求参数错误",
				"error": err.Error(),
			})
			return 
		}
		
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "POST 请求成功",
			"data": gin.H{
				"username": user.Username,
				"password": user.Password,
			},
		})
	})

	engine.Run(":1234")
}