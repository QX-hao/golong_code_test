package main

import (
	"runtime"
	"fmt"
)

func main() {
	// 打印当前cpu数量
	fmt.Println("当前协程数量：",runtime.NumCPU())

	// 设置使用cpu数量
	runtime.GOMAXPROCS(2)

}