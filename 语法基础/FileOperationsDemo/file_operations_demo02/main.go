package main

import (
	"file_operations_demo02/FileWrite"
	"os"
	"time"
)

// 文件操作 -- 写入


func main()  {
	// 写入文件(方法1) - 直接写入文件
	file := FileWrite.FileWrite1("log/"+time.Now().Format("2006-01-02")+".log",os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	FileWrite.FileClose(file)

	// 写入文件(方法2) - bufio 写入文件
	file = FileWrite.FileWrite2("log/"+time.Now().Format("2006-01-02")+".log",os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	FileWrite.FileClose(file)

	// 写入文件(方法3) - ioutil 写入文件
	FileWrite.FileWrite3("log/"+time.Now().Format("2006-01-02")+".log1", 0666)
}