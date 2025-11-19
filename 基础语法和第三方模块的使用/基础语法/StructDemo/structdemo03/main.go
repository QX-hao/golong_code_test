package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名:%s,年龄:%d,成绩:%d", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func DeBug() {
	fmt.Println(strings.Repeat("-", 20))
}

func main() {
	stu1 := Student{
		Name:  "张三",
		Age:   18,
		Score: 90,
	}

	// 取址 -- 改成 引用类型
	stu2 := &Student{
		Name:  "李四",
		Age:   19,
		Score: 80,
	}
	fmt.Println(stu1.GetInfo())
	DeBug()
	fmt.Println(stu2.GetInfo())
	DeBug()
	stu3 := stu2
	stu3.SetInfo("王五", 20, 70)
	fmt.Println(stu2.GetInfo())
}
