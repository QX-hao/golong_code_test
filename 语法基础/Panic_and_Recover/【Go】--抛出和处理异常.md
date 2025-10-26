# Go语言异常处理：panic和recover机制

在Go语言中，异常处理主要通过`panic`和`recover`两个内置函数来实现。与传统的try-catch机制不同，Go采用了一种更加简洁和明确的错误处理方式。

## panic函数

### 功能说明
- `panic`函数用于引发运行时异常
- 当程序遇到无法继续执行的严重错误时，可以主动调用`panic`
- `panic`会导致当前goroutine的执行被中断

### 使用场景
1. 程序遇到不可恢复的错误
2. 参数验证失败
3. 资源获取失败
4. 业务逻辑中的严重错误

## recover函数

### 功能说明
- `recover`函数用于捕获`panic`引发的异常
- `recover`只能在`defer`函数中调用才有效
- 成功捕获异常后，程序可以继续执行

### 使用限制
- 必须在`defer`函数内部调用
- 只能捕获当前goroutine的panic
- 如果没有panic发生，`recover`会返回`nil`

## 示例代码分析

```go
package main

import (
	"fmt"
)

func fn1() {
	fmt.Println("start")
}

func fn2() {
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	panic("我发生了错误")
}

func main() {
	fn1()
	fn2()
	fmt.Println("ending")
}
```

**执行结果：**
```
start
recover: 我发生了错误
ending
```

**关键点：**
- `panic`在`fn2`函数中被触发
- `recover`在`defer`中成功捕获异常
- 程序继续执行，输出"ending"



```go
package main

import (
	"fmt"
)

func fn1(a,b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("计算错误:", err)
		}
	}()
	return a/b
}

func main()  {
	fmt.Println(fn1(10,0))
	fmt.Println("ending")
}
```

**执行结果：**
```
计算错误: runtime error: integer divide by zero
0
ending
```

**关键点：**
- 除零操作会触发运行时panic
- `recover`捕获异常并输出错误信息
- 函数返回默认值0，程序继续执行



```go
package main

import (
	"errors"
	"fmt"
)

func readFile(fileName string) error {
	if fileName == "test.txt" {
		return nil
	} else {
		return errors.New("文件不存在")
	}
}

func fn() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("读取文件错误:", err)
		}
	}()

	err := readFile("xxx.txt")
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("start")
	fn()
	fmt.Println("ending")
}
```

**执行结果：**
```
start
读取文件错误: 文件不存在
ending
```

**关键点：**
- 自定义错误通过`panic`抛出
- `recover`捕获业务逻辑中的异常
- 优雅地处理文件不存在的情况

## 工作原理

### panic的执行流程
1. 程序遇到`panic`或运行时错误
2. 当前函数的执行立即停止
3. 开始执行所有的`defer`语句（LIFO顺序）
4. 如果`defer`中有`recover`，则捕获异常
5. 如果没有`recover`，程序崩溃并输出堆栈信息

### recover的捕获机制
1. `recover`必须放在`defer`函数中
2. 当panic发生时，会逆序执行所有已注册的defer函数
3. 在defer函数中调用`recover`可以捕获panic值
4. 捕获后，程序从panic点之后的代码继续执行
