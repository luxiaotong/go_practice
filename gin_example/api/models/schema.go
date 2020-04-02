package models

import (
	"fmt"
	"log"

	"github.com/luxiaotong/go_practice/gin_example/api/database"
)

//DBName 返回结果
type DBName struct {
	Name string `json:"db_name"`
}

//DBNames 包含一组数据库名字
var DBNames []DBName

//Databases list all databases
func Databases() []DBName {

	// Show Databases
	rows, err := database.Eloquent.Raw("SHOW DATABASES;").Rows()
	if err != nil {
		log.Fatal(err)
	}

	//rows.Scan(&DBNames)

	fmt.Println("All Databases:")
	var dbName string
	for rows.Next() {
		rows.Scan(&dbName)
		fmt.Println(dbName)
		DBNames = append(DBNames, DBName{dbName})
	}

	fmt.Printf("dbname: %v\n", DBNames)
	return DBNames
}

//Tables list all tables
func Tables(dbname string) []string {
	// Show Tables
	rows, err := database.Eloquent.Raw("SHOW TABLES FROM " + dbname).Rows()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Tables:")
	var tbName string
	var TableNames []string
	for rows.Next() {
		rows.Scan(&tbName)
		fmt.Println(tbName)
		TableNames = append(TableNames, tbName)
	}

	return TableNames
}
