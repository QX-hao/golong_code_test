# Go语言time包使用指南

## 概述

`time`包标准库中非常重要的一个包。包含了时间获取、格式化、转换、定时器等多种功能。

- **时间获取和格式化**：支持多种时间格式
- **时间戳转换**：秒级和纳秒级时间戳互转
- **定时器功能**：Ticker和Timer满足不同定时需求
- **时区处理**：支持全球时区转换
- **时间运算**：加减、比较、差值计算

## 一、基础时间操作 (timedemo01)

### 1.1 获取当前时间

```go
timeObj := time.Now()
fmt.Println(timeObj) // 2025-10-25 19:10:19.9605226 +0800 CST m=+0.000000001
```

### 1.2 提取时间组件

```go
year := timeObj.Year()
month := timeObj.Month()
day := timeObj.Day()
hour := timeObj.Hour()
minute := timeObj.Minute()
second := timeObj.Second()
```

### 1.3 时间格式化

Go语言使用特殊的格式化字符串：
- **2006** - 年份
- **01** - 月份
- **02** - 日期
- **15** - 24小时制小时
- **03** - 12小时制小时
- **04** - 分钟
- **05** - 秒

```go
// 24小时制
fmt.Println("24小时制：", timeObj.Format("2006-01-02 15:04:05"))

// 12小时制
fmt.Println("12小时制：", timeObj.Format("2006-01-02 03:04:05"))
```

### 1.4 时间戳操作

#### 获取时间戳
```go
// 秒级时间戳
timestamp := timeObj.Unix()

// 纳秒级时间戳
timestampNano := timeObj.UnixNano()
```

#### 时间戳转时间
```go
// 秒级时间戳转时间
timeObj01 := time.Unix(int64(timestamp), 0)

// 纳秒级时间戳转时间
timeObj02 := time.Unix(0, int64(timestampNano))

// 同时指定秒和纳秒（会自动相加）
timeObj03 := time.Unix(int64(timestamp), int64(timestampNano))
```

### 1.5 字符串转时间

```go
str := "2025-10-25 19:10:19"
timeObj04, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
```

## 二、定时器使用 (timedemo02)

### 2.1 Ticker定时器

Ticker提供周期性定时功能：

```go
// 创建每秒触发一次的定时器
ticker := time.NewTicker(time.Second)

// 重要：使用后需要停止，避免goroutine泄漏
defer ticker.Stop()
```

### 2.2 定时器循环

#### 方法一：使用range循环
```go
num := 5
for value := range ticker.C {
    fmt.Printf("当前时间：%v\n", key)
    num--
    if num <= 0 {
        ticker.Stop()
        break
    }
}
```

#### 方法二：使用无限循环
```go
for {
    <-ticker.C
    fmt.Println("ticker")
}
```

### 2.3 定时器通道机制

- `ticker.C` 是一个时间通道（channel）
- `<-ticker.C` 会阻塞当前goroutine，直到定时器触发
- 每次触发时，通道会发送一个时间值

### 2.4 Sleep功能

```go
fmt.Printf("程序开始%v\n", time.Now())
time.Sleep(time.Second * 5)  // 休眠5秒
fmt.Printf("程序结束%v\n", time.Now())
```

## 三、常用时间常量

```go
time.Second      // 1秒
time.Minute      // 1分钟
time.Hour        // 1小时
time.Millisecond // 1毫秒
time.Microsecond // 1微秒
time.Nanosecond  // 1纳秒
```

## 四、时间运算

### 4.1 时间加减
```go
now := time.Now()

// 加1小时
future := now.Add(time.Hour)

// 减30分钟
past := now.Add(-30 * time.Minute)
```

### 4.2 时间比较
```go
time1 := time.Now()
time2 := time1.Add(time.Hour)

fmt.Println(time1.Before(time2))  // true
fmt.Println(time1.After(time2))   // false
equal := time1.Equal(time2)      // false
```

### 4.3 时间差计算
```go
duration := time2.Sub(time1)
fmt.Println(duration)  // 1h0m0s
```

## 五、时区处理

### 5.1 本地时区
```go
localTime := time.Now()
fmt.Println("本地时间:", localTime)
```

### 5.2 指定时区
```go
// 加载时区
loc, _ := time.LoadLocation("America/New_York")
newYorkTime := time.Now().In(loc)
fmt.Println("纽约时间:", newYorkTime)
```

## 六、最佳实践

### 6.1 定时器使用注意事项
1. **必须调用Stop()**：避免goroutine泄漏
2. **合理选择定时器类型**：
   - `Ticker`：周期性任务
   - `Timer`：一次性延迟任务
   - `Sleep`：简单延迟

### 6.2 性能考虑
1. 避免频繁创建和销毁定时器
2. 对于高精度定时需求，考虑使用`time.Tick`
3. 注意时区转换的性能开销

### 6.3 错误处理
```go
// 时间解析时的错误处理
parsedTime, err := time.Parse("2006-01-02", "2025-10-25")
if err != nil {
    fmt.Println("时间解析错误:", err)
    return
}
```
