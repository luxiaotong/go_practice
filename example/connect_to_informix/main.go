package main

import (
	"database/sql"
	"log"

	_ "github.com/ibmdb/go_ibm_db"
)

func main() {
	dsn := "HOSTNAME=139.9.119.21;DATABASE=testdb;PORT=59089;UID=test;PWD=datassets"
	db, err := sql.Open("go_ibm_db", dsn)
	if err != nil {
		log.Panicf("informix open err: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Panicf("informix ping err: %v", err)
	}
	// fmt.Println(db)

	rows, err := db.Query("select tabname from systables where tabid > 100;")
	if err != nil {
		log.Panicf("informix query tables err: %v", err)
	}
	for rows.Next() {
		var tablename string
		if err := rows.Scan(&tablename); err != nil {
			log.Panicf("informix scan table err: %v", err)
		}
		log.Printf("table name: %v", tablename)
	}

	rows, err = db.Query("select c.colname from sysconstraints as co, sysindexes as i, syscolumns as c, systables as t"+
		" where t.tabname=? and i.tabid=t.tabid and c.tabid=t.tabid"+
		" and co.constrtype = 'P' and co.idxname=i.idxname and c.colno = i.part1;", "test_type")
	if err != nil {
		log.Panicf("informix query primary key err: %v", err)
	}
	for rows.Next() {
		var primarykey string
		if err := rows.Scan(&primarykey); err != nil {
			log.Panicf("informix scan primary err: %v", err)
		}
		log.Printf("primary key: %v", primarykey)
	}

	rows, err = db.Query("select c.colname, coltype, collength from systables as t, syscolumns as c where t.tabname=? and c.tabid=t.tabid;", "test_type")
	if err != nil {
		log.Panicf("mssql query fields err: %v", err)
	}
	for rows.Next() {
		var colName, colType string
		var colLen int64
		if err := rows.Scan(&colName, &colType, &colLen); err != nil {
			log.Panicf("mssql scan fields err: %v", err)
		}
		log.Printf("fields: %v, %v, %v", colName, colType, colLen)
	}
}
