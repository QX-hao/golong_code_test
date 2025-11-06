package FileRead

import (
	"fmt"
	"io"
	"os"
	"strings"
)
func DeBug() {
	fmt.Println(strings.Repeat("-", 20))
}

// 方法一：
// 		1.只读方式打开 file,_ := os.Open(filename)
// 		2.读取文件 file.Read()
// 		3.关闭文件流 file.Close()
func FileOpen1(filename string) (*os.File, error) {
	// 打开文件 -- Open() 只读模式
	if file, err := os.Open(filename); err != nil {
		fmt.Println("打开文件失败", err)
		DeBug()
		return nil, err
	} else {
		// defer file.Close() // 必须关闭文件流
		fmt.Println(file) // &{0xc0001046c8}
		DeBug()
		return file, nil
	}
}

// 读取文件内容
func FileRead1(file *os.File) {
	// buf用来存储读取的文件内容 -- 扩容机制
	var buf []byte
	// 只存128字节  -- 临时存放
	var tempBuff = make([]byte, 128)

	for {
		if n, err := file.Read(tempBuff); err != nil {
			if err == io.EOF {
				// io.EOF 表示文件读取结束
				fmt.Println("读取文件结束:", err)
				// fmt.Printf("读取文件内容:%s\n",buf)
				fmt.Println("读取文件内容1:", string(buf))
				DeBug()
				break
			} else {
				fmt.Println("读取文件失败", err)
				DeBug()
				return
			}
		} else {
			fmt.Printf("读取文件内容:%d字节\n", n)
			// fmt.Printf("读取文件内容:%s\n",buf)
			// fmt.Println("读取文件内容:",string(tempBuff[:n]))

			// 最后一次读取k可能不足128字节，需要特殊处理，只append实际读取的字节
			// 否则会append多余的0字节
			// buf = append(buf,tempBuff...) // 这样写会有问题
			buf = append(buf, tempBuff[:n]...)
			DeBug()
		}
	}
}

func CloseFile(file *os.File) {
	file.Close()
}