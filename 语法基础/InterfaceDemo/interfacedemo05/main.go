package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Setname(string)
	Getname() string
}


// Dog
type Dog struct {
	Name string
}

func (d *Dog) Setname(name string) {
	d.Name = name
}

func (d *Dog) Getname() string {
	return "小狗名字："+d.Name
}

// Cat
type Cat struct {
	Name string
}

func (c *Cat) Setname(name string) {
	c.Name = name
}

func (c *Cat) Getname() string {

	return "小猫名字："+c.Name
}

func main()  {
	fmt.Println(strings.Repeat("-",20))
	var dog_a Animal = &Dog{}
	dog_a.Setname("旺财")
	fmt.Println(dog_a.Getname())
	fmt.Println(strings.Repeat("-",20))

	var cat_a Animal = &Cat{}
	cat_a.Setname("咪咪")
	fmt.Println(cat_a.Getname())
	fmt.Println(strings.Repeat("-",20))
}