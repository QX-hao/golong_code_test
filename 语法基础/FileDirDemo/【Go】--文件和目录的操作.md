# Go    文件和目录操作

## 一、目录操作

### 1.1 创建目录

#### 1.1.1 创建单级目录
使用`os.Mkdir`函数创建单级目录：

```go
package main

import (
    "os"
    "fmt"
)

// 创建单级目录
func createSingleDir(dirname string, perm os.FileMode) {
    err := os.Mkdir(dirname, perm)
    if err != nil {
        fmt.Println("创建目录失败", err)
        return
    }
    fmt.Println("目录创建成功")
}

func main() {
    // 创建目录，权限为0666
    createSingleDir("testdir", 0666)
}
```

#### 1.1.2 创建多级目录
使用`os.MkdirAll`函数创建多级目录：

```go
package main

import (
    "os"
    "fmt"
)

// 创建多级目录
func createMultiLevelDir(dirname string, perm os.FileMode) {
    err := os.MkdirAll(dirname, perm)
    if err != nil {
        fmt.Println("创建目录失败", err)
        return
    }
    fmt.Println("多级目录创建成功")
}

func main() {
    // 创建多级目录
    createMultiLevelDir("parent/child/grandchild", 0666)
}
```

### 1.2 删除目录

#### 1.2.1 删除空目录
使用`os.Remove`函数删除空目录：

```go
package main

import (
    "os"
    "fmt"
)

// 删除空目录
func removeEmptyDir(dir string) {
    err := os.Remove(dir)
    if err != nil {
        fmt.Println("删除目录失败", err)
        return
    }
    fmt.Println("目录删除成功")
}

func main() {
    removeEmptyDir("testdir")
}
```

#### 1.2.2 删除非空目录
使用`os.RemoveAll`函数递归删除目录及其所有内容：

```go
package main

import (
    "os"
    "fmt"
)

// 删除非空目录
func removeNonEmptyDir(dir string) {
    err := os.RemoveAll(dir)
    if err != nil {
        fmt.Println("删除目录失败", err)
        return
    }
    fmt.Println("非空目录删除成功")
}

func main() {
    removeNonEmptyDir("parent")
}
```

## 二、文件操作

### 2.1 文件读取

#### 2.1.1 方法一：使用 os.Open 和 file.Read
适用于需要精细控制读取过程的场景：

```go
package main

import (
    "fmt"
    "io"
    "os"
)

// 打开文件（只读模式）
func openFile(filename string) (*os.File, error) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("打开文件失败", err)
        return nil, err
    }
    return file, nil
}

// 读取文件内容（缓冲区方式）
func readFileWithBuffer(file *os.File) {
    var buf []byte
    var tempBuff = make([]byte, 128)

    for {
        n, err := file.Read(tempBuff)
        if err != nil {
            if err == io.EOF {
                fmt.Println("文件读取结束")
                fmt.Printf("文件内容：%s\n", string(buf))
                break
            } else {
                fmt.Println("读取文件失败", err)
                return
            }
        } else {
            fmt.Printf("读取了%d字节\n", n)
            buf = append(buf, tempBuff[:n]...)
        }
    }
}

func main() {
    file, err := openFile("config.json")
    if err != nil {
        return
    }
    defer file.Close()
    
    readFileWithBuffer(file)
}
```

#### 2.1.2 方法二：使用 bufio 读取器
适用于需要逐行读取的场景：

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func readFileWithBufio(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("打开文件失败", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Println("读取文件失败", err)
    }
}

func main() {
    readFileWithBufio("config.json")
}
```

#### 2.1.3 方法三：使用 os.ReadFile
适用于读取小文件的简单场景：

```go
package main

import (
    "fmt"
    "os"
)

func readEntireFile(filename string) {
    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("读取文件失败", err)
        return
    }
    fmt.Printf("文件内容：%s\n", string(data))
}

func main() {
    readEntireFile("config.json")
}
```

### 2.2 文件写入

#### 2.2.1 方法一：直接写入文件
使用`os.OpenFile`和文件对象的写入方法：

```go
package main

import (
    "os"
    "time"
    "fmt"
)

func writeFileDirectly(filename string, flag int, perm os.FileMode) *os.File {
    file, err := os.OpenFile(filename, flag, perm)
    if err != nil {
        fmt.Println("打开文件失败", err)
        return nil
    }
    
    // 写入字节切片数据
    file.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + " 直接写入字节数据\n"))
    
    // 直接写入字符串数据
    file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " 直接写入字符串数据\n")
    
    return file
}

func main() {
    file := writeFileDirectly("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if file != nil {
        defer file.Close()
    }
}
```

#### 2.2.2 方法二：使用 bufio 写入器
适用于需要缓冲写入的场景：

```go
package main

import (
    "bufio"
    "os"
    "time"
    "fmt"
)

func writeFileWithBufio(filename string, flag int, perm os.FileMode) *os.File {
    file, err := os.OpenFile(filename, flag, perm)
    if err != nil {
        fmt.Println("打开文件失败", err)
        return nil
    }
    
    writer := bufio.NewWriter(file)
    
    // 写入缓存
    writer.WriteString(time.Now().Format("2006-01-02 15:04:05") + " bufio写入数据\n")
    
    // 将缓存数据写入文件
    writer.Flush()
    
    return file
}

func main() {
    file := writeFileWithBufio("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if file != nil {
        defer file.Close()
    }
}
```

#### 2.2.3 方法三：使用 os.WriteFile
适用于一次性写入小文件的简单场景：

```go
package main

import (
    "os"
    "fmt"
)

func writeFileSimple(filename string, perm os.FileMode) {
    data := []byte("Hello, Golang! 这是一次性写入的文件内容。")
    err := os.WriteFile(filename, data, perm)
    if err != nil {
        fmt.Println("写入文件失败", err)
        return
    }
    fmt.Println("文件写入成功")
}

func main() {
    writeFileSimple("output.txt", 0666)
}
```

### 2.3 文件复制

#### 2.3.1 方法一：使用 os.ReadFile 和 os.WriteFile
适用于小文件的快速复制：

```go
package main

import (
    "os"
    "fmt"
)

func copyFileSimple(source, destination string) {
    data, err := os.ReadFile(source)
    if err != nil {
        fmt.Println("读取源文件失败:", err)
        return
    }
    
    err = os.WriteFile(destination, data, 0666)
    if err != nil {
        fmt.Println("写入目标文件失败:", err)
        return
    }
    
    fmt.Println("文件复制成功")
}

func main() {
    copyFileSimple("source.txt", "destination.txt")
}
```

#### 2.3.2 方法二：使用流式处理
适用于大文件的复制，节省内存：

```go
package main

import (
    "os"
    "fmt"
    "io"
)

// 利用os.Open、os.Read(流处理)和os.OpenFile方法复制文件
func CopyFile2(filename string, newFilename string) {
    var buf_all []byte
    
    // 打开源文件
    if file, err := os.Open(filename); err != nil {
        fmt.Println("打开原文件失败:", err)
        return
    } else {
        defer file.Close()
        
        var buf = make([]byte, 128)
        for {
            if n, readerr := file.Read(buf); readerr != nil {
                if readerr == io.EOF {
                    fmt.Println("读取完成")
                    break
                } else {
                    fmt.Println("读取文件失败:", readerr)
                    return
                }
            } else {
                buf_all = append(buf_all, buf[:n]...)
            }
        }

        // 使用os.OpenFile创建目标文件
        if newFile, newerr := os.OpenFile(newFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); newerr != nil {
            fmt.Println("创建备份文件失败", newerr)
            return
        } else {
            defer newFile.Close()
            newFile.Write(buf_all)
        }
    }
}

func main() {
    CopyFile2("largefile.txt", "largefile_copy.txt")
}
```

## 三、文件打开模式详解

### 3.1 常用打开模式标志

`os.OpenFile`函数支持多种打开模式，可以组合使用：

```go
// 基本模式（三选一）
os.O_RDONLY  // 只读模式
os.O_WRONLY  // 只写模式  
os.O_RDWR    // 读写模式

// 附加模式（可组合）
os.O_CREATE   // 文件不存在时创建
os.O_APPEND   // 追加模式
os.O_TRUNC    // 打开时清空文件内容
os.O_EXCL     // 与O_CREATE一起使用，文件必须不存在

// 使用示例
file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
```

### 3.2 文件权限说明

文件权限使用八进制表示：

```go
0666  // 所有用户可读写
0644  // 所有者可读写，其他用户只读
0600  // 仅所有者可读写
```

## 四、错误处理最佳实践

### 4.1 使用 defer 确保资源释放

```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // 确保文件被关闭
    
    // 处理文件内容
    // ...
    
    return nil
}
```

### 4.2 检查文件是否存在

```go
func checkFileExists(filename string) bool {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return err == nil
}
```
