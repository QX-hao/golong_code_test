package main	

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string
	Age int
	Address City
}

type City struct {
	Address string
}

func main() {
	var p1 Person
	p1.Name = "张三"
	p1.Age = 18
	p1.Address.Address = "中国"

	fmt.Printf("%v\n", p1)
	fmt.Printf("%#v\n", p1)


	// go对象(结构体)转json字符串
	jsonbyte, _ := json.Marshal(p1)
	// Bson
	// fmt.Println(jsonbyte)
	fmt.Println(string(jsonbyte))


	// json字符串转go对象(结构体)
	var p2 Person
	str := `{"name":"李四","age":28,"address":{"address":"中国"}}`


	// 第一个参数得是[]byte类型
	err := json.Unmarshal([]byte(str), &p2)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
	}
	fmt.Printf("%#v\n", p2)
}	