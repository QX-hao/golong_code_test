package main

import (
	"log"
	"os"
	"time"
)

// 使用New()来创造自己的logger对象
// func New(out io.Writer, prefix string, flag int) *Logger

func main()  {
	os.Mkdir("log",0755)
	logfile,logfileerr := os.OpenFile("log/"+time.Now().Format("2006-01-02")+".log",os.O_CREATE | os.O_WRONLY | os.O_APPEND ,0666) 
	if logfileerr != nil {
		log.Fatal(logfileerr)
		return
	}
	logger :=log.New(logfile, "new:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("使用New()来创造自己的logger对象")

	defer logfile.Close()

}
