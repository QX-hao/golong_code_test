# 【Go】-- Gin框架基本使用

## 第一个Gin应用

### 代码示例

```go
package main

import "github.com/gin-gonic/gin"

// 首页处理函数
func index(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello, world!",
    })
}

func main() {
    // 忽略gin的日志（生产模式）
    gin.SetMode(gin.ReleaseMode)
    
    // 初始化Gin引擎
    r := gin.Default()
    
    // 定义路由
    r.GET("/login", index)
    
    // 启动服务器，监听1234端口
    r.Run("0.0.0.0:1234")
}
```

### 核心概念

1. **gin.H**：Gin框架提供的map[string]interface{}别名，用于构建JSON响应
2. **gin.Context**：请求上下文，包含请求和响应信息
3. **路由定义**：使用`GET`、`POST`等方法定义HTTP路由
4. **服务器启动**：`Run()`方法启动HTTP服务器

## 路由定义和处理

### GET请求处理

```go
engine.GET("/get_test", func(c *gin.Context){
    c.JSON(200, gin.H{
        "code": 200,
        "msg": "GET 请求成功",
    })
})
```

### POST请求处理（JSON数据绑定）

```go
engine.POST("/post_test", func(c *gin.Context){
    var user user.User
    
    // JSON数据绑定到结构体
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
```

### 数据结构定义

```go
package user

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
```

### 关键特性

1. **数据绑定**：`ShouldBindJSON()`自动将JSON数据绑定到结构体
2. **错误处理**：验证失败时返回详细的错误信息
3. **响应格式**：统一的JSON响应格式（code, msg, data）

## 模板渲染和静态文件

### 模板配置

```go
// 加载HTML模板
engine.LoadHTMLGlob("templates/*")
// 或者精确加载指定文件
// engine.LoadHTMLFiles("templates/index.html")
```

### 静态文件服务

```go
// 静态文件服务配置
engine.Static("/static", "./static")

// 提供单个静态文件
// engine.StaticFile("/favicon.ico", "./static/favicon.ico")
```

### 模板渲染

```go
engine.GET("/index", func(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{
        "title": "QX-hao的第一个 Gin 项目",
    })
})
```

### 路径映射关系

| HTML引用路径 | 实际文件路径 |
|-------------|-------------|
| `/static/css/style.css` | `./static/css/style.css` |
| `/static/js/main.js` | `./static/js/main.js` |
| `/static/images/logo.png` | `./static/images/logo.png` |

### 模板语法示例

```html
<!-- 使用模板变量 -->
<title>{{.title}}</title>

<!-- 引用静态资源 -->
<link rel="stylesheet" href="/static/css/style.css">
<script src="/static/js/main.js"></script>
<img src="/static/images/logo.png" alt="Logo">
```

## 错误处理和日志记录

### 全局错误处理中间件

```go
router.Use(func(c *gin.Context) {
    // 执行后续中间件和路由处理
    c.Next()

    // 检查是否有错误发生
    if len(c.Errors) > 0 {
        // 自定义错误处理
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "服务器内部错误"
        })
    }
})
```

### 错误模拟和记录

```go
router.GET("/ping", func(c *gin.Context) {
    // 模拟处理过程中发生错误
    c.Error(gin.Error{Err: errors.New("处理过程中发生错误")})
})
```

### 错误处理流程

1. **请求到达**：浏览器发送请求
2. **执行中间件**：先执行自定义错误处理中间件
3. **路由处理**：执行具体的路由处理函数
4. **错误检查**：检查`c.Errors`中是否有错误
5. **错误响应**：如果有错误，返回统一的错误格式

### 日志配置

```go
// 生产模式（减少日志输出）
gin.SetMode(gin.ReleaseMode)

// 开发模式（详细日志）
// gin.SetMode(gin.DebugMode) // 默认模式
```

## 核心API总结

### 路由方法

| 方法 | 说明 | 示例 |
|------|------|------|
| `GET` | 处理GET请求 | `r.GET("/path", handler)` |
| `POST` | 处理POST请求 | `r.POST("/path", handler)` |
| `PUT` | 处理PUT请求 | `r.PUT("/path", handler)` |
| `DELETE` | 处理DELETE请求 | `r.DELETE("/path", handler)` |

### 响应方法

| 方法 | 说明 | 示例 |
|------|------|------|
| `c.JSON()` | 返回JSON响应 | `c.JSON(200, gin.H{"msg": "ok"})` |
| `c.HTML()` | 返回HTML页面 | `c.HTML(200, "index.html", data)` |
| `c.String()` | 返回文本响应 | `c.String(200, "Hello")` |
| `c.Redirect()` | 重定向 | `c.Redirect(302, "/new-path")` |

### 数据绑定

| 方法 | 说明 | 示例 |
|------|------|------|
| `ShouldBindJSON()` | JSON数据绑定 | `c.ShouldBindJSON(&user)` |
| `ShouldBindQuery()` | Query参数绑定 | `c.ShouldBindQuery(&params)` |
| `ShouldBindUri()` | URI参数绑定 | `c.ShouldBindUri(&params)` |

