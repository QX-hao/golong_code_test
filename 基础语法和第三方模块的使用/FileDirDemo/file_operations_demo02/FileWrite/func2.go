package FileWrite

import (
	"os"
	"fmt"
	"time"
	"bufio"
)

// 写入文件(方法2) - bufio 写入文件
// 步骤说明：
// 1、打开文件 file,err := os.OpenFile("c:/test.txt",os.O_CREATE|os.O_RDWR, 0666)
// 2、创建writer对象 writer := bufio.NewWriter(file)
// 3、将数据先写入缓存
//     writer.WriteString("你好golang\r\n")
// 4、将缓存中的内容写入文件
//     writer.Flush()
// 5、关闭文件流 file.close()

func FileWrite2(filename string,flag int,perm os.FileMode) *os.File  {
	if file,err := os.OpenFile(filename,flag, perm); err != nil {
		fmt.Println("打开文件失败",err)
		DeBug()
		return nil
	} else {
		writer := bufio.NewWriter(file)
		// 写入缓存
		writer.WriteString(time.Now().Format("2006-01-02 15:04:05    ")+"info:"+"bufio.Writer.WriteString写入    "+"你好golang\r\n")
		// 将缓存数据写入文件
		writer.Flush()
		DeBug()
		return file
	}
}