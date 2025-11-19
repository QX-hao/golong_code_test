package main

import (
	"fmt"
)

// 闭包：函数内部定义的函数，对外部作用域的变量进行引用
// 写法： 函数里面嵌套函数，并且最后返回函数

func add1() func() int {
	i := 10 // 常驻内存，不污染全局
	return func() int {
		return  i + 1
	}
}


func add2() func(y int) int {
	i := 10
	return func(y int) int {
		i += y // 每次调用闭包，i会累加y
		return i 
	}
}

// 闭包的作用： 可以在函数外部访问函数内部的变量
// 闭包的注意事项： 闭包会引用外部的变量，所以外部的“变量不能被垃圾回收“，否则会导致内存泄漏
// 解决方法： 可以在闭包内部定义一个变量，将外部的变量赋值给这个变量，然后在闭包内部使用这个变量
// 闭包的应用场景： 可以用于实现计数器、缓存、延迟计算等功能

func main() {

	f := add1() // 执行闭包，返回一个函数
	fmt.Printf("f的值为：%v，数据类型为：%T\n", f, f) // f的值为：0x79c120，数据类型为：func() int


	fmt.Printf("f()的值为：%v\n", f()) // f()的值为：11
	fmt.Printf("f()的值为：%v\n", f()) // f()的值为：11
	fmt.Printf("f()的值为：%v\n", f()) // f()的值为：11


	g := add2() // 执行闭包，返回一个函数
	fmt.Printf("g的值为：%v，数据类型为：%T\n", g, g) // g的值为：0x79c120，数据类型为：func(int) int

	fmt.Printf("g(1)的值为：%v\n", g(10)) // g(1)的值为：20
	fmt.Printf("g(2)的值为：%v\n", g(10)) // g(2)的值为：30
	fmt.Printf("g(3)的值为：%v\n", g(10)) // g(3)的值为：40
}