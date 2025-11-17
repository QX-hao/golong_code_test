package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func a() {
	fmt.Println("a()")
}


func main()  {
	var user Student
	// user.Name = "张三"
	// user.Age = 18
	fmt.Printf("类型：%T\n",user)

	
}