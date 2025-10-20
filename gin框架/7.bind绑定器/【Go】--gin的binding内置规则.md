# 【Go】--gin的binding内置规则

## binding标签基本语法

```go
type User struct {
    FieldName Type `binding:"规则1,规则2,规则3"`
}
```

## 内置验证规则详解

### 1. 必填验证

```go
// required - 字段必须提供且不能为空
Name string `binding:"required"`
```

**验证场景：**
- 字符串：长度 > 0
- 数字：不为0
- 数组/切片：长度 > 0
- 指针：不为nil

### 2. 数字验证

```go
// number - 必须是有效数字
Age int `binding:"number"`

// min/max - 数值范围限制
Age int `binding:"min=18,max=100"`
Score float64 `binding:"min=0,max=100"`
```

### 3. 字符串验证

```go
// len - 固定长度
Code string `binding:"len=6"`

// min/max - 长度范围
Username string `binding:"min=3,max=20"`

// eqfield - 字段相等验证
Password string `binding:"required"`
ConfirmPassword string `binding:"eqfield=Password"`
```

### 4. 格式验证

```go
// email - 邮箱格式
Email string `binding:"required,email"`

// url - URL格式
Website string `binding:"url"`

// ip - IP地址格式
IPAddress string `binding:"ip"`

// ipv4 - IPv4地址格式
IPv4 string `binding:"ipv4"`

// ipv6 - IPv6地址格式
IPv6 string `binding:"ipv6"`
```

### 5. 自定义格式验证

```go
// contains - 包含特定字符串
Domain string `binding:"contains=.com"`

// excludes - 不包含特定字符串
Content string `binding:"excludes=script"`

// alpha - 只包含字母
FirstName string `binding:"alpha"`

// alphanum - 字母和数字
Username string `binding:"alphanum"`

// numeric - 只包含数字
Phone string `binding:"numeric"`

// hexadecimal - 十六进制格式
HexColor string `binding:"hexadecimal"`

// base64 - Base64格式
Base64Data string `binding:"base64"`
```

### 6. 日期时间验证

```go
// datetime - 日期时间格式
Birthday string `binding:"datetime=2006-01-02"`
CreatedAt string `binding:"datetime=2006-01-02 15:04:05"`
```

### 7. 文件验证

```go
// 文件大小限制（字节）
FileSize int64 `binding:"max=10485760"` // 10MB

// 文件类型验证
FileType string `binding:"oneof=jpg png gif pdf"`
```

## 组合验证规则

### 常用组合示例

```go
type User struct {
    // 用户注册验证
    Username    string `binding:"required,min=3,max=20,alphanum"`
    Email       string `binding:"required,email"`
    Password    string `binding:"required,min=8,max=50"`
    Age         int    `binding:"required,min=1,max=3"`
    Phone       string `binding:"required,numeric,len=11"`
    BirthDate   string `binding:"datetime=2006-01-02"`
}

type Product struct {
    // 商品信息验证
    Name        string  `binding:"required,min=2,max=100"`
    Price       float64 `binding:"required,min=0"`
    Stock       int     `binding:"required,min=0"`
    Category    string  `binding:"required,oneof=electronics clothing books"`
    Description string `binding:"max=500"`
}
```

## 不同绑定类型的标签

### 1. 表单绑定 (Form Data)

```go
type LoginForm struct {
    Username string `form:"username" binding:"required"`
    Password string `form:"password" binding:"required"`
}
```

### 2. JSON绑定

```go
type User struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
}
```

### 3. URI参数绑定

```go
type Params struct {
    ID   int    `uri:"id" binding:"required,number"`
    Name string `uri:"name" binding:"required"`
}
```

### 4. Query参数绑定

```go
type Query struct {
    Page     int    `form:"page" binding:"min=1"`
    PageSize int    `form:"page_size" binding:"min=1,max=100"`
    Keyword  string `form:"keyword"`
}
```

## 完整代码示例

### 示例1：用户注册验证

```go
package main

import "github.com/gin-gonic/gin"

type RegisterRequest struct {
    Username        string `json:"username" binding:"required,min=3,max=20,alphanum"`
    Email           string `json:"email" binding:"required,email"`
    Password        string `json:"password" binding:"required,min=8,max=50"`
    ConfirmPassword string `json:"confirm_password" binding:"eqfield=Password"`
    Age             int    `json:"age" binding:"required,min=18,max=100"`
    Phone           string `json:"phone" binding:"required,numeric,len=11"`
}

func main() {
    r := gin.Default()
    
    r.POST("/register", func(c *gin.Context) {
        var req RegisterRequest
        
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
                "code": 400,
            })
            return
        }
        
        c.JSON(200, gin.H{
            "message": "注册成功",
            "data": req,
        })
    })
    
    r.Run(":8080")
}
```

### 示例2：商品搜索验证

```go
type SearchRequest struct {
    Keyword  string `form:"keyword" binding:"max=100"`
    Category string `form:"category" binding:"oneof=all electronics clothing books"`
    MinPrice float64 `form:"min_price" binding:"min=0"`
    MaxPrice float64 `form:"max_price" binding:"min=0"`
    Page     int    `form:"page" binding:"min=1"`
    PageSize int    `form:"page_size" binding:"min=1,max=50"`
}

func searchHandler(c *gin.Context) {
    var req SearchRequest
    
    if err := c.ShouldBindQuery(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // 处理搜索逻辑
    c.JSON(200, gin.H{"data": "搜索结果"})
}
```

## 错误处理

### 获取详细的错误信息

```go
if err := c.ShouldBind(&req); err != nil {
    // 获取字段级别的错误
    if fieldErrors, ok := err.(validator.ValidationErrors); ok {
        for _, fieldError := range fieldErrors {
            fmt.Printf("字段: %s, 错误: %s, 值: %v\n", 
                fieldError.Field(), 
                fieldError.Tag(), 
                fieldError.Value())
        }
    }
    
    c.JSON(400, gin.H{"error": err.Error()})
    return
}
```

### 自定义错误消息

```go
import "github.com/go-playground/validator/v10"

// 注册自定义验证器
if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    v.RegisterValidation("custom_rule", func(fl validator.FieldLevel) bool {
        // 自定义验证逻辑
        return true
    })
}
```

## 最佳实践

### 1. 验证规则设计原则
- **前端友好**：提供清晰的错误信息
- **安全性**：防止SQL注入和XSS攻击
- **性能**：避免过于复杂的验证规则
- **可读性**：规则命名清晰易懂

### 2. 验证顺序建议
1. 格式验证（email, url, ip等）
2. 长度验证（min, max, len）
3. 内容验证（contains, excludes等）
4. 业务逻辑验证（eqfield等）

### 3. 性能优化
- 避免在热路径中使用复杂验证
- 对频繁验证的结果进行缓存
- 使用适当的验证粒度

## 调试技巧

### 1. 查看验证错误详情

```go
import "github.com/go-playground/validator/v10"

if err := c.ShouldBind(&req); err != nil {
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        for _, ve := range validationErrors {
            fmt.Printf("字段: %s, 标签: %s, 参数: %s, 值: %v\n",
                ve.Field(), ve.Tag(), ve.Param(), ve.Value())
        }
    }
}
```

### 2. 测试验证规则

```go
// 单元测试验证规则
func TestValidation(t *testing.T) {
    user := User{
        Username: "test",
        Email:    "invalid-email",
    }
    
    err := validate.Struct(user)
    if err == nil {
        t.Error("预期验证失败")
    }
}
```