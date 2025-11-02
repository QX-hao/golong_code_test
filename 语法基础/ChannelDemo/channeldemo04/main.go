package main

import (
	"fmt"
	"strings"
	"sync"
)

// var ch chan int ---可以读写
// var ch chan<-int ---只能写
// var ch<-chan int ---只能读

var wg sync.WaitGroup

func main() {
	// 默认是双向通道
	fmt.Println("默认是双向通道")
	ch := make(chan int,2)
	ch <- 100
	fmt.Println(<-ch)

	fmt.Println(strings.Repeat("-",20))

	// 只能写
	fmt.Println("只能写")
	ch1 := make(chan<- int,2)
	ch1 <- 200

	// .\main.go:28:16: invalid operation: 
	// cannot receive from send-only channel ch1 (variable of type chan<- int)
	// fmt.Println(<-ch1) // 编译错误

	fmt.Println(strings.Repeat("-",20))
	// 只能读
	fmt.Println("只能读")
	ch2 := make(<-chan int,2)
	// .\main.go:37:2: 
	// invalid operation: cannot send to receive-only channel ch2 (variable of type <-chan int)
	// ch2 <- 300 // 编译错误
	fmt.Println(<-ch2)
}