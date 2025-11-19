package connectmysql

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CREATE TABLE `users` (
//   `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
//   `username` varchar(30) NOT NULL COMMENT '账号',
//   `password` varchar(100) NOT NULL COMMENT '密码',
//   `createtime` int(10) NOT NULL DEFAULT 0 COMMENT '创建时间',
//    PRIMARY KEY (`id`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8

//定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
//在这里User类型可以代表mysql users表
type User struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	CreateTime int `gorm:"column:createtime"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}


func ConnectMysqlUsers() *gorm.DB {
	//连接mysql数据库
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := "3306"
	Dbname := "gormdatabase"


	//通过前面的数据库参数，拼接MYSQL DSN， 其实就是数据库连接串（数据源名称）
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",username,password,host,port,Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败：%v",err)
	}
	return db
}
// 插入数据
func InsertUser(db *gorm.DB) {
	//定义一个用户，并初始化数据
	u := User{
		Username: "xiaopan",
		Password: "123456",
		CreateTime: int(time.Now().Unix()),
	}
	//插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('tizi365','123456','1540824823')
	
	// Create返回一个DB结构体指针，DB结构体指针包含了很多方法，比如Error、RowsAffected、Statement等，
	// 我们可以通过DB结构体指针调用这些方法，来获取插入操作的结果
	// 	type DB struct {
	// 	*Config
	// 	Error        error
	// 	RowsAffected int64
	// 	Statement    *Statement
	// 	clone        int
	// }
	if db.Create(&u).Error != nil {
		log.Println("插入用户失败")
	} else {
		log.Println("插入用户成功")
	}
}

// *gorm.DB.Debug() 开启调试模式，打印出自动生成的SQL语句
// 查询并返回第一条数据
// 定义需要保存数据的struct变量
func SelectUser(db *gorm.DB)  {
	var u User
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'xiaopan') LIMIT 1
	// First()
	result := db.Where("username = ? ", "xiaopan").First(&u)

	// err == Error 直接比较错误对象与指定的错误对象，也就是指针的对比。
	// errors.Is(err, Error) 通过不断解包err，如果匹配错误类型，则返回 true。
	if errors.Is(result.Error,gorm.ErrRecordNotFound){
		log.Println("查询用户失败，用户不存在")
		return
	}
	// 打印信息
	log.Printf("查询用户成功，用户信息：%v\n",u)
	log.Printf("username:%s,password:%s,createtime:%d\n",u.Username,u.Password,u.CreateTime)
}

// 更新数据
func UpdateUser(db *gorm.DB)  {
	var u User
	//更新
	//自动生成Sql: UPDATE `users` SET `password` = '654321'  WHERE (username = 'xiaopan')
	// Update()
	db.Model(&u).Where("username = ?","xiaopan").Update("password","654321")
	if db.Error != nil {
		log.Println("更新用户失败")
	} else {
		log.Println("更新用户成功")
	}
}

func DeleteUser(db *gorm.DB)  {
	var u User
	//删除
	//自动生成Sql: DELETE FROM `users`  WHERE (username = 'xiaopan')
	// Delete()
	db.Where("username = ?","xiaopan1").Delete(&u)

	if db.Error != nil {
		log.Println("删除用户失败")
	} else {
		log.Println("删除用户成功")
	}
}

// 关闭数据库连接
func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库连接失败：%v", err)
	}
	sqlDB.Close()
}
