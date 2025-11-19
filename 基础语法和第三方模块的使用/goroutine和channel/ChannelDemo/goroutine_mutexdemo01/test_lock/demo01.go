package test_lock

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup
var count1 = 0

func Demo01()  {
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
		fmt.Println("goroutine",i,"count1:",*data)
		wait.Done()
	}(&count1)
	}
	wait.Wait()
	fmt.Println("最终结果", count1)
}
