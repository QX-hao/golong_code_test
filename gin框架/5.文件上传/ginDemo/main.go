package main

import (
	// "fmt"
	// "io"
	// "os"
	"github.com/gin-gonic/gin"
)

func main()  {
	
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		// FormFile 方法会解析 multipart form 数据，包括文件上传
		fileheader,err := c.FormFile("file")
		if err != nil {
			c.JSON(500,gin.H{
				"msg":err.Error(),
			})
			return
		}
		
		// 获取文件名和文件大小
		c.JSON(200,gin.H{
			"文件名":fileheader.Filename,
			"文件大小":fileheader.Size,
		})
		
		// 方法一
		// file, _ := fileheader.Open()
		// byteData, _ := io.ReadAll(file)
		// os.WriteFile(fileheader.Filename,byteData,0644)

		// 方法二
		c.SaveUploadedFile(fileheader,"images/"+fileheader.Filename)
	})

	r.POST("/user2",func(c *gin.Context){
		// 获取整个表单
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500,gin.H{
				"msg":err.Error(),
			})
			return
		}
		
		// 获取全部文件
		for key,fileheader := range form.File["file"]{
			// 打印文件名
			c.JSON(200,gin.H{
				"key":key,
				"文件名":fileheader.Filename,
			})
			c.SaveUploadedFile(fileheader,"images/"+fileheader.Filename)
		}
	})

	r.Run(":1234")


}