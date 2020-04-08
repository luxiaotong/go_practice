package dbi

//DBApi 是用于数据库访问统一API
type DBApi struct {
	DBType string
	DBObj  DBInterface
}

//Open 用于建立数据库连接
//	参数
//		sourceType: "mysql"|"pgsql"
//	返回值
//		*DBApi
func Open(sourceType string) *DBApi {
	var m DBInterface
	var d *DBApi

	if sourceType == "mysql" {
		m = &MySQLAPI{}
		m.Connect()
		d = &DBApi{sourceType, m}
	}
	return d
}
