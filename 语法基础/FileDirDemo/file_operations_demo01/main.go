package main

import (
	"file_operations_demo01/FileRead"
)

// 文件的操作 -- 读取

func main() {
	// 法一:
	// 打开文件 -- 只读
	file, _ := FileRead.FileOpen1("config/config.json")
	// 操作文件
	FileRead.FileRead1(file)
	// 关闭文件流
	FileRead.CloseFile(file)

	FileRead.DeBug()
	// 法二:
	// 打开文件 -- 只读
	file, _ = FileRead.FileOpen1("config/config.json")
	// 操作文件
	FileRead.FileRead2(file)
	// 关闭文件流
	FileRead.CloseFile(file)

	FileRead.DeBug()
	// 法三:

	FileRead.FileRead31("config/config.json")
	FileRead.FileRead32("config/config.json")



}
