package main

import (
	"file_operations_demo03/copyfile"
	"fmt"
	"strings"
)

func DeBug()  {
	fmt.Println(strings.Repeat("-",20))
}

func main()  {

	// 方法一:
	// 利用os.ReadFile和os.WriteFile方法复制文件
	// 其实就是 ioutil.ReadFile和ioutil.WriteFile方法复制文件
	copyfile.CopyFile1("testdir/test.txt","testdir/test_copy1.txt")
	DeBug()

	// 方法二:
	// 利用os.Open os.Read(流处理)和os.OpenFile 方法复制文件
	copyfile.CopyFile2("testdir/test.txt","testdir/test_copy2.txt")
	DeBug()

	


}