package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:@(127.0.0.1)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	// Show Databases
	rows, err := db.Raw("SHOW DATABASES;").Rows()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Databases:")
	var dbName string
	for rows.Next() {
		rows.Scan(&dbName)
		fmt.Println(dbName)
	}

	fmt.Println()
	// Show Tables
	rows, err = db.Raw("SHOW TABLES FROM csse_covid_19_daily_reports;").Rows()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Tables:")
	var tbName string
	for rows.Next() {
		rows.Scan(&tbName)
		fmt.Println(tbName)
	}

	fmt.Println()
	// Show Columns
	schema := "csse_covid_19_daily_reports"
	table := "03-30-2020"
	query := fmt.Sprintf("SELECT COLUMN_NAME FROM information_schema.columns WHERE table_schema='%s' AND table_name='%s'", schema, table)
	rows, err = db.Raw(query).Rows()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Columns:")
	var colName string
	for rows.Next() {
		rows.Scan(&colName)
		fmt.Println(colName)
	}

	defer db.Close()
}
