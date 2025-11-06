package createdir

import (
	"os"
	"fmt"
)

// Createdir1 创建目录testdir1
func Createdir1(dirname string,perm os.FileMode)  {
	err := os.Mkdir(dirname,perm)
	if err != nil {
		fmt.Println("创建目录失败",err)
		return
	}
}