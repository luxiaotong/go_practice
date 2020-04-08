package dbi

//TableInfo 是查询GetTableList返回中的结构
type TableInfo struct {
	name   string
	typeOf string //"table"|"view"
	// primary string
	// size    int
}

//ColumnInfo 是查询GetColumnList返回的结构
type ColumnInfo struct {
	name   string
	typeof string
	length int
}

//ColumnList 是查询GetColumnList返回的结构
type ColumnList struct {
	list    []ColumnInfo
	primary string
	size    int
}

//DataList 是查询GetDataList返回的结构
type DataList struct {
	list  []map[string]interface{}
	total int //或者返回totalPage int
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
