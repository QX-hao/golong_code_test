## 值类型 vs 引用类型



| 类型 | 分类 | 赋值行为 | 函数传参 |
|------|------|----------|----------|
| int, float, bool, string | 值类型 | 创建副本 | 传副本 |
| 数组 | 值类型 | 复制整个数组 | 传整个数组副本 |
| 结构体 | 值类型 | 复制整个结构体 | 传整个结构体副本 |
| 切片 | 引用类型 | 共享底层数组 | 共享底层数据 |
| 映射 | 引用类型 | 共享底层数据 | 共享底层数据 |
| 通道 | 引用类型 | 共享底层数据 | 共享底层数据 |
| 指针 | 值类型 | 复制指针地址 | 传指针副本 |

### 值类型 (Value Types)
**特点：** 变量直接存储值，赋值和传参时会创建副本

```go
package main

import "fmt"

func main() {
    // 值类型示例
    a := 10
    b := a  // 创建 a 的副本
    b = 20  // 修改 b 不会影响 a
    
    fmt.Println(a) // 10
    fmt.Println(b) // 20
}
```

### 引用类型 (Reference Types)  
**特点：** 变量存储的是内存地址，赋值和传参时共享底层数据

```go
func main() {
    // 引用类型示例
    slice1 := []int{1, 2, 3}
    slice2 := slice1  // 共享底层数组
    slice2[0] = 100   // 修改 slice2 会影响 slice1
    
    fmt.Println(slice1) // [100 2 3]
    fmt.Println(slice2) // [100 2 3]
}
```

## 具体分类

### 值类型包括：
```go
// 1. 基本数据类型
var i int
var f float64
var b bool
var s string

// 2. 数组
var arr [3]int

// 3. 结构体
type Person struct {
    Name string
    Age  int
}
var p Person

// 4. 指针（本身是值类型，但指向引用）
var ptr *int
```

### 引用类型包括：
```go
// 1. 切片 (slice)
var slice []int

// 2. 映射 (map)
var m map[string]int

// 3. 通道 (channel)
var ch chan int

// 4. 函数 (function)
var fn func()

// 5. 接口 (interface)
var err error
```

## 详细示例对比

### 值类型行为
```go
func modifyValue(x int) {
    x = 100  // 修改的是副本
}

func main() {
    num := 10
    modifyValue(num)
    fmt.Println(num) // 10，原值未改变
    
    arr1 := [3]int{1, 2, 3}
    arr2 := arr1     // 数组是值类型，会复制整个数组
    arr2[0] = 100
    fmt.Println(arr1) // [1 2 3]，原数组未变
    fmt.Println(arr2) // [100 2 3]
}
```

### 引用类型行为
```go
func modifySlice(s []int) {
    s[0] = 100  // 修改共享的底层数组
}

func main() {
    slice := []int{1, 2, 3}
    modifySlice(slice)
    fmt.Println(slice) // [100 2 3]，原切片被修改
    
    m1 := map[string]int{"a": 1}
    m2 := m1           // 映射是引用类型，共享底层数据
    m2["a"] = 100
    fmt.Println(m1)    // map[a:100]，原映射被修改
    fmt.Println(m2)    // map[a:100]
}
```

## 特殊情况：指针

```go
func modifyViaPointer(p *int) {
    *p = 100  // 通过指针修改原值
}

func main() {
    num := 10
    modifyViaPointer(&num)
    fmt.Println(num) // 100，原值被修改
    
    // 指针本身是值类型，但可以模拟引用行为
    var p1 *int = &num
    p2 := p1    // p2 是 p1 的副本，但指向同一个地址
    *p2 = 200
    fmt.Println(*p1) // 200
    fmt.Println(num) // 200
}
```
