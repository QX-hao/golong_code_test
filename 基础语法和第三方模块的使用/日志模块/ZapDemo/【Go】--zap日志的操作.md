# 【Go】--zap日志的操作

## 一、快速开始

### 安装
```bash
go get -u go.uber.org/zap
```

### 三种预置模式

#### 1. Development 模式
```go
package main

import "go.uber.org/zap"

func dev() {
    logger, _ := zap.NewDevelopment()
    logger.Info("dev this is info",
        zap.String("name", "dev"),
        zap.Int("age", 18),
        zap.Bool("isMale", true),
    )
    logger.Warn("dev this is warn")
    logger.Error("dev this is error")
}
```
**特点**：
- 文本格式输出
- Warn 和 Error 级别包含堆栈信息
- 适合开发环境

#### 2. Example 模式
```go
func test() {
    logger := zap.NewExample()
    logger.Info("exam this is info")
    logger.Warn("exam this is warn")
    logger.Error("exam this is error")
}
```
**特点**：
- JSON 格式输出
- 只包含 level 和 msg 字段
- 适合示例代码

#### 3. Production 模式
```go
func prod() {
    logger, _ := zap.NewProduction()
    logger.Info("prod this is info",
        zap.String("name", "prod"),
        zap.Int("age", 18),
        zap.Bool("isMale", true),
    )
    logger.Warn("prod this is warn")
    logger.Error("prod this is error")
}
```
**特点**：
- JSON 格式输出
- 包含时间戳和函数位置
- 适合生产环境

## 二、基本使用

### 自定义配置
```go
package main

import (
    "fmt"
    "strings"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func main() {
    cfg := zap.NewDevelopmentConfig()
    // 设置日志级别
    cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
    
    // 自定义时间格式
    cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
    
    logger, _ := cfg.Build()
    logger.Debug("这是一条 Debug 日志")  // 会被过滤掉
    logger.Info("这是一条 Info 日志")
    logger.Warn("这是一条 Warn 日志")
}
```

### Sugar Logger
```go
package main

import "go.uber.org/zap"

func main() {
    logger, _ := zap.NewDevelopment()
    
    // 使用 SugarLogger 进行格式化输出
    sugar := logger.Sugar()
    sugar.Infof("这是一条 %s 日志", "Info")
    sugar.Debugf("这是一条 %s 日志", "Debug")
    sugar.Warnf("这是一条 %s 日志", "Warn")
    sugar.Errorf("这是一条 %s 日志", "Error")
}
```

## 三、配置详解

### 日志级别
```go
// 设置日志级别
cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)  // 所有级别
cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)   // Info 及以上
cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)   // Warn 及以上
cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)  // Error 及以上
```

### 自定义编码器

#### 1. 彩色日志级别
```go
package main

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

const (
    colorRed    = "\033[31m"
    colorGreen  = "\033[32m"
    colorYellow = "\033[33m"
    colorBlue   = "\033[34m"
    colorReset  = "\033[0m"
)

func EncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
    switch level {
    case zapcore.DebugLevel:
        enc.AppendString(colorBlue + "DEBUG" + colorReset)
    case zapcore.InfoLevel:
        enc.AppendString(colorGreen + "INFO" + colorReset)
    case zapcore.WarnLevel:
        enc.AppendString(colorYellow + "WARN" + colorReset)
    case zapcore.ErrorLevel:
        enc.AppendString(colorRed + "ERROR" + colorReset)
    default:
        enc.AppendString(level.String())
    }
}

func main() {
    cfg := zap.NewDevelopmentConfig()
    cfg.EncoderConfig.EncodeLevel = EncodeLevel
    logger, _ := cfg.Build()
    logger.Debug("debug message")
    logger.Info("info message")
    logger.Warn("warn message")
    logger.Error("error message")
}
```

#### 2. 自定义前缀
```go
package main

import (
    "os"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// 定义前缀
const logPrefix = "[MyApp] "

// 自定义 Encoder
type prefixedEncoder struct {
    zapcore.Encoder
}

func (e *prefixedEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) ([]byte, error) {
    // 先调用原始的 EncodeEntry 方法生成日志行
    buf, err := e.Encoder.EncodeEntry(entry, fields)
    if err != nil {
        return nil, err
    }

    // 在日志行的最前面添加前缀
    logLine := buf.String()
    buf.Reset()
    buf.AppendString(logPrefix + logLine)

    return buf.Bytes(), nil
}

func main() {
    cfg := zap.NewDevelopmentConfig()
    cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
    
    encoder := &prefixedEncoder{
        Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig),
    }
    
    core := zapcore.NewCore(
        encoder,
        zapcore.AddSync(os.Stdout),
        zapcore.DebugLevel,
    )
    
    logger := zap.New(core, zap.AddCaller())
    logger.Info("this is info")
    logger.Warn("this is warn")
    logger.Error("this is error")
}
```

## 四、高级功能

### 多输出目标

#### 方法一：使用 Tee
```go
package main

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

func funOne() {
    cfg := zap.NewProductionConfig()
    cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

    // 文件日志
    file, _ := os.OpenFile("app1.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    filecore := zapcore.NewCore(
        zapcore.NewJSONEncoder(cfg.EncoderConfig),
        file,
        cfg.Level,
    )

    // 控制台日志
    consolecore := zapcore.NewCore(
        zapcore.NewJSONEncoder(cfg.EncoderConfig),
        zapcore.AddSync(os.Stdout),
        cfg.Level,
    )
    
    core := zapcore.NewTee(filecore, consolecore)
    logger := zap.New(core, zap.AddCaller())
    logger.Info("info1日志")
}
```

#### 方法二：使用 MultiWriteSyncer
```go
func funTwo() {
    cfg := zap.NewProductionConfig()
    cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

    // 文件日志
    file, _ := os.OpenFile("app2.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    multiWriteSyncer := zapcore.NewMultiWriteSyncer(
        zapcore.AddSync(os.Stdout),
        zapcore.AddSync(file),
    )

    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(cfg.EncoderConfig),
        multiWriteSyncer,
        cfg.Level,
    )
    logger := zap.New(core, zap.AddCaller())
    logger.Info("info2日志")
}
```

### 日志轮转和分片

#### 按时间分片
```go
package main

import (
    "os"
    "sync"
    "time"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// 自定义日志写入器
type dynamicLogWriter struct {
    mu         sync.Mutex
    currentDay string
    file       *os.File
    logDir     string
}

func (w *dynamicLogWriter) Write(p []byte) (n int, err error) {
    w.mu.Lock()
    defer w.mu.Unlock()

    // 检查是否需要切换到新的日志文件
    currentDay := time.Now().Format("2006-01-02")
    if currentDay != w.currentDay {
        // 关闭当前日志文件
        if w.file != nil {
            w.file.Close()
        }

        // 创建新的日志文件
        if err := os.MkdirAll(w.logDir, 0755); err != nil {
            return 0, err
        }
        filePath := w.logDir + "/app-" + currentDay + ".log"
        file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
            return 0, err
        }
        w.file = file
        w.currentDay = currentDay
    }

    // 写入日志
    return w.file.Write(p)
}

func initLogger() {
    cfg := zap.NewDevelopmentConfig()
    cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
    
    writer := &dynamicLogWriter{
        logDir: "logs",
    }
    
    consoleCore := zapcore.NewCore(
        zapcore.NewConsoleEncoder(cfg.EncoderConfig),
        zapcore.AddSync(os.Stdout),
        zapcore.DebugLevel,
    )
    
    fileCore := zapcore.NewCore(
        zapcore.NewConsoleEncoder(cfg.EncoderConfig),
        zapcore.AddSync(writer),
        zapcore.WarnLevel,
    )
    
    core := zapcore.NewTee(consoleCore, fileCore)
    logger := zap.New(core, zap.AddCaller())
    zap.ReplaceGlobals(logger)
}

func main() {
    initLogger()
    zap.L().Info("info1日志")
    zap.L().Warn("warn1日志")
}
```

#### 按级别分片
```go
package main

import (
    "fmt"
    "go.uber.org/zap"
    "go.uber.org/zap/buffer"
    "go.uber.org/zap/zapcore"
    "os"
    "time"
)

// 时间分片和级别分片同时做
type logEncoder struct {
    zapcore.Encoder
    errFile     *os.File
    file        *os.File
    currentDate string
}

func (e *logEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
    buff, err := e.Encoder.EncodeEntry(entry, fields)
    if err != nil {
        return nil, err
    }
    
    data := buff.String()
    buff.Reset()
    buff.AppendString("[myApp] " + data)
    data = buff.String()
    
    // 时间分片
    now := time.Now().Format("2006-01-02")
    if e.currentDate != now {
        os.MkdirAll(fmt.Sprintf("logs/%s", now), 0666)
        name := fmt.Sprintf("logs/%s/out.log", now)
        file, _ := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
        e.file = file
        e.currentDate = now
    }

    // 级别分片
    switch entry.Level {
    case zapcore.ErrorLevel:
        if e.errFile == nil {
            name := fmt.Sprintf("logs/%s/err.log", now)
            file, _ := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
            e.errFile = file
        }
        e.errFile.WriteString(buff.String())
    }

    if e.currentDate == now {
        e.file.WriteString(data)
    }
    return buff, nil
}

func InitLog() *zap.Logger {
    cfg := zap.NewDevelopmentConfig()
    cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
    
    encoder := &logEncoder{
        Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig),
    }
    
    core := zapcore.NewCore(
        encoder,
        zapcore.AddSync(os.Stdout),
        zapcore.InfoLevel,
    )
    
    logger := zap.New(core, zap.AddCaller())
    zap.ReplaceGlobals(logger)
    return logger
}

func main() {
    logger := InitLog()
    logger.Info("this is info")
    logger.Warn("this is warn")
    logger.Error("this is error1")
    logger.Error("this is error2")
}
```

## 五、实际应用

### 全局日志实例
```go
// 初始化全局日志
func initLogger() {
    logger, _ := zap.NewProduction()
    zap.ReplaceGlobals(logger)
}

// 在其他地方使用
func someFunction() {
    zap.L().Info("函数开始执行")
    defer zap.L().Info("函数执行结束")
    
    // 使用 SugarLogger
    zap.S().Infof("用户 %s 登录成功", "username")
}
```

### 结构化日志字段
```go
func processUser(userID int, username string) {
    logger := zap.L().With(
        zap.Int("userID", userID),
        zap.String("username", username),
    )
    
    logger.Info("开始处理用户")
    // ... 业务逻辑
    logger.Info("用户处理完成")
}
```

### 错误处理
```go
func riskyOperation() error {
    if err := doSomething(); err != nil {
        zap.L().Error("操作失败", 
            zap.String("operation", "doSomething"),
            zap.Error(err),
            zap.Stack("stacktrace"),
        )
        return err
    }
    return nil
}
```