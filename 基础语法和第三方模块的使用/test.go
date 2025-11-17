package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string
	Age int
	Address City
}

type City struct {
	Address string
}


func main() {
	// num := 2
	// p := &num
	// rawNum := *p
	// fmt.Println(rawNum)

	// fmt.Printf("num的内存地址：%p\n", &num)
	// fmt.Printf("num的值：%p\n", num)

	
}