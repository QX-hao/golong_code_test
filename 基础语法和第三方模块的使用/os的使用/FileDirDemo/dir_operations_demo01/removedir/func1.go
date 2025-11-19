package removedir

import (
	"os"
	"fmt"
)

// Removedir1 删除目录 
// Remove不仅可以删除目录 还可以删除文件
func Removedir1(dir string) {
	err := os.Remove(dir)
	if err != nil {
		fmt.Println("删除目录失败", err)
		return
	}
}
