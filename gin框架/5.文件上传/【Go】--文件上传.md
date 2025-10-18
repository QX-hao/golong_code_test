# 【Go】-- Gin框架文件上传



## 核心特性

- **单文件上传**：处理单个文件上传请求
- **多文件上传**：处理多个文件同时上传
- **文件验证**：获取文件名、文件大小等信息
- **文件保存**：提供多种文件保存方式



## 单文件上传

### 核心代码

```go
r.POST("/user", func(c *gin.Context) {
    // FormFile 方法会解析 multipart form 数据，包括文件上传
    fileheader, err := c.FormFile("file")
    if err != nil {
        c.JSON(500, gin.H{
            "msg": err.Error(),
        })
        return
    }
    
    // 获取文件名和文件大小
    c.JSON(200, gin.H{
        "文件名": fileheader.Filename,
        "文件大小": fileheader.Size,
    })
    
    // 保存文件
    c.SaveUploadedFile(fileheader, "images/"+fileheader.Filename)
})
```

### 关键API

- **`c.FormFile("file")`**：获取单个上传文件
- **`fileheader.Filename`**：获取原始文件名
- **`fileheader.Size`**：获取文件大小（字节）
- **`c.SaveUploadedFile()`**：保存上传的文件

## 多文件上传

### 核心代码

```go
r.POST("/user2", func(c *gin.Context) {
    // 获取整个表单
    form, err := c.MultipartForm()
    if err != nil {
        c.JSON(500, gin.H{
            "msg": err.Error(),
        })
        return
    }
    
    // 获取全部文件
    for key, fileheader := range form.File["file"] {
        // 打印文件名
        c.JSON(200, gin.H{
            "key": key,
            "文件名": fileheader.Filename,
        })
        c.SaveUploadedFile(fileheader, "images/"+fileheader.Filename)
    }
})
```

### 关键API

- **`c.MultipartForm()`**：获取完整的multipart表单数据
- **`form.File["file"]`**：获取指定字段的所有文件
- **循环处理**：遍历处理每个上传的文件

## 文件保存方法

Gin框架提供了两种文件保存方式：

### 方法一：手动读取和保存

```go
// 方法一：手动读取文件内容并保存
file, _ := fileheader.Open()
byteData, _ := io.ReadAll(file)
os.WriteFile(fileheader.Filename, byteData, 0644)
```

**优点**：
- 完全控制文件处理过程
- 可以在保存前对文件内容进行处理

**缺点**：
- 代码相对复杂
- 需要手动处理文件流

### 方法二：使用内置方法保存

```go
// 方法二：使用内置方法保存
c.SaveUploadedFile(fileheader, "images/"+fileheader.Filename)
```

**优点**：
- 代码简洁，一行代码完成保存
- 自动处理文件流和错误
- 推荐使用的方式

**缺点**：
- 灵活性相对较低

## 完整代码示例

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 单文件上传接口
	r.POST("/user", func(c *gin.Context) {
		// FormFile 方法会解析 multipart form 数据，包括文件上传
		fileheader, err := c.FormFile("file")
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err.Error(),
			})
			return
		}
		
		// 获取文件名和文件大小
		c.JSON(200, gin.H{
			"文件名": fileheader.Filename,
			"文件大小": fileheader.Size,
		})
		
		// 保存文件到images目录
		c.SaveUploadedFile(fileheader, "images/"+fileheader.Filename)
	})

	// 多文件上传接口
	r.POST("/user2", func(c *gin.Context) {
		// 获取整个表单
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err.Error(),
			})
			return
		}
		
		// 获取全部文件
		for key, fileheader := range form.File["file"] {
			// 打印文件名
			c.JSON(200, gin.H{
				"key": key,
				"文件名": fileheader.Filename,
			})
			// 保存每个文件
			c.SaveUploadedFile(fileheader, "images/"+fileheader.Filename)
		}
	})

	r.Run(":1234")
}
```

## 使用说明

### 1. 启动服务器

```bash
cd 5.文件上传/ginDemo
go run main.go
```

服务器将在 `http://localhost:1234` 启动

### 2. 单文件上传测试

使用Postman或curl测试单文件上传：

**请求信息：**
- 方法：POST
- URL：`http://localhost:1234/user`
- Body类型：form-data
- 参数：`file`（文件类型）

**响应示例：**
```json
{
    "文件名": "example.jpg",
    "文件大小": 102400
}
```

### 3. 多文件上传测试

**请求信息：**
- 方法：POST
- URL：`http://localhost:1234/user2`
- Body类型：form-data
- 参数：`file`（文件类型，可上传多个）

**响应示例：**
```json
{
    "key": 0,
    "文件名": "file1.jpg"
}
{
    "key": 1,
    "文件名": "file2.png"
}
```
