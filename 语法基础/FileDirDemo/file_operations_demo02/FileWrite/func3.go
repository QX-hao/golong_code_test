package FileWrite

import (
	"os"
	"time"
	"io/ioutil"
)

// 写入文件(方法3) - ioutil 写入文件
// 步骤说明：
// 1、准备数据 str := "hello golang"
// 2、写入文件 err := ioutil.WriteFile("c:/test.txt", []byte(str), 0666)

// 一样的废弃了，变成os.WriteFile
// WriteFile会清空文件内容后写入新内容
func FileWrite3(filename string,perm os.FileMode) {
	ioutil.WriteFile(filename, []byte(time.Now().Format("2006-01-02 15:04:05    ")+"info:"+"ioutil.WriteFile写入    "+"hello golang\r\n"), perm)
}
