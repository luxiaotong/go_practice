package dbi

//DBInterface 提供数据库抽象接口
type DBInterface interface {
	GetTableList() []TableInfo
	GetColumnList(tbname string) ColumnList
	GetDataList(tbname string, pageNo, pageSize int) DataList
	GetLSN() string
	// Connect()
	// Close()
}

//循环引用问题
