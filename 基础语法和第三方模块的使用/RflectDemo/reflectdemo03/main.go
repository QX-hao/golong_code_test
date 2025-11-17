package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Student struct {
	Name  string `json:"name" form:"name"`
	Age   int    `json:"age" form:"age"`
	Score int    `json:"score" form:"score"`
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

func (s Student) Abc_test() int {
	fmt.Println("测试Method()的取值问题")
	return 0
}

// 打印结构体所有字段
func PrintStuctField(v interface{}) {
	s := reflect.TypeOf(v)
	if s.Kind() != reflect.Struct && s.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体")
		return
	}

	// 通过Type去获取字段

	// 使用Field()获取结构体某一个字段  返回reflect.StructField类型(结构体)
	DeBug()
	fmt.Println("使用field获取") // 3
	fistField := s.Field(0)
	fmt.Println("Field返回值:", fistField) // {Name  string json:"name" 0 [0] false}
	fmt.Printf("%#v\n", fistField)      // reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x527280), Tag:"json:\"name\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	// 从reflect.StructField结构体中获取字段的名称
	fmt.Println("字段名称:", fistField.Name)
	// 从reflect.StructField结构体中获取字段的类型
	fmt.Println("字段类型:", fistField.Type)
	// 从reflect.StructField结构体中获取字段的标签
	fmt.Println("字段标签:", fistField.Tag)
	// 从reflect.StructField结构体中获取字段的标签中的json值
	fmt.Println("字段标签中的json值:", fistField.Tag.Get("json"))
	// 从reflect.StructField结构体中获取字段的标签中的form值
	fmt.Println("字段标签中的form值:", fistField.Tag.Get("form"))
	DeBug()

	// 使用FieldByName()
	fmt.Println("使用FieldByName()获取") // 3
	if field1, ok := s.FieldByName("Age"); !ok {
		fmt.Println("结构体中没有Age字段")
		return
	} else {
		fmt.Println("字段名称:", field1.Name)
		fmt.Println("字段类型:", field1.Type)
		fmt.Println("字段标签:", field1.Tag)
		fmt.Println("字段标签中的json值:", field1.Tag.Get("json"))
		fmt.Println("字段标签中的form值:", field1.Tag.Get("form"))
	}
	DeBug()

	// 使用NumField()获取结构体字段的数量
	fmt.Println("使用NumField()获取结构体字段的数量") // 3
	fmt.Println("结构体字段的数量:", s.NumField())
	DeBug()

	// 通过Value变量获取属性对应的值
	field := reflect.ValueOf(v)
	// 从reflect.Value结构体中获取字段对应的值
	fmt.Println("Name字段对应的值:", field.FieldByName("Name"))
	fmt.Println("Age字段对应的值:", field.FieldByName("Age"))
	fmt.Println("Score字段对应的值:", field.FieldByName("Score"))

	DeBug()

	// 遍历结构体的所有字段
	fmt.Println("遍历结构体的所有字段")
	for i := 0; i < s.NumField(); i++ {
		fmt.Println("for循环获取字段名称:", s.Field(i).Name)
		fmt.Println("for循环获取字段类型:", s.Field(i).Type)
		fmt.Println("for循环获取字段标签:", s.Field(i).Tag)
		fmt.Println("for循环获取字段标签中的json值:", s.Field(i).Tag.Get("json"))
		fmt.Println("for循环获取字段标签中的form值:", s.Field(i).Tag.Get("form"))
	}
	DeBug()
}

// 打印结构体的方法
func PrintStructMethod(v interface{}) {
	DeBug()
	// Method() 和 Field类似
	// Method() 是以ASCALL码排序的,与方法定义前后无关

	s := reflect.TypeOf(v)
	if s.Kind() != reflect.Struct && s.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体")
		return
	}

	// MethodByName()获取结构体的某个方法  --  与 FieldByName() 类似
	fmt.Println("使用MethodByName()获取结构体的某个方法")
	if method1, ok := s.MethodByName("GetInfo"); !ok {
		fmt.Println("结构体中没有GetInfo方法")
		return
	} else {
		fmt.Println("方法名称:", method1.Name)
		fmt.Println("方法类型:", method1.Type)
	}
	DeBug()

	fmt.Println("结构体的方法数量:", s.NumMethod())
	// 使用遍历结合 Method()获取结构体的所有方法
	for i := 0; i < s.NumMethod(); i++ {
		fmt.Println("方法名称:", s.Method(i).Name)
		fmt.Println("方法类型:", s.Method(i).Type)
	}
	DeBug()

	// Call()调用结构体的方法

	// 无参
	Value := reflect.ValueOf(v)
	result_getinfo := Value.MethodByName("GetInfo").Call(nil)
	fmt.Println("GetInfo方法的返回值:", result_getinfo)
	DeBug()

	// 有参
	params := []reflect.Value{
		reflect.ValueOf("测试带参Call"),
		reflect.ValueOf(18),
		reflect.ValueOf(100),
	}
	// 调用SetInfo方法
	Value.MethodByName("SetInfo").Call(params)
	fmt.Println(Value.MethodByName("GetInfo").Call(nil))
	DeBug()
}

// 利用反射修改结构体的字段值
func ChangeStruct(v interface{}) {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		fmt.Println("传入的参数不是结构体指针")
		return
	} else if reflect.TypeOf(v).Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体指针")
		return
	}
	reflect.ValueOf(v).Elem().FieldByName("Name").SetString("测试修改结构体name")
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
	PrintStuctField(stu1)
	PrintStructMethod(&stu1)

	fmt.Printf("%#v\n", stu1)
	DeBug()

	DeBug()
	ChangeStruct(&stu1)
	fmt.Printf("%#v\n", stu1)
	DeBug()

}
