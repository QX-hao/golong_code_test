package main

import (
	"fmt"
	"time"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func Test() {
	for i := 0; i<10; i++ {
		fmt.Println("Test() 第",i,"次")
		time.Sleep(time.Microsecond * 50) // 50微秒
	}
}

func GoroutineDemo() {
	// 协程执行完成后，调用wg.Done()  协程计数器减1
	defer wg.Done()

	// 协程计数器加1
	wg.Add(1)
	for i := 0; i<10; i++ {
		fmt.Println("GoroutineDemo() 第",i,"次")
		time.Sleep(time.Microsecond * 50) // 50微秒
	}
}

func main()  {

	fmt.Println(strings.Repeat("-",20))
	fmt.Println("开启协程前")
	var time1 = time.Now()
	Test()
	GoroutineDemo()
	var time2 = time.Now()
	fmt.Println("耗时：", time2.Sub(time1))


	fmt.Println(strings.Repeat("-",20))
	fmt.Println("开启协程后")
	var time3 = time.Now()
	go GoroutineDemo()
	Test()
	var time4 = time.Now()
	fmt.Println("耗时：", time4.Sub(time3))
	// 等待协程执行完成  等待协程计数器为0
	wg.Wait()
	
}
