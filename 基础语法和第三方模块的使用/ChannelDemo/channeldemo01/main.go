package main

import (
	"fmt"
	// "time"
)

func main()  {
	// channel的声明--队列（先进先出）
	// ch1 := make(chan int,3)
	var ch chan int
	ch = make(chan int,3)

	// 给管道传数据
	ch <- 1
	ch <- 2
	ch <- 3

	// 从管道取数据
	a := <-ch
	b := <-ch
	c := <-ch
	fmt.Println("从管道取的数据：",a)
	fmt.Println("从管道取的数据：",b)
	fmt.Println("从管道取的数据：",c)

	// ch的长度和容量
	fmt.Println("ch的长度：",len(ch))
	fmt.Println("ch的容量：",cap(ch))
	

	// channel无key
	// channel的遍历
	ch1 := make(chan int,3)
	ch1 <- 11
	ch1 <- 22
	ch1 <- 33
	// 关闭管道 -- 不关闭 for range 会报错deadlock,因为无索引range会取完数据后会在取一次
	// 也就是比len多取一次
	close(ch1)

	// for i := 0; i < len(ch1); i++ {
	// 	fmt.Println("遍历管道：",<-ch1)
	// }

	for v := range ch1 {
		fmt.Println("遍历管道：",v)
	}
}