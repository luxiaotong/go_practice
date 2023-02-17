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

	rows, err = db.Query("select c.colname, c.coltype, c.collength, c.extended_id from systables as t, syscolumns as c where t.tabname=? and c.tabid=t.tabid;", "test_type")
	if err != nil {
		log.Panicf("mssql query fields err: %v", err)
	}
	for rows.Next() {
		var colName string
		var colType, colLen, colExt int64
		if err := rows.Scan(&colName, &colType, &colLen, &colExt); err != nil {
			log.Panicf("mssql scan fields err: %v", err)
		}
		log.Printf("fields: %v, (%v, %v), %v", colName, colType, colExt, colLen)
		t := ""
		switch colType {
		case 0:
			t = "CHAR"
		case 1:
			t = "SMALLINT"
		case 2, 258:
			t = "INTEGER"
		case 3:
			t = "FLOAT"
		case 4:
			t = "SMALLFLOAT"
		case 5:
			t = "DECIMAL"
		case 6, 262:
			t = "SERIAL"
		case 7:
			t = "DATE"
		case 8:
			t = "MONEY"
		case 9:
			t = "NULL"
		case 10:
			t = "DATETIME"
		case 11:
			t = "BYTE"
		case 12:
			t = "TEXT"
		case 13:
			t = "VARCHAR"
		case 14:
			t = "INTERVAL"
		case 15:
			t = "NCHAR"
		case 16:
			t = "NVARCHAR"
		case 17:
			t = "INT8"
		case 18:
			t = "SERIAL8"
		case 19:
			t = "SET"
		case 20:
			t = "MULTISET"
		case 21:
			t = "LIST"
		case 23:
			t = "COLLECTION"
		case 40:
			t = "LVARCHAR"
		case 41:
			t = "BLOB"
			if colExt == 5 {
				t = "BOOLEAN"
			}
		case 43:
			t = "LVARCHAR"
		case 45:
			t = "BOOLEAN"
		case 52:
			t = "BIGINT"
		case 53:
			t = "BIGSERIAL"
		}
		log.Printf("fields: %v, %v, %v", colName, t, colLen)
	}
}
