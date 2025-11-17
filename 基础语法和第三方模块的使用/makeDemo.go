package main

import "fmt"

func appendDemo()  {
	var slice1 = make([]int, 5)

	slice2 := append(slice1, 100)
	fmt.Println(slice1)
	fmt.Println(slice2)

}

func main() {
	// 创建长度为 3，容量为 5 的整型切片
	slice1 := make([]int, 3, 5)
	fmt.Printf("slice1: %v, len: %d, cap: %d\n", slice1, len(slice1), cap(slice1))
	// 输出: slice1: [0 0 0], len: 3, cap: 5

	// 创建长度和容量都为 3 的字符串切片
	slice2 := make([]string, 3)
	fmt.Printf("slice2: %v, len: %d, cap: %d\n", slice2, len(slice2), cap(slice2))
	// 输出: slice2: [  ], len: 3, cap: 3

	// 不指定容量，容量默认等于长度
	slice3 := make([]int, 3)
	fmt.Printf("slice3: %v, len: %d, cap: %d\n", slice3, len(slice3), cap(slice3))
	// 输出: slice3: [0 0 0], len: 3, cap: 3

	//
	appendDemo()

}
