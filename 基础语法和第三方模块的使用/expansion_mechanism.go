package main

import "fmt"

func main() {
	// 面试题
	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2]
	s2 := s1
	s3 := append(append(append(s1, 1), 2), 3)
	s1[0] = 1000
	
	fmt.Println("=== 内存地址分析 ===")
	fmt.Printf("arr 地址: %p\n", &arr)
	fmt.Printf("s1 变量地址: %p, 底层数组地址: %p\n", &s1, s1)
	fmt.Printf("s2 变量地址: %p, 底层数组地址: %p\n", &s2, s2)
	fmt.Printf("s3 变量地址: %p, 底层数组地址: %p\n", &s3, s3)
	
	fmt.Println("\n=== 数据内容 ===")
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("s1: %v (长度: %d, 容量: %d)\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v (长度: %d, 容量: %d)\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: %v (长度: %d, 容量: %d)\n", s3, len(s3), cap(s3))
	
	fmt.Println("\n=== 关键验证 ===")
	fmt.Printf("s2[0] = %d (应该是1000，因为s2和s1共享底层数组)\n", s2[0])
	fmt.Printf("s3[0] = %d (没有引用s1的元素，所以不是1000)\n", s3[0])
}
