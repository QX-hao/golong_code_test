package main

import "fmt"

func main()  {
	var fn [10]func() // [f,f,f,f,f,f,f,f,f,f]

	for i := 0; i < len(fn); i++ {
		fn[i] = func(){
			fmt.Println(i)
		}
	}  // [f(0),f(1),f(2),f(3),f(4),f(5),f(6),f(7),f(8),f(9)]
	
	for _,f := range fn {
		f()
	}

	
}