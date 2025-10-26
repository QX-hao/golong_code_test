package main

import (
	"fmt"
	"time"
)

// golong定时器
func main()  {

	ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()


	num := 5
	for key := range ticker.C {	
		fmt.Printf("当前时间：%v\n", key)
		num--
		if num <= 0 {
			// 终止定时器
			ticker.Stop()
			break
		}
	}


// sleep()

fmt.Printf("程序开始%v\n", time.Now())
time.Sleep(time.Second * 5)
fmt.Printf("程序结束%v\n", time.Now())


}