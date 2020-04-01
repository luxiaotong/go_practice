package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=luxiaotong dbname=csse_covid_19_daily_reports sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// Show Databases
	rows, err := db.Raw("SELECT	datname FROM pg_database;").Rows()
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
	rows, err = db.Raw("SELECT table_name FROM information_schema.tables" +
		" WHERE table_catalog='csse_covid_19_daily_reports'" +
		" AND table_schema='public';").Rows()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Tables in csse_covid_19_daily_reports:")
	var tbName string
	for rows.Next() {
		rows.Scan(&tbName)
		fmt.Println(tbName)
	}

	fmt.Println()
	// Show Columns
	table := "03-30-2020"
	query := fmt.Sprintf("SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE TABLE_NAME = '%s';", table)
	rows, err = db.Raw(query).Rows()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Columns in 03-30-2020 csse_covid_19_daily_reports:")
	var colName string
	for rows.Next() {
		rows.Scan(&colName)
		fmt.Println(colName)
	}
	defer db.Close()
}
