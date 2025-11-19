package main

import (
	"fmt"
	"strings"
)

/*
	接口声明：
		type 接口名 interface{
			方法1(参数) 返回值
			方法2(参数) 返回值
		}
*/

// 接口是一个规范，定义了一组方法，实现了该接口的结构体必须实现这些方法
type Usber interface {
	start(string,string) int
	stop()
}

// 接口中有方法，可以用结构体或者自定义类型实现这个接口

type Phone struct {
	Name string
}

func (p Phone) start(brand string, model string) int {
	fmt.Println("品牌：", brand)
	fmt.Println("型号：", model)
	fmt.Println("手机名称：", p.Name)
	fmt.Println("启动成功")
	return 0
}

func (p Phone) stop() {
	fmt.Println("手机名称：", p.Name)
	fmt.Println("关闭成功")
}

type Camera struct {

}

func (c Camera) start(brand string, model string) int {
	fmt.Println("品牌：", brand)
	fmt.Println("型号：", model)
	fmt.Println("相机启动")
	fmt.Println("启动成功")
	return 0
}

func (c Camera) stop() {
	fmt.Println("相机关闭")
	fmt.Println("关闭成功")
}

func main() {
	p := Phone{
		Name: "IPhone",
	}
	p.start("Apple", "IPhone16")
	p.stop()

	var p1 Usber
	p1 = p
	// 接口类型变量可以指向实现了该接口的结构体实例
	// p1.Name // 接口类型变量不能直接访问结构体的字段
	p1.start("HuaWei", "HuaWeiP50")
	p1.stop()

	fmt.Println(strings.Repeat("-", 20))

	var c Camera
	// c := Camera{}

	var c1 Usber = c
	// c1 = c
	c1.start("Nikon", "NikonD850")
	c1.stop()
}