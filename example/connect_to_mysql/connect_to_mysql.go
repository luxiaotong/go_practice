package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db1, err := sql.Open("mysql", "root:@(127.0.0.1:6666)/csse_covid_19_daily_reports")
	fmt.Println(db1)
	err = db1.Ping()
	fmt.Println(err)

	db2, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/csse_covid_19_daily_reports")
	fmt.Println(db1)
	err = db2.Ping()
	fmt.Println(err)

	db, err := gorm.Open("mysql", "root:@(127.0.0.1)/?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("mysql", "root:@(127.0.0.1:6377)/csse_covid_19_daily_reports?charset=utf8&parseTime=True&loc=Local")
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
	// rows, err = db.Raw("SHOW TABLES FROM csse_covid_19_daily_reports;").Rows()
	rows, err = db.Raw("SELECT TABLE_NAME, TABLE_TYPE FROM information_schema.tables where table_schema='csse_covid_19_daily_reports';").Rows()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Tables:")
	var tbName, tbType string
	for rows.Next() {
		rows.Scan(&tbName, &tbType)
		fmt.Println(tbName, tbType)
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
