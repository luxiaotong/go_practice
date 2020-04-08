package db

import (
	"github.com/luxiaotong/go_practice/example/source_database_abstracter/db/dbi"
	"github.com/luxiaotong/go_practice/example/source_database_abstracter/db/mysql"
)

//DBObj 是用于数据库访问统一API
var DBObj dbi.DBInterface

//Open 用于建立数据库连接
//	参数
//		sourceType: "mysql"|"pgsql"
//	返回值
//		*DBInterface
// func Open(sourceType string) dbi.DBInterface {
// 	var m dbi.DBInterface

// 	if sourceType == "mysql" {
// 		m = &mysql.MySQLAPI{}
// 		m.Connect()
// 	}
// 	return m
// }

//GetDB 获取访问DB的实例
//	参数
//		sourceType: Source Database类型，"mysql"|"pgsql"
//		host: 数据库Host
//		port: 数据库端口号
//		user: 数据库用户名
//		passwd: 数据库密码
//		dbname: 数据库名
//	返回
//		DBInterface: 实现DBInterface的实例
//		error: 错误信息
func GetDB(sourceType, host string, port int, user, passwd, dbname string) (dbi.DBInterface, error) {
	//TODO:单例
	// bcOnce.Do(func() {
	// 	bc = &baiduChecker{}
	// 	config.RegisterListener("baiduCheck", bc)
	// })
	// return bc
	// return &mysql.Impl{dbname: dbname}, nil
	return &mysql.Impl{DBName: dbname}, nil
}
