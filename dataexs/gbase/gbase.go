package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "139.9.119.21:55258",
		DBName:               "test",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Panicf("open sql error: %v", err)
	}
	rows, err := db.Query("select area_code, area_name from area limit 0, 10")
	if err != nil {
		log.Panicf("query db error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var code int32
		var name string
		if err := rows.Scan(&code, &name); err != nil {
			log.Panicf("rows scan error: %v", err)
		}
		log.Printf("area code: %v, name: %v", code, name)
	}
}
