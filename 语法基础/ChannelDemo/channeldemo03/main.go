package main

import (
	"fmt"
	"time"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// 向inChan中存放数字
func putNum(inChan chan int)  {
	for i := 2; i<120000; i++ {
		inChan <- i
	}
}

// 从inChan中取数字，判断是否为素数，是则放入primeChan
func primeNum()  {

}

// printPrime 打印素数的方法
func printPrime()  {
	
}


func main() {

	// 存放数字
	inChan := make(chan int,1000)

	// 存放素数
	primeChan := make(chan int,1000)

	// 退出信号
	exitChan := make(chan int,1000)

}