package test_lock

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count = 0


// Lock()
    // 在这个过程中，数据不会被其他协程修改
// Unlock()

// 声明一个互斥锁
var mutex sync.Mutex

func Demo02()  {
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

		fmt.Println("goroutine",i,"count:",*data)
		wg.Done()
	}(&count)
	}
	wg.Wait()
	fmt.Println("最终结果", count)
}
