package dbi

import (
	"fmt"
)

//MySQLAPI 是DBInterface抽象的MySQL具体实现（基于gorm?)
type MySQLAPI struct {
	//GormDB *gorm.DB
}

//Connect 是用于连接Mysql
func (m *MySQLAPI) Connect() {
	fmt.Println("MySQL Connect")
}

//GetDatabaseList 是用于获取数据库列表
func (m *MySQLAPI) GetDatabaseList() {
	fmt.Println("MySQL GetDatabaseList")
}

//GetTableList 是用于获取指定数据库的Table列表
func (m *MySQLAPI) GetTableList(dbname string) {
	fmt.Println("MySQL GetTableList from ", dbname)
}

//GetViewList 是用于获取指定数据库的视图列表
func (m *MySQLAPI) GetViewList(dbname string) {
	fmt.Println("MySQL GetViewList from ", dbname)
}

//GetColumnList 是用于获取指定数据库的字段列表
func (m *MySQLAPI) GetColumnList(tbname string) {
	fmt.Println("MySQL GetColumnList from ", tbname)
}

//GetLSN 是用于获取指定数据库的Log Sequence Number
func (m *MySQLAPI) GetLSN() {
	fmt.Println("MySQL GetLSN")
}

//Close 是用于关闭数据库连接
func (m *MySQLAPI) Close() {
	fmt.Println("MySQL Close")
}
