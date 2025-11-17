package main

import (
	"fmt"
)

func main() {
	// 初始长度为5，预分配容量为10，为后续追加预留空间
	initialLength := 5
	preAllocatedCapacity := 10
	slice1 := make([]int, initialLength, preAllocatedCapacity)
	
	fmt.Println("=== 切片初始状态 ===")
	fmt.Printf("slice1 内容: %v\n", slice1)
	fmt.Printf("slice1 长度: %d, 容量: %d\n", len(slice1), cap(slice1))
	fmt.Printf("slice1 变量地址: %p\n", &slice1)
	fmt.Printf("slice1 底层数组地址: %p\n", slice1)
	
	// 优化2：批量追加元素，减少函数调用次数
	fmt.Println("\n=== 批量追加元素 ===")
	elementsToAppend := []int{10, 20, 30}
	slice1 = append(slice1, elementsToAppend...)
	
	fmt.Printf("追加后 slice1 内容: %v\n", slice1)
	fmt.Printf("追加后 slice1 长度: %d, 容量: %d\n", len(slice1), cap(slice1))
	fmt.Printf("追加后 slice1 变量地址: %p\n", &slice1)
	fmt.Printf("追加后 slice1 底层数组地址: %p\n", slice1)
	
	// 优化3：创建新切片时考虑容量共享问题
	fmt.Println("\n=== 创建新切片 slice2 ===")
	slice2 := make([]int, len(slice1), cap(slice1)+1) // 预分配额外容量
	copy(slice2, slice1) // 显式复制数据，避免共享底层数组
	slice2 = append(slice2, 100)
	
	fmt.Printf("slice2 内容: %v\n", slice2)
	fmt.Printf("slice2 长度: %d, 容量: %d\n", len(slice2), cap(slice2))
	fmt.Printf("slice2 变量地址: %p\n", &slice2)
	fmt.Printf("slice2 底层数组地址: %p\n", slice2)
	
	// 优化4：验证切片独立性
	fmt.Println("\n=== 验证切片独立性 ===")
	fmt.Println("修改 slice1[0] 为 1000")
	slice1[0] = 1000
	
	fmt.Printf("修改后 slice1 内容: %v\n", slice1)
	fmt.Printf("修改后 slice2 内容: %v\n", slice2)
	
	// 优化5：性能对比演示
	fmt.Println("\n=== 性能优化对比 ===")
	demonstratePerformanceOptimization()
	
	// 优化6：内存使用分析
	fmt.Println("\n=== 内存使用分析 ===")
	analyzeMemoryUsage()
}

// 优化7：提取独立函数，提高代码可维护性
func demonstratePerformanceOptimization() {
	// 方法1：不预分配容量（性能较差）
	var inefficientSlice []int
	for i := 0; i < 1000; i++ {
		inefficientSlice = append(inefficientSlice, i)
	}
	
	// 方法2：预分配容量（性能较好）
	efficientSlice := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		efficientSlice = append(efficientSlice, i)
	}
	
	fmt.Printf("不预分配容量 - 最终容量: %d (扩容次数较多)\n", cap(inefficientSlice))
	fmt.Printf("预分配容量   - 最终容量: %d (无扩容)\n", cap(efficientSlice))
}

func analyzeMemoryUsage() {
	// 演示不同创建方式的内存使用
	sliceA := make([]int, 100)      // 长度100，容量100
	sliceB := make([]int, 0, 100)   // 长度0，容量100
	sliceC := make([]int, 50, 100)  // 长度50，容量100
	
	fmt.Printf("sliceA - 长度: %d, 容量: %d (已使用全部容量)\n", len(sliceA), cap(sliceA))
	fmt.Printf("sliceB - 长度: %d, 容量: %d (预留增长空间)\n", len(sliceB), cap(sliceB))
	fmt.Printf("sliceC - 长度: %d, 容量: %d (部分使用，预留空间)\n", len(sliceC), cap(sliceC))
	
	// 内存优化建议
	fmt.Println("\n内存优化建议:")
	fmt.Println("1. 根据实际需求预分配容量")
	fmt.Println("2. 避免不必要的切片复制")
	fmt.Println("3. 及时释放不再使用的大切片")
	fmt.Println("4. 使用 copy() 替代 append() 进行数据迁移")
}

// 优化8：添加辅助函数，提高代码复用性
func printSliceInfo(name string, slice []int) {
	fmt.Printf("%s - 内容: %v, 长度: %d, 容量: %d\n", 
		name, slice, len(slice), cap(slice))
}

func createOptimizedSlice(initialData []int, preAllocatedCap int) []int {
	if preAllocatedCap < len(initialData) {
		preAllocatedCap = len(initialData)
	}
	
	result := make([]int, len(initialData), preAllocatedCap)
	copy(result, initialData)
	return result
}