package createdir

import (
	"os"
	"fmt"
)

// Createdie2 创建多级目录
func Createdie2(manydir string, perm os.FileMode) {
	err := os.MkdirAll(manydir, perm)
	if err != nil {
		fmt.Println("创建多级目录失败", err)
		return
	}
}
