package mysql

import (
	"fmt"

	"github.com/luxiaotong/go_practice/example/source_database_abstracter/db/dbi"
)

//Impl 是DBInterface抽象的MySQL具体实现
type Impl struct {
	//gormDB *gorm.DB
	DBName string
}

//Connect 是用于连接Mysql
func (m *Impl) Connect() {
	fmt.Println("MySQL Connect")
	// TODO
	// m.GormDB = db
}

//GetTableList 是用于获取指定数据库的Table列表
//	参数
//
//	返回
//		[]TableInfo
func (m *Impl) GetTableList() []dbi.TableInfo {
	fmt.Println("MySQL GetTableList from ", m.DBName)
	return []dbi.TableInfo{}
}

//GetColumnList 是用于获取指定表的字段列表
//	参数
//		tbname
//	返回
//		ColumnList
func (m *Impl) GetColumnList(tbname string) dbi.ColumnList {
	fmt.Println("MySQL GetColumnList from ", tbname)
	return dbi.ColumnList{}
}

//GetDataList 是用于获取指定表的数据列表
//	参数
//		tbname
//		pageNo
//		pageSize
//	返回
//		DataList
func (m *Impl) GetDataList(tbname string, pageNo, pageSize int) dbi.DataList {
	//SELECT ALL，不需要选择字段
	fmt.Println("MySQL GetDataList from ", tbname)
	return dbi.DataList{}
}

//GetLSN 是用于获取指定数据库的Log Sequence Number
func (m *Impl) GetLSN() string {
	fmt.Println("MySQL GetLSN")
	return "TODO"
}

//Close 是用于关闭数据库连接
// func (m *Impl) Close() {
// 	fmt.Println("MySQL Close")
// }
