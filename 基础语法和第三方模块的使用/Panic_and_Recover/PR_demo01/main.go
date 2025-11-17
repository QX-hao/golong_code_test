package main

import (
	"fmt"
)

// panic 会导致程序崩溃，但是可以通过 recover 来恢复程序的执行
// panic 可以在任何地方引发 
// recover 只能在 defer 中调用，否则会导致编译错误

func fn1() {
	fmt.Println("start")
}

func fn2() {
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	panic("我发生了错误")
}

func main() {
	fn1()
	fn2()
	fmt.Println("ending")
}