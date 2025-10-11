package main

import (
	"fmt"
)

func main() {
	// go中声明变量后必须用上，否则会报错，或者用匿名变量_来忽略
	var _,b = 10,20
	fmt.Println(b)
}
