package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明
	var arr_int [5]int = [5]int{1, 2, 3, 4, 5}
	var arr_string [3]string

	fmt.Println(arr_int, arr_string)
	fmt.Println("数组数据类型：", reflect.TypeOf(arr_int))

	// 切片
	// 方法一：切割数组获取
	slice_int := arr_int[:2] // 省略起始索引 0，默认从数组开头开始切割
	fmt.Printf("arr_int[:2] 切片内容：%v，切片数据类型：%T\n", slice_int, reflect.TypeOf(slice_int))
	// 方法二：直接声明
	var slice_int2 []int
	fmt.Println("切片数据类型：", reflect.TypeOf(slice_int2))

	var arr_test = [6]int{1, 2, 3, 4, 5, 6}
	arr3 := arr_test[2:]
	fmt.Println(arr3)
	fmt.Println(len(arr3), cap(arr3))

	var arr4 = arr3[1:2]
	fmt.Println(len(arr4), cap(arr4))
}
