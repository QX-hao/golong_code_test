package main

import (
	"errors"
	"fmt"
)

func readFile(fileName string) error {
	if fileName == "test.txt" {
		return nil
	} else {
		return errors.New("文件不存在")
	}
}

func fn() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("读取文件错误:", err)
		}
	}()

	// 调用 readFile 函数 判断是否读取文件成功
	// 如果读取文件失败 则 panic
	err := readFile("xxx.txt")
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("start")
	fn()
	fmt.Println("ending")
}
