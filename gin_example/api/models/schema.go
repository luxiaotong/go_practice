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

//Databases list all databases
func Databases() []DBName {

	// Show Databases
	rows, err := database.Eloquent.Raw("SHOW DATABASES;").Rows()
	if err != nil {
		log.Fatal(err)
	}

	var DBNames []DBName
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
	var tbNames []string
	for rows.Next() {
		rows.Scan(&tbName)
		fmt.Println(tbName)
		tbNames = append(tbNames, tbName)
	}

	return tbNames
}

//Columns list all columns
func Columns(dbname string, tablename string) []string {
	// Show Columns
	query := fmt.Sprintf("SELECT COLUMN_NAME FROM information_schema.columns "+
		"WHERE table_schema='%s' AND table_name='%s'", dbname, tablename)
	rows, err := database.Eloquent.Raw(query).Rows()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Columns:")
	var colName string
	var colNames []string
	for rows.Next() {
		rows.Scan(&colName)
		fmt.Println(colName)
		colNames = append(colNames, colName)
	}

	return colNames
}
