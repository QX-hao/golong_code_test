package removedir

import (
	"os"
	"fmt"
)

// Removedir2 删除目录 可以删除多级目录
func Removedir2(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		fmt.Println("删除目录失败", err)
		return
	}
}