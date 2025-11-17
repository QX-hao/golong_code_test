# Go 语言 strings 包常用方法详解

| 函数名            | 功能描述                     |
| ----------------- | ---------------------------- |
| `Contains`        | 检查字符串是否包含子串       |
| `ContainsAny`     | 检查是否包含任何指定字符     |
| `ContainsRune`    | 检查是否包含特定Unicode字符  |
| `HasPrefix`       | 检查前缀                     |
| `HasSuffix`       | 检查后缀                     |
| `Index`           | 查找子串第一次出现位置       |
| `LastIndex`       | 查找子串最后一次出现位置     |
| `IndexAny`        | 查找任何指定字符第一次出现   |
| `IndexRune`       | 查找Unicode字符位置          |
| `LastIndexAny`    | 查找任何指定字符最后一次出现 |
| `Count`           | 统计子串出现次数             |
| `Compare`         | 按字典序比较字符串           |
| `EqualFold`       | 不区分大小写比较             |
| `ToUpper`         | 转换为大写                   |
| `ToLower`         | 转换为小写                   |
| `ToTitle`         | 转换为标题格式               |
| `Trim`            | 去除指定字符                 |
| `TrimSpace`       | 去除前后空白字符             |
| `TrimLeft`        | 去除左侧指定字符             |
| `TrimRight`       | 去除右侧指定字符             |
| `TrimPrefix`      | 去除前缀                     |
| `TrimSuffix`      | 去除后缀                     |
| `Replace`         | 替换字符串                   |
| `ReplaceAll`      | 替换所有匹配项               |
| `Map`             | 使用函数进行字符替换         |
| `Split`           | 分割字符串                   |
| `SplitN`          | 限制分割次数                 |
| `SplitAfter`      | 分割但保留分隔符             |
| `Fields`          | 按空白字符分割               |
| `FieldsFunc`      | 使用自定义函数分割           |
| `Join`            | 连接字符串切片               |
| `Repeat`          | 重复字符串                   |
| `strings.Builder` | 高效字符串构建               |

## 1. 字符串判断与查询

### Contains - 检查包含关系
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Contains 检查字符串是否包含子串
    fmt.Println(strings.Contains("hello world", "hello")) // true
    fmt.Println(strings.Contains("hello world", "foo"))   // false
    
    // ContainsAny 检查是否包含任何指定字符
    fmt.Println(strings.ContainsAny("hello", "aeiou")) // true
    fmt.Println(strings.ContainsAny("hello", "xyz"))   // false
    
    // ContainsRune 检查是否包含特定 Unicode 字符
    fmt.Println(strings.ContainsRune("hello", 'h')) // true
    fmt.Println(strings.ContainsRune("hello", 'z')) // false
}
```

### 前缀/后缀判断
```go
func prefixSuffixDemo() {
    s := "hello world"
    
    // HasPrefix 检查前缀
    fmt.Println(strings.HasPrefix(s, "hello")) // true
    fmt.Println(strings.HasPrefix(s, "world")) // false
    
    // HasSuffix 检查后缀
    fmt.Println(strings.HasSuffix(s, "world")) // true
    fmt.Println(strings.HasSuffix(s, "hello")) // false
}
```

### 字符位置查找
```go
func searchDemo() {
    s := "hello world, hello go"
    
    // Index 查找子串第一次出现位置
    fmt.Println(strings.Index(s, "world"))     // 6
    fmt.Println(strings.Index(s, "python"))    // -1
    
    // LastIndex 查找子串最后一次出现位置
    fmt.Println(strings.LastIndex(s, "hello")) // 13
    
    // IndexAny 查找任何指定字符第一次出现位置
    fmt.Println(strings.IndexAny(s, "aeiou"))  // 1 (e的位置)
    
    // IndexRune 查找 Unicode 字符位置
    fmt.Println(strings.IndexRune(s, '世'))    // -1
    
    // LastIndexAny 查找任何指定字符最后一次出现位置
    fmt.Println(strings.LastIndexAny(s, "aeiou")) // 16 (o的位置)
}
```

## 2. 字符串计数与比较

### Count - 统计出现次数
```go
func countDemo() {
    s := "banana"
    
    // Count 统计子串出现次数
    fmt.Println(strings.Count(s, "a"))     // 3
    fmt.Println(strings.Count(s, "an"))    // 2
    fmt.Println(strings.Count(s, "ana"))   // 1
    fmt.Println(strings.Count("five", "")) // 5 (注意空串的情况)
}
```

### Compare - 字符串比较
```go
func compareDemo() {
    // Compare 按字典序比较两个字符串
    // 返回: 
    //   -1 如果 a < b
    //    0 如果 a == b
    //    1 如果 a > b
    
    fmt.Println(strings.Compare("apple", "banana")) // -1
    fmt.Println(strings.Compare("apple", "apple"))  // 0
    fmt.Println(strings.Compare("banana", "apple")) // 1
    
    // 通常更推荐直接使用比较运算符
    fmt.Println("apple" < "banana") // true
    fmt.Println("apple" == "apple") // true
}
```

### EqualFold - 不区分大小写比较
```go
func equalFoldDemo() {
    // EqualFold 不区分大小写比较
    fmt.Println(strings.EqualFold("Go", "go"))     // true
    fmt.Println(strings.EqualFold("Hello", "HELLO")) // true
    fmt.Println(strings.EqualFold("Go", "Python")) // false
}
```

## 3. 字符串变换

### 大小写转换
```go
func caseDemo() {
    s := "Hello World"
    
    // ToUpper 转换为大写
    fmt.Println(strings.ToUpper(s)) // "HELLO WORLD"
    
    // ToLower 转换为小写
    fmt.Println(strings.ToLower(s)) // "hello world"
    
    // Title 已废弃，使用 ToTitle
    fmt.Println(strings.ToTitle(s)) // "HELLO WORLD"
    
    // 特殊的大小写处理
    fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "ı")) // "I"
}
```

### 修剪操作
```go
func trimDemo() {
    s := "   hello world   "
    
    // TrimSpace 去除前后空白字符
    fmt.Printf("'%s'\n", strings.TrimSpace(s)) // 'hello world'
    
    // Trim 去除指定字符
    fmt.Println(strings.Trim("!!!hello!!!", "!")) // "hello"
    
    // TrimLeft 去除左侧指定字符
    fmt.Println(strings.TrimLeft("!!!hello!!!", "!")) // "hello!!!"
    
    // TrimRight 去除右侧指定字符
    fmt.Println(strings.TrimRight("!!!hello!!!", "!")) // "!!!hello"
    
    // TrimPrefix 去除前缀
    fmt.Println(strings.TrimPrefix("hello world", "hello ")) // "world"
    
    // TrimSuffix 去除后缀
    fmt.Println(strings.TrimSuffix("hello world", " world")) // "hello"
    
    // 使用 TrimFunc 自定义修剪规则
    result := strings.TrimFunc("123hello456", func(r rune) bool {
        return r >= '0' && r <= '9'
    })
    fmt.Println(result) // "hello"
}
```

### 替换操作
```go
func replaceDemo() {
    s := "hello world, hello go"
    
    // Replace 替换字符串
    // 参数：原字符串，旧子串，新子串，替换次数（-1表示全部替换）
    fmt.Println(strings.Replace(s, "hello", "hi", 1))  // "hi world, hello go"
    fmt.Println(strings.Replace(s, "hello", "hi", -1)) // "hi world, hi go"
    
    // ReplaceAll 替换所有（Go 1.12+）
    fmt.Println(strings.ReplaceAll(s, "hello", "hi")) // "hi world, hi go"
    
    // Map 使用函数进行字符替换
    result := strings.Map(func(r rune) rune {
        if r == 'o' {
            return '0' // 将 o 替换为 0
        }
        return r
    }, "hello world")
    fmt.Println(result) // "hell0 w0rld"
}
```

## 4. 字符串分割与连接

### 分割操作
```go
func splitDemo() {
    s := "apple,banana,orange"
    
    // Split 分割字符串
    parts := strings.Split(s, ",")
    fmt.Printf("%#v\n", parts) // []string{"apple", "banana", "orange"}
    
    // SplitN 限制分割次数
    partsN := strings.SplitN(s, ",", 2)
    fmt.Printf("%#v\n", partsN) // []string{"apple", "banana,orange"}
    
    // SplitAfter 分割但保留分隔符
    partsAfter := strings.SplitAfter(s, ",")
    fmt.Printf("%#v\n", partsAfter) // []string{"apple,", "banana,", "orange"}
    
    // Fields 按空白字符分割
    fields := strings.Fields("  hello   world\tgo\nlang  ")
    fmt.Printf("%#v\n", fields) // []string{"hello", "world", "go", "lang"}
    
    // FieldsFunc 使用自定义函数分割
    fieldsFunc := strings.FieldsFunc("apple,banana;orange", func(r rune) bool {
        return r == ',' || r == ';'
    })
    fmt.Printf("%#v\n", fieldsFunc) // []string{"apple", "banana", "orange"}
}
```

### 连接操作
```go
func joinDemo() {
    // Join 连接字符串切片
    fruits := []string{"apple", "banana", "orange"}
    result := strings.Join(fruits, ", ")
    fmt.Println(result) // "apple, banana, orange"
    
    // 复杂连接示例
    words := []string{"hello", "world", "go"}
    fmt.Println(strings.Join(words, "::")) // "hello::world::go"
}
```

## 5. 重复与填充

### Repeat - 重复字符串
```go
func repeatDemo() {
    // Repeat 重复字符串
    fmt.Println(strings.Repeat("go", 3))    // "gogogo"
    fmt.Println(strings.Repeat("-", 10))    // "----------"
    fmt.Println(strings.Repeat("hello ", 2)) // "hello hello "
}
```

## 6. 字符串遍历与处理

### 字符遍历
```go
func rangeDemo() {
    s := "hello世界"
    
    // 使用 range 遍历字符串（正确处理 Unicode）
    for i, r := range s {
        fmt.Printf("位置 %d: 字符 %c (Unicode: %U)\n", i, r, r)
    }
    
    // 统计字符数（不是字节数）
    fmt.Printf("字符串 '%s' 的长度:\n", s)
    fmt.Printf("字节数: %d\n", len(s))
    fmt.Printf("字符数: %d\n", utf8.RuneCountInString(s))
}
```

## 7. 实际应用示例

### 综合应用：字符串处理工具
```go
package main

import (
    "fmt"
    "strings"
    "unicode"
)

// 字符串处理工具函数
func stringUtils() {
    // 1. 清理用户输入
    userInput := "  Hello, World!  "
    cleaned := strings.TrimSpace(userInput)
    fmt.Printf("清理后: '%s'\n", cleaned)
    
    // 2. 检查文件扩展名
    filename := "document.pdf"
    if strings.HasSuffix(strings.ToLower(filename), ".pdf") {
        fmt.Println("这是一个PDF文件")
    }
    
    // 3. 分割CSV数据
    csvData := "John,25,Engineer\nJane,30,Designer"
    lines := strings.Split(csvData, "\n")
    for i, line := range lines {
        fields := strings.Split(line, ",")
        fmt.Printf("第%d行: %v\n", i+1, fields)
    }
    
    // 4. 构建URL路径
    baseURL := "https://api.example.com"
    endpoints := []string{"users", "123", "profile"}
    fullURL := baseURL + "/" + strings.Join(endpoints, "/")
    fmt.Printf("完整URL: %s\n", fullURL)
    
    // 5. 密码强度检查
    password := "Hello123"
    hasUpper := strings.IndexFunc(password, unicode.IsUpper) >= 0
    hasLower := strings.IndexFunc(password, unicode.IsLower) >= 0
    hasDigit := strings.IndexFunc(password, unicode.IsDigit) >= 0
    
    if hasUpper && hasLower && hasDigit {
        fmt.Println("密码强度足够")
    } else {
        fmt.Println("密码需要包含大小写字母和数字")
    }
}

func main() {
    stringUtils()
}
```

### 性能优化：使用 Builder
```go
func builderDemo() {
    // 对于大量字符串拼接，使用 Builder 更高效
    var builder strings.Builder
    
    names := []string{"Alice", "Bob", "Charlie"}
    for i, name := range names {
        if i > 0 {
            builder.WriteString(", ")
        }
        builder.WriteString(name)
    }
    
    result := builder.String()
    fmt.Println(result) // "Alice, Bob, Charlie"
    
    // 重置 Builder 重用
    builder.Reset()
    builder.WriteString("重新开始")
    fmt.Println(builder.String()) // "重新开始"
}
```

## 8. 重要注意事项

1. **字符串不可变**：所有操作都返回新字符串
2. **Unicode 安全**：大多数函数正确处理 Unicode 字符
3. **性能考虑**：大量字符串操作时考虑使用 `strings.Builder`
4. **空字符串处理**：注意空字符串和 nil 的区别

```go
func importantNotes() {
    // 空字符串 vs nil
    var s1 string        // 零值是空字符串 ""
    var s2 *string       // 指针，零值是 nil
    
    fmt.Printf("s1: '%s', is empty: %v\n", s1, s1 == "")
    fmt.Printf("s2: %v\n", s2)
    
    // 正确处理可能为空的情况
    processString("")
    processString("hello")
}

func processString(s string) {
    if strings.TrimSpace(s) == "" {
        fmt.Println("输入字符串为空或仅包含空白字符")
        return
    }
    fmt.Printf("处理字符串: %s\n", s)
}
```
