# 【Go】--channel

## 1. Channel

### 1.1 什么是Channel

Channel是Go语言中的一种并发原语，可以看作是一个先进先出（FIFO）的队列，用于在goroutine之间安全地传递数据。

### 1.2 Channel的特性

- **类型安全**：每个channel只能传递特定类型的数据
- **线程安全**：channel操作是原子的，无需额外同步
- **阻塞机制**：当channel满时写入阻塞，空时读取阻塞
- **关闭机制**：关闭后无法写入，但可以继续读取剩余数据

## 2. Channel的基本操作

### 2.1 声明和创建

```go
// 声明一个channel
var ch chan int

// 创建channel（无缓冲）
ch = make(chan int)

// 创建带缓冲的channel
ch = make(chan int, 3)  // 容量为3
```

### 2.2 数据操作

```go
// 发送数据到channel
ch <- 100

// 从channel接收数据
value := <-ch

// 关闭channel
close(ch)
```

## 3. 代码示例分析

### 3.1 基础操作

```go
package main

import (
	"fmt"
)

func main() {
	// channel的声明--队列（先进先出）
	var ch chan int
	ch = make(chan int, 3)

	// 给管道传数据
	ch <- 1
	ch <- 2
	ch <- 3

	// 从管道取数据
	a := <-ch
	b := <-ch
	c := <-ch
	fmt.Println("从管道取的数据：", a)
	fmt.Println("从管道取的数据：", b)
	fmt.Println("从管道取的数据：", c)

	// ch的长度和容量
	fmt.Println("ch的长度：", len(ch))
	fmt.Println("ch的容量：", cap(ch))

	// channel无key
	// channel的遍历
	ch1 := make(chan int, 3)
	ch1 <- 11
	ch1 <- 22
	ch1 <- 33
	// 关闭管道 -- 不关闭 for range 会报错deadlock
	close(ch1)

	for v := range ch1 {
		fmt.Println("遍历管道：", v)
	}
}
```

**核心要点：**

1. **缓冲channel操作**：
   
   ```go
   ch := make(chan int, 3)
   ch <- 1
   ch <- 2  
   ch <- 3
   ```
   
2. **channel长度和容量**：
   ```go
   fmt.Println("ch的长度:", len(ch))  // 当前元素数量
   fmt.Println("ch的容量:", cap(ch))  // 总容量
   ```

3. **channel遍历**：
   ```go
   close(ch1)  // 必须关闭才能使用for-range
   for v := range ch1 {
       fmt.Println("遍历管道:", v)
   }
   ```

**关键：**

- 使用`for-range`遍历channel时，必须提前关闭channel
- 遍历会读取所有数据，比`len(ch)`多取一次（检查是否关闭）

### 3.2 goroutine间通信

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func input(ch chan int) {
	defer wg.Done()
	
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("向管道放入的数据:", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch)
}

func output(ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("从管道取的数据:", v)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	var ch = make(chan int, 5)
	wg.Add(1)
	go input(ch)
	wg.Add(1)
	go output(ch)
	wg.Wait()
	fmt.Println("主函数结束")
}
```

**核心要点：**

1. **生产者-消费者模式**：
   ```go
   func input(ch chan int) {
       for i := 0; i < 5; i++ {
           ch <- i
           fmt.Println("向管道放入的数据:", i)
       }
       close(ch)  // 生产完成后关闭
   }
   
   func output(ch chan int) {
       for v := range ch {
           fmt.Println("从管道取的数据:", v)
       }
   }
   ```

2. **WaitGroup同步**：
   ```go
   var wg sync.WaitGroup
   wg.Add(1)
   go input(ch)
   wg.Add(1) 
   go output(ch)
   wg.Wait()
   ```

**关键：**

- 使用`WaitGroup`确保所有goroutine完成后再结束主程序
- `for-range`会自动检测channel关闭，避免死锁

### 3.3 素数计算并发模式

```go
package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// 向inChan中存放数字
func putNum(inChan chan int) {
	defer wg.Done()
	for i := 2; i < 120000; i++ {
		inChan <- i
	}
	close(inChan)
}

// 从inChan中取数字，判断是否为素数，是则放入primeChan
func primeNum(inChan chan int, primeChan chan int, exitChan chan bool) {
	defer wg.Done()
	for num := range inChan {
		isNum := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isNum = false
				break
			}
		}
		if isNum {
			primeChan <- num
		}
	}
	// 利用exitChan通知printPrime方法退出
	exitChan <- true
}

// printPrime 打印素数的方法
func printPrime(primeChan chan int) {
	defer wg.Done()
	for num := range primeChan {
		fmt.Println(num)
	}
}

func main() {
	// 存放数字
	inChan := make(chan int, 1000)
	wg.Add(1)
	go putNum(inChan)

	// 存放素数
	primeChan := make(chan int, 1000)

	// 退出信号--通知primeChan 什么时候退出
	exitChan := make(chan bool, 16)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go primeNum(inChan, primeChan, exitChan)
	}

	wg.Add(1)
	go printPrime(primeChan)

	// 判断exitChan是否存满
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			<-exitChan
		}
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(strings.Repeat("-", 20))
}
```

**核心要点：**

1. **多阶段管道处理**：
   ```go
   // 阶段1：数字生成
   inChan := make(chan int, 1000)
   go putNum(inChan)
   
   // 阶段2：素数判断（多个worker）
   primeChan := make(chan int, 1000)
   for i := 0; i < 10; i++ {
       go primeNum(inChan, primeChan, exitChan)
   }
   
   // 阶段3：结果输出
   go printPrime(primeChan)
   ```

2. **退出信号机制**：
   ```go
   exitChan := make(chan bool, 16)
   
   // 等待所有worker完成
   for i := 0; i < 10; i++ {
       <-exitChan
   }
   close(primeChan)  // 安全关闭
   ```

**关键：**

- 使用专门的exit channel来协调多个worker的退出
- 避免直接关闭还在使用的channel
- 实现了高效的并行素数计算

### 3.4 单向Channel

```go
package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 默认是双向通道
	fmt.Println("默认是双向通道")
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(<-ch)

	fmt.Println(strings.Repeat("-", 20))

	// 只能写
	fmt.Println("只能写")
	ch1 := make(chan<- int, 2)
	ch1 <- 200

	// .\main.go:28:16: invalid operation: 
	// cannot receive from send-only channel ch1 (variable of type chan<- int)
	// fmt.Println(<-ch1) // 编译错误

	fmt.Println(strings.Repeat("-", 20))
	// 只能读
	fmt.Println("只能读")
	ch2 := make(<-chan int, 2)
	// .\main.go:37:2: 
	// invalid operation: cannot send to receive-only channel ch2 (variable of type <-chan int)
	// ch2 <- 300 // 编译错误
	fmt.Println(<-ch2)
}
```

**核心要点：**

1. **单向channel声明**：
   ```go
   var ch chan int        // 双向通道（可读可写）
   var ch1 chan<- int     // 只能写
   var ch2 <-chan int     // 只能读
   ```

2. **编译时类型检查**：
   ```go
   ch1 <- 200             // 正确：可以写入
   // fmt.Println(<-ch1) // 错误：不能从只写channel读取
   
   // ch2 <- 300          // 错误：不能向只读channel写入
   fmt.Println(<-ch2)     // 正确：可以读取
   ```

**关键：**

- 单向channel提供编译时的安全性保证
- 常用于函数参数，限制channel的操作权限

### 3.5 Select多路复用

```go
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		stringChan <- fmt.Sprintf("hello%d", i)
	}

	// select -- IO 多路复用的解决方案
	// select和switch的区别：
	// 1. select会随机执行一个case，而switch会顺序执行
	// 2. select可以监听多个channel，而switch只能监听一个channel

	// select 不需要close channel
	for {
		select {
		case num := <-intChan:
			fmt.Println("intChan:", num)
			time.Sleep(time.Millisecond * 50)
		case str := <-stringChan:
			fmt.Println("stringChan:", str)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Println("数据获取完毕")
			return
		}
	}

	fmt.Println(strings.Repeat("-", 20))
}
```

**核心要点：**

1. **Select语句结构**：
   ```go
   for {
       select {
       case num := <-intChan:
           fmt.Println("intChan:", num)
       case str := <-stringChan:
           fmt.Println("stringChan:", str)
       default:
           fmt.Println("数据获取完毕")
           return
       }
   }
   ```

2. **Select与Switch的区别**：
   - Select随机执行一个可用的case
   - Switch顺序执行case
   - Select可以监听多个channel

**关键：**

- Select是Go语言实现I/O多路复用的核心机制
- 不需要关闭channel即可使用select
- default case用于处理非阻塞操作
