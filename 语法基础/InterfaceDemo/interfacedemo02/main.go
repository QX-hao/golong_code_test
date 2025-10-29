package main

import (
	"fmt"
	"strings"
)

type Usber interface {
	start()
	stop()
}

// 手机类型
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println("品牌：", p.Name)
	fmt.Println("启动成功")
}

func (p Phone) stop() {
	fmt.Println("品牌：", p.Name)
	fmt.Println("关闭成功")
}

// 相机类型
type Camera struct {

}

func (c Camera) start() {
	fmt.Println("相机启动")
	fmt.Println("启动成功")
}

func (c Camera) stop() {
	fmt.Println("相机关闭")
	fmt.Println("关闭成功")
}


// 电脑类型
type Computer struct {

}


// 这里类似多态，usb可以是Phone类型，也可以是Camera类型
// 类似  var usb Usber = Phone{} 或者 var usb Usber = Camera{}
func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}



func main() {
	var computer Computer
	var phone Phone
	phone.Name = "IPhone"
	var camera Camera

	computer.work(phone)
	fmt.Println(strings.Repeat("-", 20))

	// func() {
	// 	defer func() {
	// 		if err := recover(); err != nil {
	// 			fmt.Println("异常被捕获:", err)
	// 		}
	// 	}()
		
	// 	fmt.Println("尝试调用Camera...")
	// 	computer.work(camera)
	// }()

	computer.work(camera)

	// phone和camera都实现了Usber接口，所以可以作为参数传递给computer.work()
}