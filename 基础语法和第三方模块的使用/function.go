package main

import (
	"fmt"
)

func user_info() (name string, salary float64) {

	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入工资：")
	fmt.Scanln(&salary)

	return name, salary
}

func main() {
	name, salary := user_info()

	fmt.Println("-----------------")
	fmt.Println("用户信息如下：")
	fmt.Println("姓名：" + name)
	fmt.Println("工资：" + fmt.Sprintf("%f", salary))
}