package dbi

//DBInterface 提供数据库抽象接口
type DBInterface interface {
	TableList() ([]TableInfo, error)
	TableDetail(tbname string) (TableDetail, error)
	DataList(tbname string, pageNo, pageSize int) (DataList, error)
	LastModified() (string, error)
	// Connect()
	// Close()
}

//循环引用问题
// ColumnList命名
