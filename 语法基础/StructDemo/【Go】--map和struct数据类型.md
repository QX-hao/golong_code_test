# Go 语言中的 map 和 struct 数据类型详解

`map用于存储键值对，而struct用于组织不同类型的数据字段。`

- **Map**：适合处理动态的、无序的键值对数据
- **Struct**：适合组织固定的、有结构的数据字段

## 1. Map（映射表）数据类型

### 1.1 基本概念

Map是一种无序的键值对集合，基于哈希表实现。

**特性：**

- 键必须是可比较的类型（string、int等）
- 值可以是任意类型
- 无序存储，遍历顺序不确定

### 1.2 初始化方法

#### 1.2.1 字面量初始化

```go
// 整数键的map
mp1 := map[int]string{
    0: "a",
    1: "b", 
    2: "c",
    3: "d",
    4: "e",
}

// 字符串键的map
mp2 := map[string]int{
    "a": 0,
    "b": 22,
    "c": 33,
}
```

#### 1.2.2 使用make函数初始化

```go
// 指定初始容量
mp1 := make(map[string]int, 8)
mp2 := make(map[string][]int, 10)
```

**重要：**

- Map是引用类型，未初始化的map可以访问但无法存放元素
- 必须为map分配内存后才能使用

### 1.3 访问操作

#### 1.3.1 基本访问

```go
mp := map[string]int{
    "a": 0,
    "b": 1,
    "c": 2,
    "d": 3,
}

fmt.Println(mp["a"]) // 输出: 0
fmt.Println(mp["b"]) // 输出: 1
fmt.Println(mp["d"]) // 输出: 3
fmt.Println(mp["f"]) // 输出: 0 (不存在键返回零值)
```

#### 1.3.2 存在性检查

```go
if val, exist := mp["f"]; exist {
    fmt.Println(val)
} else {
    fmt.Println("key不存在")
}
```

### 1.4 存储操作

#### 1.4.1 基本存储

```go
mp := make(map[string]int, 10)
mp["a"] = 1
mp["b"] = 2
fmt.Println(mp) // 输出: map[a:1 b:2]
```

#### 1.4.2 覆盖操作

```go
mp := make(map[string]int, 10)
mp["a"] = 1
mp["b"] = 2

if _, exist := mp["b"]; exist {
    mp["b"] = 3 // 覆盖原有值
}
fmt.Println(mp) // 输出: map[a:1 b:3]
```

### 1.5 特殊键值处理

#### 1.5.1 NaN键的特殊情况

```go
import "math"

func main() {
    mp := make(map[float64]string, 10)
    mp[math.NaN()] = "a"
    mp[math.NaN()] = "b"
    mp[math.NaN()] = "c"
    
    _, exist := mp[math.NaN()]
    fmt.Println(exist) // 输出: false
    fmt.Println(mp)    // 输出: map[NaN:c NaN:a NaN:b]
}
```

**注意事项：**

- NaN作为键时会出现多个相同键的情况
- 无法正常判断NaN键是否存在
- 无法删除NaN键值对
- 应避免使用NaN作为map的键

### 1.6 删除操作

```go
func main() {
    mp := map[string]int{
        "a": 0,
        "b": 1,
        "c": 2,
        "d": 3,
    }
    
    fmt.Println(mp) // 输出: map[a:0 b:1 c:2 d:3]
    delete(mp, "a")
    fmt.Println(mp) // 输出: map[b:1 c:2 d:3]
}
```

### 1.7 遍历操作

```go
func main() {
    mp := map[string]int{
        "a": 0,
        "b": 1,
        "c": 2,
        "d": 3,
    }
    
    for key, val := range mp {
        fmt.Println(key, val)
    }
    // 可能的输出（无序）:
    // c 2
    // d 3
    // a 0
    // b 1
}
```

### 1.8 清空操作

#### 1.8.1 Go 1.21之前的方法

```go
func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
    }
    
    for k, _ := range m {
        delete(m, k)
    }
    fmt.Println(m) // 输出: map[]
}
```

#### 1.8.2 Go 1.21及之后的方法

```go
func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
    }
    
    clear(m)
    fmt.Println(m) // 输出: map[]
}
```

### 1.9 长度获取

```go
func main() {
    mp := map[string]int{
        "a": 0,
        "b": 1,
        "c": 2,
        "d": 3,
    }
    fmt.Println(len(mp)) // 输出: 4
}
```

## 2. Struct（结构体）数据类型

### 2.1 基本概念

Struct是用于存储一组不同类型数据的复合类型。

**特性：**

- 可以存储不同类型的数据字段
- 通过组合实现类似继承的功能
- 支持方法绑定

### 2.2 结构体声明

#### 2.2.1 基本声明

```go
type Programmer struct {
    Name     string
    Age      int
    Job      string
    Language []string
}

type Person struct {
    name string
    age  int
}
```

#### 2.2.2 相邻字段简化声明

```go
type Rectangle struct {
    height, width, area int
    color               string
}
```

**命名规则：**

- 结构体本身及其字段都遵守大小写命名的暴露方式
- 大写字母开头的字段可被外部包访问
- 小写字母开头的字段只能在包内访问

### 2.3 实例化方法

#### 2.3.1 字段名初始化（推荐）

```go
programmer := Programmer{
    Name:     "jack",
    Age:      19,
    Job:      "coder",
    Language: []string{"Go", "C++"},
}
```

#### 2.3.2 省略字段名初始化（不推荐）

```go
programmer := Programmer{
    "jack",
    19,
    "coder",
    []string{"Go", "C++"},
}
```

### 2.4 构造函数模式

#### 2.4.1 基本构造函数

```go
type Person struct {
    Name    string
    Age     int
    Address string
    Salary  float64
}

func NewPerson(name string, age int, address string, salary float64) *Person {
    return &Person{Name: name, Age: age, Address: address, Salary: salary}
}
```

#### 2.4.2 选项模式（函数式选项模式）

```go
type Person struct {
    Name     string
    Age      int
    Address  string
    Salary   float64
    Birthday string
}

// 选项函数类型
type PersonOptions func(p *Person)

// 具体的选项函数
func WithName(name string) PersonOptions {
    return func(p *Person) {
        p.Name = name
    }
}

func WithAge(age int) PersonOptions {
    return func(p *Person) {
        p.Age = age
    }
}

func WithAddress(address string) PersonOptions {
    return func(p *Person) {
        p.Address = address
    }
}

func WithSalary(salary float64) PersonOptions {
    return func(p *Person) {
        p.Salary = salary
    }
}

// 构造函数
func NewPerson(options ...PersonOptions) *Person {
    p := &Person{}
    for _, option := range options {
        option(p)
    }
    
    // 默认值处理
    if p.Age < 0 {
        p.Age = 0
    }
    
    return p
}

// 使用示例
func main() {
    p1 := NewPerson(
        WithName("John Doe"),
        WithAge(25),
        WithAddress("123 Main St"),
        WithSalary(10000.00),
    )

    p2 := NewPerson(
        WithName("Mike jane"),
        WithAge(30),
    )
}
```

### 2.5 组合（Composition）

#### 2.5.1 显式组合

```go
type Person struct {
    name string
    age  int
}

type Student struct {
    p      Person
    school string
}

type Employee struct {
    p   Person
    job string
}

// 使用示例
student := Student{
    p:      Person{name: "jack", age: 18},
    school: "lili school",
}
fmt.Println(student.p.name) // 输出: jack
```

#### 2.5.2 匿名组合（类似继承）

```go
type Person struct {
    name string
    age  int
}

type Student struct {
    Person  // 匿名组合
    school string
}

type Employee struct {
    Person  // 匿名组合
    job string
}

// 使用示例
student := Student{
    Person: Person{name: "jack", age: 18},
    school: "lili school",
}
fmt.Println(student.name) // 直接访问，输出: jack
```

### 2.6 指针操作

```go
p := &Person{
    name: "jack"
}

// 指针可以直接访问字段，无需解引用
fmt.Println(p.name) // 输出: jack
```

## 3. Map和Struct的对比

### 3.1 使用场景对比

| 特性 | Map | Struct |
|------|-----|--------|
| **数据结构** | 键值对集合 | 字段集合 |
| **键类型** | 必须可比较 | 固定字段名 |
| **顺序** | 无序 | 字段定义顺序固定 |
| **内存布局** | 哈希表，动态 | 连续内存，静态 |
| **性能** | O(1)访问 | O(1)字段访问 |
| **适用场景** | 动态键值数据 | 固定结构数据 |

### 3.2 性能考虑

- **Map**：适合动态键值对，但哈希计算有开销
- **Struct**：内存连续，访问速度快，但结构固定

## 4. 最佳实践

### 4.1 Map使用建议

1. **预分配容量**：使用make时指定合理容量减少扩容
2. **存在性检查**：访问前检查键是否存在
3. **避免NaN键**：避免使用math.NaN()作为键
4. **并发安全**：多goroutine访问时使用sync.Map

### 4.2 Struct使用建议

1. **使用选项模式**：复杂结构体使用函数式选项模式
2. **合理使用组合**：优先使用组合而非模拟继承
3. **字段命名规范**：遵循Go的命名约定
4. **方法绑定**：为结构体绑定相关方法

## 5. 实际应用示例

### 5.1 学生管理系统

```go
type Student struct {
    ID      int
    Name    string
    Grades  map[string]float64
    Contact struct {
        Phone string
        Email string
    }
}

func NewStudent(id int, name string) *Student {
    return &Student{
        ID:     id,
        Name:   name,
        Grades: make(map[string]float64),
    }
}

// 使用示例
student := NewStudent(1, "Alice")
student.Grades["Math"] = 95.5
student.Grades["English"] = 88.0
student.Contact.Phone = "123-456-7890"
student.Contact.Email = "alice@example.com"
```

### 5.2 配置管理系统

```go
type Config struct {
    Database struct {
        Host     string
        Port     int
        Username string
        Password string
    }
    Server struct {
        Port    int
        Timeout time.Duration
    }
    Features map[string]bool
}

func LoadConfig() *Config {
    config := &Config{
        Features: make(map[string]bool),
    }
    
    // 加载配置逻辑...
    return config
}
```

