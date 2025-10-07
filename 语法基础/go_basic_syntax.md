# Go语言基础语法指南

## 1. 程序结构

### 1.1 包声明
```go
package main  // 每个Go程序必须有一个main包
```

### 1.2 导入语句
```go
import "fmt"  // 导入标准库fmt包
import (
    "os"
    "math"
)
```

## 2. 基本语法

### 2.1 变量声明

#### 单变量声明
```go
// 标准声明格式
var 变量名 类型 = 值

// 示例
var a int = 10       // 显式类型声明
var b = 20          // 类型推断(编译器自动推断为int)
c := 30             // 短变量声明(只能在函数内使用)
var d string        // 默认初始化为空字符串""
```

#### 多变量声明
```go
// 相同类型变量声明
var x, y int = 10, 20

// 不同类型变量声明(使用var块)
var (
    name string = "John"
    age  int    = 30
    isActive bool = true
)

// 短变量多声明(只能在函数内使用)
a, b := 1, "hello"

// 变量交换(Go特有语法)
a, b = b, a  // 无需临时变量

// 声明并忽略某些返回值
_, err := someFunction()  // 使用_忽略第一个返回值
```

### 2.2 常量
```go
const Pi = 3.14
```

### 2.3 控制结构

#### if语句
```go
if x > 0 {
    fmt.Println("正数")
}
```

#### for循环
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

## 3. 函数

### 3.1 函数定义
```go
func add(a int, b int) int {
    return a + b
}
```

### 3.2 多返回值
```go
func swap(x, y string) (string, string) {
    return y, x
}
```

## 4. 数据类型

### 4.1 基本类型
- int, float64, bool, string等

### 4.2 复合类型
- 数组
- 切片
- 映射(map)
- 结构体

## 5. 接口

```go
type Shape interface {
    Area() float64
}
```

## 6. 并发

### 6.1 goroutine
```go
go func() {
    fmt.Println("并发执行")
}()
```

### 6.2 channel
```go
ch := make(chan int)
```