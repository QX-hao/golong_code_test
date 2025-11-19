package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// 向inChan中存放数字
func putNum(inChan chan int)  {
	defer wg.Done()
	for i := 2; i<120000; i++ {
		inChan <- i
	}
	close(inChan)
}

// 从inChan中取数字，判断是否为素数，是则放入primeChan
func primeNum(inChan chan int,primeChan chan int,exitChan chan bool)  {
	defer wg.Done()
	for num := range inChan {
		isNum := true
		for i := 2; i<num; i++ {
			if num %i == 0 {
				isNum = false
				break
			}
		}
		if isNum {
			primeChan <- num
		}
	}
	// close(primeChan) // 比较危险，关闭primeChan后就无法再从primeChan中发送和取数据
	// 利用exitChan通知printPrime方法退出
	exitChan <- true
}

// printPrime 打印素数的方法
func printPrime(primeChan chan int)  {
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

	for i := 0; i<10; i++ {
		wg.Add(1)
		go primeNum(inChan,primeChan,exitChan)
	}
	

	wg.Add(1)
	go printPrime(primeChan)

	// 判断exitChan是否存满
	wg.Add(1)
	go func () {
		for i := 0; i<10; i++ {
			<- exitChan
		}
		close(primeChan)
		wg.Done()
	}()
	

	wg.Wait()
	fmt.Println(strings.Repeat("-",20))
	

}