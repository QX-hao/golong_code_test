# Go log模块使用

对于更复杂的日志需求，可以考虑使用第三方日志库如`zap`、`logrus`等，它们提供了更丰富的功能和更好的性能

### 1.主要特性
- 线程安全的日志记录
- 灵活的日志格式配置
- 支持多种输出目标（控制台、文件等）\- 内置错误处理和程序终止功能
- 支持自定义Logger创建

## 2. 基础日志输出

### 2.1 基本日志方法

log包提供了三种基本的日志输出方法：

```go
package main

import "log"

func main() {
    // Println - 自动添加换行符
    log.Println("这是一条日志信息")
    
    // Printf - 格式化输出
    str := "this is a test of log"
    log.Printf("%s", str)
    
    // Print - 不添加换行符
    log.Print("这条信息没有换行")
}
```

### 2.2 错误处理日志

log包还提供了在记录日志后终止程序的特殊方法：

```go
package main

import "log"

func main() {
    // Fatal系列函数会在日志输出后调用os.Exit(1)
    log.Fatalf("触发fatal错误，程序将退出")
    
    // Panic系列函数会在日志输出后调用panic()
    log.Panicln("触发panic错误")
}
```

**注意**：Fatal和Panic方法调用后程序会终止，后续代码不会执行。

## 3. 日志格式配置

### 3.1 设置日志标志

使用`SetFlags`方法可以配置日志的输出格式：

```go
package main

import "log"

func main() {
    // 设置日志格式标志
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.Println("设置完格式后的日志信息")
}
```

### 3.2 可用的日志标志

log包定义了以下常量来控制日志格式：

```go
const (
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123
    Llongfile                     // 文件全路径名+行号：/a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
```

## 4. 日志前缀配置

### 4.1 设置日志前缀

使用`SetPrefix`方法可以为日志添加前缀：

```go
package main

import "log"

func main() {
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    
    // 默认无前缀
    log.Println("没有设置前缀的日志信息")
    
    // 设置前缀
    log.SetPrefix("INFO: ")
    log.Println("设置完日志前缀的信息")
    
    // 获取当前前缀
    prefix := log.Prefix()
    log.Printf("当前前缀是：%s", prefix)
}
```

## 5. 文件日志输出

### 5.1 将日志输出到文件

使用`SetOutput`方法可以将日志重定向到文件：

```go
package main

import (
    "log"
    "os"
    "time"
)

// 全局文件变量
var logfile *os.File

func init() {
    // 创建日志目录
    os.Mkdir("log", 0755)
    
    // 打开或创建日志文件
    var err error
    logfile, err = os.OpenFile("log/"+time.Now().Format("2006-01-02")+".log", 
        os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("打开日志文件失败:%v", err)
        return
    }
    
    // 配置日志
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetOutput(logfile)
    log.SetPrefix("Prefix:")
}

func main() {
    // 确保文件关闭
    defer logfile.Close()
    
    log.Println("这是一条写入文件的日志信息")
}
```

### 5.2 文件打开模式说明

- `os.O_CREATE`：如果文件不存在则创建
- `os.O_WRONLY`：只写模式打开文件
- `os.O_APPEND`：追加模式，新内容添加到文件末尾
- `0666`：文件权限设置

## 6. 自定义Logger

### 6.1 使用log.New()创建自定义Logger

`log.New()`函数允许创建具有特定配置的Logger实例：

```go
package main

import (
    "log"
    "os"
    "time"
)

func main() {
    // 创建日志目录
    os.Mkdir("log", 0755)
    
    // 打开日志文件
    logfile, err := os.OpenFile("log/"+time.Now().Format("2006-01-02")+".log", 
        os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
        return
    }
    defer logfile.Close()
    
    // 创建自定义Logger
    logger := log.New(logfile, "new:", log.Ldate|log.Ltime|log.Lshortfile)
    logger.Println("使用New()来创造自己的logger对象")
}
```

### 6.2 多个Logger实例

可以创建多个Logger实例用于不同的日志目的：

```go
package main

import (
    "log"
    "os"
)

func main() {
    // 创建不同用途的Logger
    infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
    errorLogger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
    
    infoLogger.Println("这是一条信息日志")
    errorLogger.Println("这是一条错误日志")
}
```

## 7. 高级用法

### 7.1 日志级别管理

虽然log包本身不提供日志级别，但可以通过自定义实现：

```go
package main

import (
    "log"
    "os"
    "io"
)

type LogLevel int

const (
    DEBUG LogLevel = iota
    INFO
    WARN
    ERROR
    FATAL
)

type Logger struct {
    debug *log.Logger
    info  *log.Logger
    warn  *log.Logger
    error *log.Logger
    fatal *log.Logger
    level LogLevel
}

func NewLogger(out io.Writer, level LogLevel) *Logger {
    flags := log.Ldate | log.Ltime | log.Lshortfile
    return &Logger{
        debug: log.New(out, "DEBUG: ", flags),
        info:  log.New(out, "INFO: ", flags),
        warn:  log.New(out, "WARN: ", flags),
        error: log.New(out, "ERROR: ", flags),
        fatal: log.New(out, "FATAL: ", flags),
        level: level,
    }
}

func (l *Logger) Debug(v ...interface{}) {
    if l.level <= DEBUG {
        l.debug.Println(v...)
    }
}

func (l *Logger) Info(v ...interface{}) {
    if l.level <= INFO {
        l.info.Println(v...)
    }
}

// 其他级别方法类似...
```

### 7.2 并发安全的日志记录

log包的所有方法都是线程安全的，可以在并发环境中安全使用：

```go
package main

import (
    "log"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            log.Printf("协程 %d 正在执行", id)
        }(i)
    }
    
    wg.Wait()
    log.Println("所有协程执行完毕")
}
```

## 8. 最佳实践

### 8.1 错误处理

```go
package main

import (
    "log"
    "os"
)

func setupLogger() (*os.File, error) {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return nil, err
    }
    
    log.SetOutput(file)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    return file, nil
}

func main() {
    logFile, err := setupLogger()
    if err != nil {
        log.Fatal("初始化日志失败:", err)
    }
    defer logFile.Close()
    
    // 应用程序逻辑...
    log.Println("应用程序启动成功")
}
```
