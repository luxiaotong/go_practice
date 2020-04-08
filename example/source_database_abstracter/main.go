package main

import "github.com/luxiaotong/go_practice/example/source_database_abstracter/dbi"

func main() {
	d := dbi.Open("mysql")
	d.DBObj.GetDatabaseList()
	d.DBObj.GetTableList("DBNAME")
	d.DBObj.GetColumnList("TBNAME")
	defer d.DBObj.Close()
}
