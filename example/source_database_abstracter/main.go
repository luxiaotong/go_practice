package main

import (
	"fmt"
	"log"

	"github.com/luxiaotong/go_practice/example/source_database_abstracter/db"
	"github.com/luxiaotong/go_practice/example/source_database_abstracter/db/dbi"
)

func main() {
	d, err := db.GetDB("mysql", "127.0.0.1", 6379, "root", "", "csse_covid_19_daily_reports")
	if err != nil {
		log.Fatal("GetDB Failed: ", err)
	}
	// defer d.DBObj.Close()

	m, ok := d.(dbi.DBInterface)
	// m, ok := d.(mysql.Impl)
	if ok {
		fmt.Println("succ: ", m)
	} else {
		fmt.Println("fail: ", ok)
	}

	d.TableList()
	d.TableDetail("TBNAME")
}
