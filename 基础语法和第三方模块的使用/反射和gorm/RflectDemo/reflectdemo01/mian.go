package main

import (
	"fmt"
	"reflect"
	"strings"
)

type MyInt int
type any interface{}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}



func PrintType(t interface{})  {
	fmt.Println(t,"类型是",reflect.TypeOf(t))

	// 类型名称  Name()
	fmt.Println(t,"的类型名称是",reflect.TypeOf(t).Name())
	// 种类   Kind()  --  底层类型
	fmt.Println(t,"的种类是",reflect.TypeOf(t).Kind())
	fmt.Println(strings.Repeat("-",20))
}


func main()  {
	a := 10
	PrintType(a)

	b := "你好 golong"
	PrintType(b)

	c := true
	PrintType(c)

	var d MyInt  = 100
	PrintType(d)	// main.MyInt

	var e any = "qx-hao is cool"
	PrintType(e) 

	var f = Person{
		Name: "qx-hao",
		Age:  18,
	}
	PrintType(f)

	// 指针 切片 数组 的类型名称都是 空字符串
	var g *int = &a
	PrintType(g)

	var h []int = []int{1,2,3,4,5}
	PrintType(h)

	var i [3]int = [3]int{1,2,3}
	PrintType(i)
}