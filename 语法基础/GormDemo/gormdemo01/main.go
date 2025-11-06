package main

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

type CarInfo struct {
	Id int
	Brand string
	Model string
	Configuration string
	Range_km float64
}

// Tablename 自定义表名
func (CarInfo) Tablename() string {
	return "car_info"
}

func main()  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/doncar_top?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("数据库连接失败", err)
	}
	log.Println("数据库连接成功")
	defer func () {
		sqlDB, _ := db.DB()
		db.Find(&CarInfo{})

		log.Println("查询到的car_info表数据:", a)
		
		sqlDB.Close()
	}()
}