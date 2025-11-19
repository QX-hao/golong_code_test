package main

import (
	"goroutine_mutexdemo01/test_lock"
	"fmt"
	"strings"
)

func main() {

	// 不加锁
	test_lock.Demo01()
	fmt.Println(strings.Repeat("-", 50))
	// 加入互斥锁
	test_lock.Demo02()
	fmt.Println(strings.Repeat("-", 50))

	// 读写锁
	test_lock.Demo03()
	fmt.Println(strings.Repeat("-", 50))
}
