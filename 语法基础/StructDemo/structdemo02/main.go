package main

import (
	"fmt"
	"encoding/json"
)

type Class struct {
	ClassName string
	Student []Student
}

type Student struct {
	Id int
	Name string
	Age int
}



func main() {
	c := Class{
		ClassName:"1班",
		Student: make([]Student, 0),
	}

	for i := 0; i<5; i++ {
		s := Student{
			Id: i,
			Name: fmt.Sprintf("学生%d", i),
			Age: 18,
		}
		c.Student = append(c.Student, s)
	}

	// fmt.Printf("%#v\n", c)
	jsonByte,err := json.Marshal(c)
	if err != nil {
		fmt.Println("json转换失败")
		return
	}else {
		fmt.Println(string(jsonByte))
	}

	// {"ClassName":"1班","Student":[{"Id":0,"Name":"学生0","Age":18},{"Id":1,"Name":"学生1","Age":18},{"Id":2,"Name":"学生2","Age":18},{"Id":3,"Name":"学生3","Age":18},{"Id":4,"Name":"学生4","Age":18}]}
}