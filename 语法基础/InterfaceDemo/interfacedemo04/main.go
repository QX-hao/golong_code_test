package main

import (
	"fmt"
	"strings"
)

// 结构体  值接收者 和  指针接收者

type Usber interface {
	Start()
	Stop()
}

type Phone struct {
	Brand string
	Price int
}

// 值接收者  调用方法时，使用值类型或指针类型都可以
func (p Phone) Start() {
	fmt.Println(p.Brand, "手机开始工作了")
}

func (p Phone) Stop() {
	fmt.Println(p.Brand, "手机停止工作了")
}

// 指针接收者  调用方法时，必须使用指针类型
// func (p *Phone) Start() {
// 	fmt.Println(p.Brand, "手机开始工作了")
// }

// func (p *Phone) Stop() {
// 	fmt.Println(p.Brand, "手机停止工作了")
// }

func main() {
	fmt.Println(strings.Repeat("-", 20))

	var p = Phone{
		Brand: "Apple",
		Price: 1000,
	}
	var p2 Usber = p
	p2.Start()
	p2.Stop()

	fmt.Println(strings.Repeat("-", 20))


	// 指针接收者  调用方法时，必须使用指针类型
	var p3 = &Phone{
		Brand: "HuaWei",
		Price: 9999,
	}
	var p4 Usber = p3
	p4.Start()
	p4.Stop()

	fmt.Println(strings.Repeat("-", 20))

}
