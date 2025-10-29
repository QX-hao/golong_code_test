package main

import (
	"fmt"
	"strings"
)

// 定义空接口
// 方法一：使用type关键字定义空接口
type Any interface {
	// 空接口没有方法，所有类型都实现了空接口
}

// 根据不同数据类型进行操作
// 方法一：
func MyPrint(x interface{}) {
	if v, ok := x.(string); ok {
		fmt.Println("变量x的类型为：string，值为：", v)
	} else if v, ok := x.(int); ok {
		fmt.Println("变量x的类型为：int，值为：", v)
	} else {
		fmt.Println("变量x的类型不是string或int")
	}
}

// 方法二：
// x.(type)只能结合switch来使用
func MyPrint2(x interface{}) {
	// 断言
	switch x.(type) {
		case nil:
			fmt.Println("变量x的类型为：nil")
		case string:
			fmt.Println("变量x的类型为：string，值为：", x.(string))
		case int:
			fmt.Println("变量x的类型为：int，值为：", x.(int))
		default:
			fmt.Println("变量x的类型不是string或int")
	}
}

func main()  {
	// 使用空接口

	var a Any = 10
	fmt.Printf("类型为%T\n变量a的值为：%v", a, a)


	// 方法二：直接定义
	var a2  interface{} = "hello"
	fmt.Printf("类型为%T\n变量a2的值为：%v\n", a2, a2)

	str,ok := a2.(string)
	if ok {
		fmt.Println("变量a2的类型为：string，值为：", str)
	} else {
		fmt.Println("变量a2的类型不是string")
	}

	fmt.Println(strings.Repeat("-",20))

	MyPrint("hello")
	MyPrint(10)
	MyPrint(10.5)

	MyPrint2("xiaopan")
	MyPrint2(11)
	MyPrint2(11.5)
	
	
}