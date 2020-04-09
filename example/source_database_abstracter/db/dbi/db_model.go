package dbi

//TableInfo 是查询GetTableList返回中的结构
type TableInfo struct {
	Name   string
	TypeOf string //"table"|"view"
	// PrimaryKey string
	// Size    int
}

//ColumnInfo 是查询GetColumnList返回的结构
type ColumnInfo struct {
	Name   string
	TypeOf string
	Length int
}

//TableDetail 是查询GetTableDetail返回的结构
//List: Column列表
//PrimaryKey: 主键
//Size: 表行数
type TableDetail struct {
	ColumnList []ColumnInfo
	PrimaryKey string
	Size       int
}

//DataList 是查询GetDataList返回的结构
type DataList struct {
	List  []map[string]interface{}
	Total int //或者返回totalPage int
}

/*
TableList:
[
	{
		"name": "123123",
		"type": "table"|"view",
		"primary": "",
		"size": 100
	},
	{...}
]

ColumnList:
{
	list: [
		{
			"name":
			"type":
			"length":
		},
		{...}
	],
	primary: "",
	tableSize: 100,
}
*/
