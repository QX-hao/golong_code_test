package main

import (
	"fmt"
	"reflect"
)

func PrintValue(v interface{}) {

	switch reflect.ValueOf(v).Kind() {
		case reflect.Int:
			fmt.Println("num的值是", reflect.ValueOf(v).Int() + 100)
		default:
			fmt.Println( "还无其他类型")
	}
}

func Alter(v interface{})  {
	val := reflect.ValueOf(v)
	// 检查是否是指针类型
	if val.Kind() == reflect.Ptr {
		// 获取指针指向元素 -- 类似于 *v
		elem := val.Elem()
		// 检查指针指向元素是否是int类型
		if elem.Kind() == reflect.Int {
			elem.SetInt(100)
		}
	} else {
		fmt.Println("v不是指针类型")
	}
}

func Alter2(v interface{})  {
	if ptr, ok := v.(*int); ok {
		*ptr = 1000
	}
}

func main() {
	var a int = 10
	PrintValue(a)
	fmt.Println("a的值是:",a)
	Alter(&a)
	fmt.Println("Alter后:",a)
	Alter2(&a)
	fmt.Println("Alter2后:",a)
}


