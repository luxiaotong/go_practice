package main

import (
	"log"

	"github.com/luxiaotong/go_practice/example/source_database_abstracter/db"
)

func main() {
	d, err := db.GetDB("mysql", "127.0.0.1", 6379, "root", "", "csse_covid_19_daily_reports")
	if err != nil {
		log.Fatal("GetDB Failed: ", err)
	}
	// defer d.DBObj.Close()

	d.GetTableList()
	d.GetColumnList("TBNAME")
}
