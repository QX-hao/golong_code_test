package main

import (
	"log"
)
// 配置日志前缀
// func Prefix() string
// func SetPrefix(prefix string)

func main()  {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("没有设置前缀的日志信息")
	log.SetPrefix("INFO: ")
	log.Println("设置完日志前缀的信息")
}