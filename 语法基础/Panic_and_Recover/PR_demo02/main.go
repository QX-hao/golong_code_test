package main

import (
	"fmt"
)

func fn1(a,b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("计算错误:", err)
		}
	}()
	return a/b
}

func main()  {
	fmt.Println(fn1(10,0))
	fmt.Println("ending")
}