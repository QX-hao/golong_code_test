package main

import (
	"fmt"
	"sync"
	"time"
)

// groutine结合channel
// 1. 当channel没有数据时，读取会阻塞
// 2. 当channel没有空间时，写入会阻塞
// 3. 当channel关闭时，读取会返回0值，写入会报错

var wg sync.WaitGroup

func input(ch chan int)  {
	defer wg.Done()
	
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("向管道放入的数据：",i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch)
}

func output(ch chan int)  {
	defer wg.Done()
	for v := range ch {
		fmt.Println("从管道取的数据：",v)
		time.Sleep(time.Millisecond * 50)
	}
}

func main()  {
	
	var ch = make(chan int,5)
	wg.Add(1)
	go input(ch)
	wg.Add(1)
	go output(ch)
	wg.Wait()
	fmt.Println("主函数结束")
}