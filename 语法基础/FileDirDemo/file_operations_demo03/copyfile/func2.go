package copyfile

import (
	"os"
	"fmt"
	"io"
)

// 方法二:
// 利用os.Open os.Read(流处理)和os.OpenFile 方法复制文件


func CopyFile2(filename string,newFilename string)  {
	var buf_all []byte
	if file,err := os.Open(filename); err != nil{
		fmt.Println("打开原文件失败:",err)
		return
	} else {
		defer file.Close()
		var buf  = make([]byte,128)
		for {
			if n,readerr := file.Read(buf); readerr != nil {
				if readerr == io.EOF {
					fmt.Println("读取完成")
					break
				} else {
					fmt.Println("读取文件失败:",readerr)
					return
				}
			} else {
				buf_all = append(buf_all,buf[:n]...)
			}	
		}

		if newFile,newerr := os.OpenFile(newFilename,os.O_CREATE|os.O_WRONLY|os.O_TRUNC,066); newerr != nil {
			fmt.Println("创建备份文件失败",newerr)
			return
		} else {
			defer newFile.Close()
			newFile.Write(buf_all)
		}
	}
}