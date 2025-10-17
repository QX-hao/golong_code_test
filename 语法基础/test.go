package main

import (
	"fmt"
	"github.com/QX-hao/golang_code_test/algorithm"
)

func main() {
	num := 2
	p := &num
	rawNum := *p
	fmt.Println(rawNum)
	
}