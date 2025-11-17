package FileRead

import (
	"bufio"
	"fmt"
	"os"
	"io"
)

// 方法二:bufio 读取文件
// 		1.只读方式打开 file,_ := os.Open(filename)
// 		2.创建bufio读取器 reader := bufio.NewReader(file)
// 		3.读取文件内容 reader.ReadString('\n')
// 		4.关闭文件流 file.Close()

func FileRead2(file *os.File ) {
	var FileStr string
	reader := bufio.NewReader(file)
	for{

		// 按照\n（换行符）来分割读取 -- 一次读取一行
		if str,err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完成2")
				fmt.Println("文件内容:",FileStr)
				break
				
			}
				fmt.Println("读取文件失败:",err)
				return
		} else {
			// fmt.Println("读取文件内容:",str)
			FileStr += str
		}
	}
	DeBug()

}
