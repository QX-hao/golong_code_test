package main

import (
	"log"
)

func main()  {
	// 日志输出到控制台
	log.Println("这是一条日志信息")

	str := "this is a test of log"
	log.Printf("%s",str)

	// Fatal系列函数会在日志输出后调用os.Exit(1)，导致程序退出
	log.Fatalf("触发fatal错误")
	log.Panicln("触发panic错误")
}