# 【Go】--互斥锁和读写锁

## 作用

互斥锁（Mutex）和读写锁（RWMutex）是Go语言中用于保护共享资源、防止数据竞争的重要同步机制。

## 1. 数据竞争问题

### 1.1 数据竞争

数据竞争（Data Race）发生在多个goroutine并发访问同一共享变量，且至少有一个访问是写入操作时。如果没有适当的同步机制，会导致不可预测的结果。

### 1.2 数据竞争的危害

- **结果不确定性**：相同的代码可能产生不同的结果
- **内存损坏**：可能导致程序崩溃或数据损坏
- **难以调试**：问题可能只在特定条件下出现

## 2. 互斥锁（Mutex）

### 2.1 互斥锁

互斥锁（Mutual Exclusion Lock）是最基本的同步原语，保证同一时间只有一个goroutine可以访问受保护的临界区。

### 2.2 互斥锁操作

```go
var mutex sync.Mutex

// 加锁
mutex.Lock()

// 临界区代码
// ...

// 解锁
mutex.Unlock()
```

## 3. 代码示例分析

### 3.1 数据竞争演示

```go
package test_lock

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup
var count1 = 0

func Demo01() {
	wait.Add(10)
	for i := 0; i < 10; i++ {
		go func(data *int) {
			// 模拟访问耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			// 访问数据
			temp := *data
			// 模拟计算耗时
			// 拉大时间，让数据竞争更严重（更好观察结果）
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			ans := 1
			// 修改数据
			*data = temp + ans
			fmt.Println("goroutine", i, "count1:", *data)
			wait.Done()
		}(&count1)
	}
	wait.Wait()
	fmt.Println("最终结果", count1)
}
```

**核心要点：**

1. **无锁并发访问**：
   
   ```go
   var count1 = 0
   
   go func(data *int) {
       temp := *data           // 读取共享变量
       // ... 计算耗时
       *data = temp + ans      // 写入共享变量
   }(&count1)
   ```
   
2. **数据竞争特征**：
   
   - 多个goroutine同时读写`count1`变量
   - 读-改-写操作不是原子的
   - 最终结果通常小于预期值

**关键：**

- 典型的数据竞争
- 由于goroutine执行顺序不确定，结果不可预测
- 需要同步机制来保证正确性

### 3.2 互斥锁解决方案

```go
package test_lock

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count = 0

// 声明一个互斥锁
var mutex sync.Mutex

func Demo02() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(data *int) {
			// 加锁
			mutex.Lock()

			// 模拟访问耗时
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			// 访问数据
			temp := *data
			// 模拟计算耗时
			// 拉大时间，让数据竞争更严重（更好观察结果）
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			ans := 1
			// 修改数据
			*data = temp + ans

			// 解锁
			mutex.Unlock()

			fmt.Println("goroutine", i, "count:", *data)
			wg.Done()
		}(&count)
	}
	wg.Wait()
	fmt.Println("最终结果", count)
}
```

**核心要点：**

1. **互斥锁保护临界区**：
   ```go
   var mutex sync.Mutex
   
   go func(data *int) {
       // 加锁
       mutex.Lock()
       
       temp := *data
       // ... 计算耗时
       *data = temp + ans
       
       // 解锁
       mutex.Unlock()
   }(&count)
   ```

2. **锁的正确使用**：
   - 在访问共享资源前加锁
   - 在操作完成后立即解锁
   - 使用`defer`确保解锁被执行

**关键：**

- 互斥锁保证了操作的原子性
- 最终结果总是正确的（10个goroutine，每个+1，结果应为10）
- 性能会有一定损耗，但保证了正确性

### 3.3 读写锁（RWMutex）

```go
package test_lock

import (
	"fmt"
	"sync"
	"time"
)

var wg3 sync.WaitGroup
var rw sync.RWMutex

// 实现可以多人读  但是只能一人写
func write(count int) {
	defer wg3.Done()
	rw.Lock()
	defer rw.Unlock()
	fmt.Println("goroutine", count, "写操作>>>>>>>")
	time.Sleep(time.Second * 2)
}

func read(count int) {
	defer wg3.Done()
	fmt.Println("goroutine", count, "<<<<<<<<<读操作")
	time.Sleep(time.Second * 2)
}

func Demo03() {
	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go write(i)
	}

	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go read(i)
	}

	wg3.Wait()
}
```

**核心要点：**

1. **读写锁操作**：
   ```go
   var rw sync.RWMutex
   
   // 写锁操作
   func write(count int) {
       rw.Lock()           // 加写锁
       defer rw.Unlock()   // 解读锁
       // 写操作...
   }
   
   // 读锁操作  
   func read(count int) {
       rw.RLock()          // 加读锁
       defer rw.RUnlock()  // 解读锁
       // 读操作...
   }
   ```

2. **读写锁特性**：
   - **读锁**：多个goroutine可以同时持有读锁
   - **写锁**：写锁是排他的，持有写锁时不能有读锁或其他写锁
   - **升级规则**：读锁不能升级为写锁

**关键：**

- 读写锁适用于"读多写少"的场景
- 提高了并发读取的性能
- 写操作仍然需要独占访问

## 4. 锁的类型对比

### 4.1 互斥锁 vs 读写锁

| 特性 | 互斥锁（Mutex） | 读写锁（RWMutex） |
|------|----------------|------------------|
| 并发读取 | 不支持 | 支持多个goroutine同时读取 |
| 写入操作 | 完全互斥 | 完全互斥 |
| 性能开销 | 较低 | 较高（需要维护读锁计数） |
| 适用场景 | 读写频率相当 | 读多写少 |

### 4.2 选择原则

**使用互斥锁的情况：**
- 读写操作频率相当
- 临界区代码执行时间短
- 代码逻辑简单

**使用读写锁的情况：**
- 读操作远多于写操作
- 读操作耗时较长
- 需要最大化读取并发性

## 5. 锁的使用场景

### 5.1 共享数据保护

```go
// 保护共享map
var dataMap = make(map[string]string)
var mapMutex sync.RWMutex

func GetValue(key string) string {
    mapMutex.RLock()
    defer mapMutex.RUnlock()
    return dataMap[key]
}

func SetValue(key, value string) {
    mapMutex.Lock()
    defer mapMutex.Unlock()
    dataMap[key] = value
}
```

### 5.2 计数器保护

```go
// 原子计数器
var counter int
var counterMutex sync.Mutex

func Increment() {
    counterMutex.Lock()
    defer counterMutex.Unlock()
    counter++
}

func GetCount() int {
    counterMutex.Lock()
    defer counterMutex.Unlock()
    return counter
}
```

### 5.3 资源池管理

```go
// 连接池管理
type ConnectionPool struct {
    pool []*Connection
    mutex sync.Mutex
}

func (p *ConnectionPool) Get() *Connection {
    p.mutex.Lock()
    defer p.mutex.Unlock()
    
    if len(p.pool) > 0 {
        conn := p.pool[0]
        p.pool = p.pool[1:]
        return conn
    }
    return nil
}
```

## 6. 指南

### 6.1 锁的使用

1. **保持锁的粒度适中**：
   
   ```go
   // 错误：锁的粒度过大
   mutex.Lock()
   // 大量非临界区代码...
   // 共享资源访问
   mutex.Unlock()
   
   // 正确：只保护必要的临界区
   // 非临界区代码...
   mutex.Lock()
   // 共享资源访问
   mutex.Unlock()
   // 非临界区代码...
   ```
   
2. **使用defer确保解锁**：
   ```go
   func safeOperation() {
       mutex.Lock()
       defer mutex.Unlock()  // 确保在任何情况下都会解锁
       
       // 可能panic的代码
       if err != nil {
           panic("error")
       }
   }
   ```

### 6.2 避免死锁

1. **锁的顺序一致性**：
   ```go
   // 错误：可能产生死锁
   func operation1() {
       mutexA.Lock()
       mutexB.Lock()
       // ...
       mutexB.Unlock()
       mutexA.Unlock()
   }
   
   func operation2() {
       mutexB.Lock()  
       mutexA.Lock()  // 可能死锁
       // ...
   }
   
   // 正确：保持一致的锁顺序
   func operation1() {
       mutexA.Lock()
       mutexB.Lock()
       // ...
   }
   
   func operation2() {
       mutexA.Lock()  // 先获取A锁
       mutexB.Lock()  // 再获取B锁
       // ...
   }
   ```

2. **使用超时机制**：
   ```go
   // 使用context实现超时
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   
   select {
   case <-acquireLock(ctx, mutex):
       // 成功获取锁
   case <-ctx.Done():
       // 超时处理
   }
   ```

### 6.3 性能优化

1. **减少锁竞争**：
   - 使用更细粒度的锁
   - 考虑使用无锁数据结构
   - 使用本地缓存减少锁的使用频率

2. **读写分离**：
   ```go
   // 使用读写锁优化读多写少的场景
   var rw sync.RWMutex
   var data []string
   
   // 多个goroutine可以同时读取
   func ReadData() []string {
       rw.RLock()
       defer rw.RUnlock()
       return data
   }
   
   // 写操作仍然需要互斥
   func WriteData(newData []string) {
       rw.Lock()
       defer rw.Unlock()
       data = newData
   }
   ```

## 7. 高级主题

### 7.1 条件变量（Cond）

条件变量用于在特定条件下等待或通知goroutine：

```go
var mutex sync.Mutex
cond := sync.NewCond(&mutex)

// 等待条件
func waitForCondition() {
    mutex.Lock()
    for !condition {
        cond.Wait()  // 释放锁并等待
    }
    // 条件满足，执行操作
    mutex.Unlock()
}

// 通知条件满足
func signalCondition() {
    mutex.Lock()
    condition = true
    cond.Signal()  // 通知一个等待的goroutine
    mutex.Unlock()
}
```

### 7.2 原子操作

对于简单的计数器操作，可以使用原子操作避免锁的开销：

```go
import "sync/atomic"

var counter int64

func Increment() {
    atomic.AddInt64(&counter, 1)
}

func GetCount() int64 {
    return atomic.LoadInt64(&counter)
}
```

## 8. 调试和测试

### 8.1 数据竞争检测

使用Go内置的竞争检测器：

```bash
# 编译时启用竞争检测
go build -race main.go

# 运行程序
go run -race main.go
```

### 8.2 性能分析

使用pprof分析锁竞争：

```go
import _ "net/http/pprof"

// 在程序中启动pprof服务器
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```
