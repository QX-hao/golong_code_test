package main

import (
	"log"
)


// Ldate         = 1 << iota     // 日期：2009/01/23
// Ltime                         // 时间：01:23:23
// Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
// Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
// Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
// LUTC                          // 使用UTC时间
// LstdFlags     = Ldate | Ltime // 标准logger的初始值
func main()  {
	// SetFlags默认值为 LstdFlags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("设置完格式后的日志信息")
}