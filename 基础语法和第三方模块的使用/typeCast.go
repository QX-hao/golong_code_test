package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	// 类型转换

	// 整型转字符串
	fmt.Println(strings.Repeat("-",10))
	fmt.Println("整型转字符串")
	var a int = 65
	b := string(a)
	fmt.Println(b)

	c := strconv.Itoa(a)
	fmt.Println(c,reflect.TypeOf(c))

	// 字符串转整型
	fmt.Println(strings.Repeat("-",10))
	fmt.Println("字符串转整型")
	var s2 string = "100"
	i,_ := strconv.Atoi(s2)
	fmt.Println(i,reflect.TypeOf(i))

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(i,reflect.TypeOf(i))


}