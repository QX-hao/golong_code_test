# 【Go】--闭包

## 什么是闭包

### 1.1 闭包

闭包（Closure）是函数和其相关引用环境组合而成的实体。在Go语言中，闭包是一个函数值，它引用了其函数体之外的变量。

**核心特征：**
- 函数可以访问其词法作用域外的变量
- 被引用的变量在函数调用期间保持存活
- 闭包可以"记住"创建时的环境状态

### 1.2 闭包与普通函数的区别

| 特性 | 普通函数 | 闭包 |
|------|----------|------|
| 变量作用域 | 只能访问参数和局部变量 | 可以访问外部作用域变量 |
| 状态保持 | 无状态，每次调用独立 | 有状态，可以记住上次调用的状态 |
| 内存管理 | 调用结束后释放局部变量 | 外部变量在闭包存在期间保持存活 |

## 2. 闭包实现原理

### 2.1 内存布局

```go
func outer() func() int {
    x := 10  // 外部变量
    return func() int {
        x++   // 闭包引用外部变量
        return x
    }
}
```

**内存结构示意图：**
```
栈内存 (Stack)
┌─────────────┐
│ outer函数帧  │
│ x = 10      │ ← 闭包捕获
└─────────────┘

堆内存 (Heap)
┌─────────────┐
│ 闭包结构体   │
│ 函数指针     │ → func() int
│ 捕获变量指针 │ → &x
└─────────────┘
```

### 2.2 变量捕获机制

Go编译器在遇到闭包时：
1. **分析变量引用**：确定哪些外部变量被闭包引用
2. **变量逃逸分析**：将被引用的变量分配到堆内存
3. **创建闭包结构**：生成包含函数指针和捕获变量指针的结构体

## 3. 闭包语法与示例

### 3.1 基础闭包示例

```go
package main

import "fmt"

func main() {
    // 示例1：简单的计数器
    counter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }()
    
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
}
```

### 3.2 带参数的闭包工厂

```go
// 创建配置器闭包
func createConfigurator(initialValue int) func(int) int {
    current := initialValue
    return func(delta int) int {
        current += delta
        return current
    }
}

func main() {
    config := createConfigurator(100)
    fmt.Println(config(10))  // 110
    fmt.Println(config(-5))  // 105
    fmt.Println(config(20))  // 125
}
```

### 3.3 多变量闭包

```go
func createBankAccount(initialBalance float64) func(float64) (float64, bool) {
    balance := initialBalance
    
    return func(amount float64) (float64, bool) {
        if balance + amount < 0 {
            return balance, false // 余额不足
        }
        balance += amount
        return balance, true
    }
}

func main() {
    account := createBankAccount(1000.0)
    
    balance, success := account(500.0)
    fmt.Printf("存款后余额: %.2f, 成功: %v\n", balance, success) // 1500.00, true
    
    balance, success = account(-2000.0)
    fmt.Printf("取款后余额: %.2f, 成功: %v\n", balance, success) // 1500.00, false
}
```

## 4. 典型应用场景

### 4.1 数据封装（信息隐藏）

```go
// 创建私有数据封装
func createPrivateData() (get func() string, set func(string)) {
    privateData := "初始数据"
    
    get = func() string {
        return privateData
    }
    
    set = func(newData string) {
        // 可以添加验证逻辑
        if len(newData) > 0 {
            privateData = newData
        }
    }
    
    return get, set
}

func main() {
    getData, setData := createPrivateData()
    
    fmt.Println("初始数据:", getData()) // 初始数据: 初始数据
    
    setData("修改后的数据")
    fmt.Println("修改后数据:", getData()) // 修改后数据: 修改后的数据
    
    // 外部无法直接访问privateData变量，实现了数据封装
}
```

### 4.2 函数工厂（动态生成函数）

```go
// 创建不同级别的日志函数
func createLogger(level string) func(string) {
    prefix := fmt.Sprintf("[%s]", level)
    
    return func(message string) {
        fmt.Printf("%s %s\n", prefix, message)
    }
}

func main() {
    infoLog := createLogger("INFO")
    warnLog := createLogger("WARN")
    errorLog := createLogger("ERROR")
    
    infoLog("应用程序启动")   // [INFO] 应用程序启动
    warnLog("内存使用较高")   // [WARN] 内存使用较高
    errorLog("数据库连接失败") // [ERROR] 数据库连接失败
}
```

### 4.3 延迟执行（回调函数）

```go
// 创建延迟任务执行器
func createDelayedExecutor(delay time.Duration) func(func()) {
    return func(task func()) {
        time.AfterFunc(delay, task)
    }
}

func main() {
    delayedExec := createDelayedExecutor(2 * time.Second)
    
    fmt.Println("任务提交时间:", time.Now().Format("15:04:05"))
    
    delayedExec(func() {
        fmt.Println("任务执行时间:", time.Now().Format("15:04:05"))
        fmt.Println("延迟任务执行完成!")
    })
    
    // 等待任务执行
    time.Sleep(3 * time.Second)
}
```

### 4.4 中间件模式

```go
// HTTP中间件工厂
func createMiddleware(name string) func(http.HandlerFunc) http.HandlerFunc {
    return func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            fmt.Printf("[%s] 开始处理请求: %s\n", name, r.URL.Path)
            
            next(w, r) // 调用下一个处理程序
            
            fmt.Printf("[%s] 请求处理完成，耗时: %v\n", name, time.Since(start))
        }
    }
}

func main() {
    // 模拟HTTP处理
    handler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("处理业务逻辑")
        time.Sleep(100 * time.Millisecond) // 模拟处理时间
    }
    
    // 应用多个中间件
    loggingMiddleware := createMiddleware("日志")
    timingMiddleware := createMiddleware("计时")
    
    finalHandler := loggingMiddleware(timingMiddleware(handler))
    
    // 模拟请求处理
    fmt.Println("=== 模拟HTTP请求处理 ===")
    finalHandler(nil, &http.Request{URL: &url.URL{Path: "/api/test"}})
}
```

### 4.5 状态机实现

```go
// 创建状态机
func createStateMachine(initialState string) (getState func() string, transition func(string) bool) {
    currentState := initialState
    
    // 状态转移规则
    transitions := map[string][]string{
        "start":    {"running"},
        "running":  {"paused", "stopped"},
        "paused":   {"running", "stopped"},
        "stopped":  {"start"},
    }
    
    getState = func() string {
        return currentState
    }
    
    transition = func(newState string) bool {
        // 检查状态转移是否合法
        validTransitions, exists := transitions[currentState]
        if !exists {
            return false
        }
        
        for _, validState := range validTransitions {
            if validState == newState {
                currentState = newState
                return true
            }
        }
        return false
    }
    
    return getState, transition
}

func main() {
    getState, transition := createStateMachine("start")
    
    fmt.Println("初始状态:", getState()) // start
    
    if transition("running") {
        fmt.Println("状态转移成功:", getState()) // running
    }
    
    if !transition("start") {
        fmt.Println("非法状态转移: running → start") // 非法状态转移
    }
}
```

## 5. 闭包陷阱与最佳实践

### 5.1 常见陷阱

#### 5.1.1 循环变量捕获问题

```go
// 问题代码：所有闭包共享同一个i
func problematicLoop() {
    var funcs []func()
    
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() {
            fmt.Println(i) // 所有闭包都输出3
        })
    }
    
    for _, f := range funcs {
        f() // 输出: 3, 3, 3
    }
}

// 解决方案1：创建局部变量副本
func solution1() {
    var funcs []func()
    
    for i := 0; i < 3; i++ {
        j := i // 创建副本
        funcs = append(funcs, func() {
            fmt.Println(j) // 正确输出: 0, 1, 2
        })
    }
    
    for _, f := range funcs {
        f()
    }
}

// 解决方案2：通过参数传递
func solution2() {
    var funcs []func()
    
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func(x int) func() {
            return func() {
                fmt.Println(x) // 正确输出: 0, 1, 2
            }
        }(i))
    }
    
    for _, f := range funcs {
        f()
    }
}
```

#### 5.1.2 内存泄漏风险

```go
// 可能造成内存泄漏的闭包
func createMemoryLeak() func() {
    largeData := make([]byte, 100*1024*1024) // 100MB数据
    
    return func() {
        // 即使不再需要largeData，由于闭包引用，它不会被GC回收
        fmt.Println("闭包执行")
    }
}

// 解决方案：及时释放引用
func createSafeClosure() func() {
    largeData := make([]byte, 100*1024*1024)
    
    // 使用完成后显式释放
    closure := func() {
        fmt.Println("闭包执行")
    }
    
    // 如果不再需要largeData，可以设置为nil
    // largeData = nil
    
    return closure
}
```

### 5.2 最佳实践

1. **明确变量作用域**：清晰区分局部变量和捕获变量
2. **避免过度捕获**：只捕获必要的变量
3. **及时释放资源**：对于大型数据，使用后及时释放引用
4. **文档化闭包行为**：为复杂的闭包添加注释说明
5. **测试闭包状态**：确保闭包在不同调用间状态正确

## 6. 性能考虑

### 6.1 性能影响分析

闭包相比普通函数的性能开销：
- **内存开销**：每个闭包需要额外的内存存储捕获变量
- **创建开销**：闭包创建比普通函数稍慢
- **调用开销**：闭包调用与普通函数调用性能相当

### 6.2 优化建议

```go
// 优化前：频繁创建闭包
func unoptimized() {
    for i := 0; i < 1000; i++ {
        closure := func() {
            // 处理逻辑
        }
        closure()
    }
}

// 优化后：复用闭包
func optimized() {
    // 在循环外创建闭包
    closure := func() {
        // 处理逻辑
    }
    
    for i := 0; i < 1000; i++ {
        closure() // 复用同一个闭包
    }
}
```

### 6.3 基准测试示例

```go
func BenchmarkClosure(b *testing.B) {
    counter := 0
    closure := func() {
        counter++
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        closure()
    }
}

func BenchmarkRegularFunction(b *testing.B) {
    counter := 0
    regularFunc := func(c *int) {
        *c++
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        regularFunc(&counter)
    }
}
```
