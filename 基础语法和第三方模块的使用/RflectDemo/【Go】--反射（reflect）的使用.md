# reflect

Go语言的反射机制通过 `reflect` 包实现，允许程序在运行时检查类型信息、操作变量值、调用方法等

## 核心概念

### 1. reflect.Type 和 reflect.Value

反射的核心是两个基本类型：

- **reflect.Type**：表示Go语言中的类型信息
- **reflect.Value**：表示Go语言中的值信息

### 2. 种类（Kind）

`reflect.Kind` 表示Go语言的基本类型种类，如 `Int`、`String`、`Struct`、`Ptr` 等。

## 基本类型反射操作

### 1. 获取类型信息

```go
func PrintType(t interface{}) {
    fmt.Println(t, "类型是", reflect.TypeOf(t))
    
    // 类型名称 Name()
    fmt.Println(t, "的类型名称是", reflect.TypeOf(t).Name())
    
    // 种类 Kind() -- 底层类型
    fmt.Println(t, "的种类是", reflect.TypeOf(t).Kind())
}
```

**示例输出分析：**
- `int` 类型：类型名称 `int`，种类 `int`
- `string` 类型：类型名称 `string`，种类 `string`
- `bool` 类型：类型名称 `bool`，种类 `bool`
- 自定义类型 `MyInt`：类型名称 `MyInt`，种类 `int`（底层类型）
- 接口类型 `any`：类型名称 `空字符串`，种类 `interface`
- 结构体 `Person`：类型名称 `Person`，种类 `struct`
- 指针类型：类型名称 `空字符串`，种类 `ptr`
- 切片类型：类型名称 `空字符串`，种类 `slice`
- 数组类型：类型名称 `空字符串`，种类 `array`

### 2. 获取值信息

```go
func PrintValue(v interface{}) {
    switch reflect.ValueOf(v).Kind() {
    case reflect.Int:
        fmt.Println("num的值是", reflect.ValueOf(v).Int() + 100)
    default:
        fmt.Println("还无其他类型")
    }
}
```

## 值修改操作

### 1. 通过反射修改值

```go
func Alter(v interface{}) {
    val := reflect.ValueOf(v)
    
    // 检查是否是指针类型
    if val.Kind() == reflect.Ptr {
        // 获取指针指向元素 -- 类似于 *v
        elem := val.Elem()
        
        // 检查指针指向元素是否是int类型
        if elem.Kind() == reflect.Int {
            elem.SetInt(100)
        }
    } else {
        fmt.Println("v不是指针类型")
    }
}
```

### 2. 通过类型断言修改值

```go
func Alter2(v interface{}) {
    if ptr, ok := v.(*int); ok {
        *ptr = 1000
    }
}
```

## 结构体反射操作

### 1. 获取结构体字段信息

```go
type Student struct {
    Name  string `json:"name" form:"name"`
    Age   int    `json:"age" form:"age"`
    Score int    `json:"score" form:"score"`
}

func PrintStuctField(v interface{}) {
    s := reflect.TypeOf(v)
    
    // 获取字段数量
    fmt.Println("结构体字段的数量:", s.NumField())
    
    // 遍历所有字段
    for i := 0; i < s.NumField(); i++ {
        field := s.Field(i)
        fmt.Println("字段名称:", field.Name)
        fmt.Println("字段类型:", field.Type)
        fmt.Println("字段标签:", field.Tag)
        fmt.Println("JSON标签:", field.Tag.Get("json"))
        fmt.Println("Form标签:", field.Tag.Get("form"))
    }
}
```

### 2. 获取和修改结构体字段值

```go
func ChangeStruct(v interface{}) {
    if reflect.TypeOf(v).Kind() != reflect.Ptr {
        fmt.Println("传入的参数不是结构体指针")
        return
    }
    
    if reflect.TypeOf(v).Elem().Kind() != reflect.Struct {
        fmt.Println("传入的参数不是结构体指针")
        return
    }
    
    // 修改结构体字段值
    reflect.ValueOf(v).Elem().FieldByName("Name").SetString("修改后的姓名")
}
```

### 3. 方法反射操作

```go
func PrintStructMethod(v interface{}) {
    s := reflect.TypeOf(v)
    
    // 获取方法数量
    fmt.Println("结构体的方法数量:", s.NumMethod())
    
    // 遍历所有方法
    for i := 0; i < s.NumMethod(); i++ {
        method := s.Method(i)
        fmt.Println("方法名称:", method.Name)
        fmt.Println("方法类型:", method.Type)
    }
}

func CallMethods(v interface{}) {
    val := reflect.ValueOf(v)
    
    // 调用无参方法
    results := val.MethodByName("GetInfo").Call(nil)
    fmt.Println("GetInfo方法返回值:", results[0].Interface())
    
    // 调用有参方法
    params := []reflect.Value{
        reflect.ValueOf("李四"),
        reflect.ValueOf(20),
        reflect.ValueOf(95),
    }
    val.MethodByName("SetInfo").Call(params)
}
```

## 事项

### 1. 类型名称和种类的区别

- **类型名称（Name）**：用户定义的类型名称，如 `MyInt`、`Person`
- **种类（Kind）**：底层基本类型，如 `int`、`struct`、`ptr`

### 2. 指针类型处理

- 修改值必须传递指针类型
- 使用 `Kind() == reflect.Ptr` 检查是否为指针
- 使用 `Elem()` 获取指针指向的元素

### 3. 可导出性规则

- 只有首字母大写的字段和方法才能通过反射访问
- 私有字段和方法无法通过反射获取或操作

### 4. 接收者类型影响

- **值类型实例**：只能看到值接收者方法
- **指针类型实例**：可以看到值接收者和指针接收者方法

## 示例

### 1. 类型判断和值操作

```go
var a int = 10
PrintValue(a)  // 输出: num的值是 110

// 修改值
Alter(&a)      // 通过反射修改
fmt.Println("Alter后:", a)  // 输出: 100

Alter2(&a)     // 通过类型断言修改
fmt.Println("Alter2后:", a) // 输出: 1000
```

### 2. 结构体操作

```go
student := Student{
    Name:  "张三",
    Age:   18,
    Score: 90,
}

PrintStuctField(student)  // 输出字段信息
PrintStructMethod(&student) // 输出方法信息（注意传递指针）
ChangeStruct(&student)    // 修改字段值
```
