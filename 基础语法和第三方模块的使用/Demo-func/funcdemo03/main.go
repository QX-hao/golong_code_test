package main

import (
	"fmt"
)

func f3() (a int) {
	defer func() {
		a++
	}()
	return a
}

func f5() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	fmt.Println(f3())
	fmt.Println(f5())

	// defer在执行之前就要计算参数--也就是在注册的时候已经将参数计算好
	// 所以这里的 calc("A", x, y) 会立即执行，得到返回值 3
	// 然后将 3 作为参数传递给 calc("AA", x, 3)
	// 同理，calc("B", x, y) 会立即执行，得到返回值 12
	// 然后将 12 作为参数传递给 calc("BB", x, 12)

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

	/*
		注册顺序
			defer calc("AA", x, calc("A", x, y))
			defer calc("BB", x, calc("B", x, y))
		执行顺序
			defer calc("BB", x, calc("B", x, y))
			defer calc("AA", x, calc("A", x, y))
	*/

	// 1.calc("A", x, y)
	// 2.calc("B", x, y)
	// 3.calc("BB", x, 12)
	// 4.calc("AA", x, 3)
}
