package dbi

//DBInterface 提供数据库抽象接口
type DBInterface interface {
	GetTableList() []TableInfo
	GetTableDetail(tbname string) TableDetail
	GetDataList(tbname string, pageNo, pageSize int) DataList
	GetLSN() string
	// Connect()
	// Close()
}

//循环引用问题
// ColumnList命名
