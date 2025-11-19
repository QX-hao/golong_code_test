package main

import (
	"time"
	"fmt"
)

func main()  {
	
	timeObj := time.Now()

	fmt.Println(timeObj) // 2025-10-25 19:10:19.9605226 +0800 CST m=+0.000000001

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("%d-%d-%d %d:%d:%d\n",year,month,day,hour,minute,second) // 2025-10-25 19:10:19

	// 格式刷输出
	// 其他语言 -- “Y-m-d H:i:s”
	// go不同--go 诞生时间 2006年1月2日 15（03）点04分05秒 
	// go -- “2006-01-02 15（03）:04:05” 
	// 15（03） 表示24（12）小时制

	fmt.Println("24小时制：",timeObj.Format("2006-01-02 15:04:05")) // 2025-10-25 07:10:19
	fmt.Println("12小时制：",timeObj.Format("2006-01-02 03:04:05")) // 2025-10-25 07:10:19

	// 获取时间戳
	// 时间转时间戳
	timestamp := timeObj.Unix()
	// 时间转时间戳（纳秒）
	timestampNano := timeObj.UnixNano()
	fmt.Println("时间戳：",timestamp) // 1700000000
	fmt.Println("纳秒时间戳：",timestampNano) // 1700000000960522600

	// 时间戳转时间
	timestamp01 := 1761391654
	timestampNano01 := 1761392204342767200
	// 参数一：时间戳
	// 参数二：纳秒时间戳
	timeObj01 := time.Unix(int64(timestamp01),0)
	timeObh02 := time.Unix(0,int64(timestampNano01))
	fmt.Printf("时间戳转时间：%v\n纳秒时间戳转时间：%v\n",timeObj01.Format("2006-01-02 15:04:05"),timeObh02.Format("2006-01-02 15:04:05")) // 2025-10-25 07:10:19

	// 同时给时间戳和纳秒时间戳，会自动合并将纳秒时间戳合并到时间戳中然后返回（相加）
	timeObj03 := time.Unix(int64(timestamp01),int64(timestampNano01))
	fmt.Printf("时间戳转时间：%v\n",timeObj03.Format("2006-01-02 15:04:05")) // 2025-10-25 07:10:19

	// 字符串转时间
	str := "2025-10-25 19:10:19"

	timeObj04, _ := time.ParseInLocation("2006-01-02 15:04:05",str,time.Local)
	fmt.Println("当前时区：",time.Local)
	fmt.Printf("字符串转时间：%v\n",timeObj04) // 字符串转时间：2025-10-25 19:10:19 +0800 CST
	fmt.Printf("转时间戳：%v\n",timeObj04.Unix()) // 转时间戳：1761390619

}