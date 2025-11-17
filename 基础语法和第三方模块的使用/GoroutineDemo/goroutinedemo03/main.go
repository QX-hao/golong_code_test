package main

import (
	"fmt"
	"time"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func Groutine1() {
	defer wg.Done()
	wg.Add(1)
	for num := 2; num<30000; num++ {
		flag := true
		for i := 2; i<num ; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(num,"是质数")
		}
	}
}

func Groutine2() {
	defer wg.Done()
	wg.Add(1)
	for num := 30000; num<60000; num++ {
		flag := true
		for i := 2; i<num ; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(num,"是质数")
		}
	}
}

func Groutine3() {
	defer wg.Done()
	wg.Add(1)
	for num := 60000; num<100000; num++ {
		flag := true
		for i := 2; i<num ; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(num,"是质数")
		}
	}
}


func main() {
	fmt.Println(strings.Repeat("-",20))
	var time1 = time.Now()
	for num := 2; num<100000; num++ {
		flag := true
		for i := 2; i<num ; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(num,"是质数")
		}
	}
	var time2 = time.Now()
	
	fmt.Println(strings.Repeat("-",20))
	

	fmt.Println(strings.Repeat("-",20))
	var time3 = time.Now()
	go Groutine1()
	go Groutine2()
	go Groutine3()
	var time4 = time.Now()
	// 等待协程执行完成  等待协程计数器为0
	wg.Wait()
		fmt.Println("无协程耗时：", time2.Sub(time1))
	fmt.Println("开启协程耗时：", time4.Sub(time3))

}