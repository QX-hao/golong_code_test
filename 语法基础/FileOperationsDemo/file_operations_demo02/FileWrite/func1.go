package FileWrite

import (
	"os"
	"fmt"
	"strings"
	"time"
)

// 写入文件(方法1) - 直接写入
// 步骤说明：
// 1、打开文件 file,err := os.OpenFile("c:/test.txt",os.O_CREATE|os.O_RDWR, 0666)
// 2、写入文件
//     file.Write([]byte(str))        // 写入字节切片数据
//     file.WriteString("直接写入的字符串数据") // 直接写入字符串数据
// 3、关闭文件流 file.close()

// OpwenFile()的flag参数 说明：
// os.O_CREATE：如果文件不存在，则创建文件
// os.O_RDWR：以读写模式打开文件
// os.O_WRONLY：以只写模式打开文件
// os.O_RDONLY：以只读模式打开文件
// os.O_APPEND：追加写入模式，新写入的内容会追加到文件末尾
// os.O_TRUNC：如果文件存在，清空文件内容
// 0666：文件权限(八进制)，读写权限都允许所有用户 -- 和linux的权限一样  4 2 1

func DeBug()  {
	fmt.Println(strings.Repeat("-", 20))
}

func FileWrite1(filename string,flag int,perm os.FileMode) *os.File  {
	if file,err := os.OpenFile(filename,flag, perm); err != nil {
		fmt.Println("打开文件失败",err)
		DeBug()
		return nil
	} else {
		file.Write([]byte(time.Now().Format("2006-01-02 15:04:05    ")+"info:"+"os.File.Write写入    "+"直接写入的字节切片数据\n"))
		file.WriteString(time.Now().Format("2006-01-02 15:04:05    ")+"info:"+"os.File.WriteString写入    "+"直接写入的字符串数据\n")
		DeBug()
		return file
	}
}

func FileClose(file *os.File)  {
	file.Close()
	DeBug()
}