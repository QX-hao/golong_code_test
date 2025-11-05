package FileRead

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 注意：ioutil已经被废弃，不建议使用
// 建议使用os.Open()和io/ioutil.ReadAll()来替代

// 方法三:ioutil 读取文件
// 		封装了打开和关闭文件的方法
// 		ioutil.ReadFile(filename)

func FileRead31(filename string) {
	// ReadFile()不是以流来读取文件内容，而是一次性读取到内存中(适合小文件不适合大文件)
	// 从 Go 1.16 开始，ioutil.ReadFile()此函数仅调用 os.ReadFile()
	if bytes, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println("读取文件失败31:", err)
		return
	} else {
		fmt.Println("读取文件成功31:", string(bytes))
	}
	DeBug()
}

func FileRead32(filename string) {
	// ReadFile()不是以流来读取文件内容，而是一次性读取到内存中
	// ReadFile 读取由 filename 指定的文件并返回其内容。
	// 成功调用返回 err == nil，而不是 err == EOF。
	// 因为 ReadFile 函数 读取整个文件时，不会将读取过程中遇到的文件结束符 (EOF) 视为错误。 需上报。
	if bytes, err := os.ReadFile(filename); err != nil {
		fmt.Println("读取文件失败32:", err)
		return
	} else {
		fmt.Println("读取文件成功32:", string(bytes))
	}
	DeBug()
}
