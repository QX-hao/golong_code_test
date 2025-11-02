package main

import (
	"fmt"
	"strings"
	"time"
)

func main()  {
	
	intChan := make(chan int, 10)
	for i := 0; i<10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 10)
	for i := 0; i<10; i++ {
		stringChan <- fmt.Sprintf("hello%d",i)
	}

	// func () {
	// 	close(intChan)
	// 	close(stringChan)
	// }()

	// select -- IO 多路复用的解决方案
	// select和switch的区别：
	// 1. select会随机执行一个case，而switch会顺序执行
	// 2. select可以监听多个channel，而switch只能监听一个channel


	// select 不需要close channel
	for {
		select {
			case num := <-intChan:
				fmt.Println("intChan:",num)
				time.Sleep(time.Millisecond * 50)
			case str := <-stringChan:
				fmt.Println("stringChan:",str)
				time.Sleep(time.Millisecond * 50)
			default:
				fmt.Println("数据获取完毕")
				return 
		}
	}
	
	fmt.Println(strings.Repeat("-",20))
}