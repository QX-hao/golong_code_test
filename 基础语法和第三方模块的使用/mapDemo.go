package main

import (
	"fmt"
)

func main() {
	var stu = map[string]map[string]int{
		"stu1": {
			"age":   18,
			"score": 90,
		},
		"stu2": {
			"age":   19,
			"score": 85,
		},
	}

	// 遍历map
	for key, value := range stu {
		fmt.Printf("学生 %s: 年龄 %d, 成绩 %d\n", key, value["age"], value["score"])
	}
}
