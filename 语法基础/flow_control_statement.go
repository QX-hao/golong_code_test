package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main()  {
	var is_login bool = true

	if is_login {
		fmt.Println("用户已登录")
	} else {
		fmt.Println("用户未登录")
	}


	// 多分支选择语句
	fmt.Println(strings.Repeat("-",10))
	fmt.Println("多分支选择语句")
	var score int
	fmt.Print("请输入成绩（0-100）：")
	fmt.Scanln(&score)

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	// 
	var a = 10
	if reflect.TypeOf(a).String() == "int" {
		fmt.Println("a是整型")
	}

	var choice int
	fmt.Print("请输入操作选择（1-5）：")
	fmt.Scanln(&choice)

	switch choice {
	case 1:fmt.Println("普通攻击")
	case 2:fmt.Println("使用道具")
	case 3:fmt.Println("超级攻击")
	case 4:fmt.Println("使用技能")
	case 5:fmt.Println("退出游戏")
	default:fmt.Println("选择了其他选项")
	}


	//循环语句
	fmt.Println(strings.Repeat("-",10))
	fmt.Println("循环语句")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for true {
	// 	fmt.Println("无限循环")
	// }

	// 1-100的和
	var sum int 
	for i := 1; i <= 100; i++ {
		sum +=i
	}
	fmt.Println("1-100的和为：",sum)

}
