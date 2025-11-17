package copyfile

import (
	"os"
	"fmt"
)

// 方法一:
// 利用os.ReadFile和os.WriteFile方法复制文件
func CopyFile1(filename string,newFilename string)  {
	if file,err := os.ReadFile(filename); err != nil {
		fmt.Println("读取文件失败:",err)
		return
	}else {
		if err := os.WriteFile(newFilename,file,0666); err != nil {
			fmt.Println("写入文件失败:",err)
			return
		}
	}

}