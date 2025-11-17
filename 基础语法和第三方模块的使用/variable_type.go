package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a_int int = 100
	var a_int8 int8 = 10 
	var a_int32 int32 = 1000
	var a_int64 int64 = 10000

	fmt.Println(a_int,reflect.TypeOf(a_int))
	fmt.Println(a_int8,reflect.TypeOf(a_int8))
	fmt.Println(a_int32,reflect.TypeOf(a_int32))
	fmt.Println(a_int64,reflect.TypeOf(a_int64))

	// 当不指明变量类型时，默认类型为int(根据电脑配置来确定为32位或64位)
	var a = 1
	fmt.Println(a,reflect.TypeOf(a))

	fmt.Println("-----------------")

	//格式化输出-进制转换
	var b = 10

	// 转换成十进制
	fmt.Printf("十进制：%d \n",b)
	//转换成二进制
	fmt.Printf("二进制：%b \n",b)
	//转换成八进制
	fmt.Printf("八进制：%o \n",b)
	//转换成十六进制
	fmt.Printf("十六进制：%x \n",b)

	fmt.Println("-----------------")
	//浮点型格式化输出
	//float64精度大概15-16位
	//float32精度大概6-7位
	fmt.Printf("浮点型格式化输出")

	fmt.Printf("默认格式：%f \n", 3.1415926)
	fmt.Printf("保留2位小数：%.2f \n", 3.1415926)
	fmt.Printf("保留3位小数：%.3f \n", 3.1415926)
	fmt.Printf("科学计数法：%e \n", 3.1415926)
	fmt.Printf("保留2位小数的科学计数法：%.2e \n", 3.1415926)
	fmt.Printf("保留3位小数的科学计数法：%.3e \n", 3.1415926)

	// 格式化输出-字符串
	fmt.Printf("默认格式：%s \n", "hello world")
	fmt.Printf("保留2位：%.2s \n", "hello world")
	fmt.Printf("保留3位：%.3s \n", "hello world")

	// 格式化输出-布尔型
	fmt.Printf("默认格式：%t \n", true)
	fmt.Printf("保留2位：%.2t \n", true)
	fmt.Printf("保留3位：%.3t \n", true)

	//布尔型类型
	var boo = 2>3
	fmt.Println(boo,reflect.TypeOf(boo))




}