package main

import (
	"dir_operations_demo01/createdir"
	"dir_operations_demo01/removedir"
	"strings"
	"fmt"
)

func DeBug()  {
	fmt.Println(strings.Repeat("=",20))
}

func main()  {
	createdir.Createdir1("testdir1",0666)
	DeBug()
	createdir.Createdie2("testdir2/testdir3",0666)

	DeBug()
	removedir.Removedir1("testdir1")

	DeBug()
	removedir.Removedir2("testdir2")
}