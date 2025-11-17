package main

import (
	"gormdemo01/connectmysql"
)

func main() {
	// 连接数据库
	db := connectmysql.ConnectMysqlUsers()
	// 插入数据
	connectmysql.InsertUser(db)
	// 查询数据
	connectmysql.SelectUser(db)
	// 更新数据
	connectmysql.UpdateUser(db)
	// 删除数据
	connectmysql.DeleteUser(db)
	// 关闭数据库连接
	connectmysql.CloseDB(db)
}
