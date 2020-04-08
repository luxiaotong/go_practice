package dbi

//DBInterface 提供数据库抽象接口
type DBInterface interface {
	Connect()
	GetDatabaseList()
	// TODO 分页
	GetTableList(dbname string)
	GetViewList(dbname string) //View和Table是否单独提供？
	GetColumnList(tbname string)
	GetLSN() //获取数据库的Log Sequence Number
	Close()
}
