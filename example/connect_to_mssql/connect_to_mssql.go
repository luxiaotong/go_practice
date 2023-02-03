package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// dsn := "sqlserver://test:C#2sZwp3@139.9.119.21:51433"
	query := url.Values{}
	query.Add("database", "testdb")
	username := "test"
	password := "C#2sZwp3"
	hostname := "139.9.119.21"
	port := 51433
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s:%d", hostname, port),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		log.Panicf("mssql open err: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Panicf("mssql ping err: %v", err)
	}
	// fmt.Println(db)

	rows, err := db.Query("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES")
	if err != nil {
		log.Panicf("mssql query tables err: %v", err)
	}
	for rows.Next() {
		var tablename string
		if err := rows.Scan(&tablename); err != nil {
			log.Panicf("mssql scan table err: %v", err)
		}
		log.Printf("table name: %v", tablename)
	}

	rows, err = db.Query("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE WHERE OBJECTPROPERTY(OBJECT_ID(CONSTRAINT_SCHEMA + '.' + QUOTENAME(CONSTRAINT_NAME)), 'IsPrimaryKey') = 1 AND TABLE_NAME = @p1", "test_type")
	if err != nil {
		log.Panicf("mssql query primary key err: %v", err)
	}
	for rows.Next() {
		var primarykey string
		if err := rows.Scan(&primarykey); err != nil {
			log.Panicf("mssql scan primary err: %v", err)
		}
		log.Printf("primary key: %v", primarykey)
	}

	rows, err = db.Query("SELECT COLUMN_NAME, DATA_TYPE, CHARACTER_MAXIMUM_LENGTH, NUMERIC_PRECISION, NUMERIC_SCALE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = @p1", "test_type")
	if err != nil {
		log.Panicf("mssql query fields err: %v", err)
	}
	for rows.Next() {
		var colName, colType string
		var charLen, numPrecision, numScale sql.NullInt64
		if err := rows.Scan(&colName, &colType, &charLen, &numPrecision, &numScale); err != nil {
			log.Panicf("mssql scan fields err: %v", err)
		}
		log.Printf("fields: %v, %v, %v, %v, %v", colName, colType, charLen, numPrecision, numScale)
	}
}
