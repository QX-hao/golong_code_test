package main

import (
	"fmt"
)

func main() {
	num := 2
	p := &num
	rawNum := *p
	fmt.Println(rawNum)

	
}