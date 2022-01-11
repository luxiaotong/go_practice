package pca

import (
	"errors"
)

//DBInterface 提供数据库抽象接口
type DBInterface interface {
	TableList() ([]*TableInfo, error)
	TableDetail(tbName string) (*TableFieldsInfo, error)

	// DataList 是用于获取指定表的数据列表
	//  param: string tbName, 表名
	//  param: []string fields 返回的字段列表，nil表示全部
	//  param: []string orders 排序列表，nil表示不排序
	//  param: int pageNo, 页码
	//  param: int pageSize， 页大小
	//  return: []map[string]interface{}, 数据列表
	//  return: int, 数据总行数
	//  return: error, 错误信息
	DataList(tbName string, fields, orders []string, pageNo, pageSize uint32) ([]map[string]interface{}, uint32, error)
	LastModified() (string, error)
	Close()

	// ToPostgresType 转换字段类型为PostgreSQL的字段类型
	//  param: []*v1.FieldInfo fields 要转换的字段类型列表
	//  return: []*v1.FieldInfo 转换好的类型
	ToPostgresType(fields []*FieldInfo) []*FieldInfo
}

var (
	ErrDBNotConnect = errors.New("db is not connected")
)
