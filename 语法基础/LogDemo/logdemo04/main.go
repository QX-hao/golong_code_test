package main

import (
	"log"
	"os"
	"time"
)


// 配置日志输出位置
// func SetOutput(w io.Writer)

// 可以写到init中
var logfile *os.File

func init() {
	os.Mkdir("log",0755)
	logfile,logfileerr := os.OpenFile("log/"+time.Now().Format("2006-01-02")+".log",os.O_CREATE | os.O_WRONLY | os.O_APPEND ,0666) 
	if logfileerr != nil {
		log.Fatalf("打开日志文件失败:%v",logfileerr)
		return
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logfile)
	log.SetPrefix("Prefix:")
}

func main()  {
	defer logfile.Close()
	log.Println("这是一条日志信息")
}
