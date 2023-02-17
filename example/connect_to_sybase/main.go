package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/thda/tds"
)

func main() {
	cnxStr := "tds://tester:guest1234@139.9.119.21:55000/testdb?charset=utf8"
	db, err := sql.Open("tds", cnxStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Panicf("sybase ping err: %v", err)
	}
	fmt.Println(db)

	rows, err := db.Query("select name, type from sysobjects where user_name(uid)=user_name()")
	if err != nil {
		log.Panicf("sybase query tables err: %v", err)
	}
	for rows.Next() {
		var tablename string
		var typ string
		if err := rows.Scan(&tablename, &typ); err != nil {
			log.Panicf("sybase scan table err: %v", err)
		}
		log.Printf("table name: %v, type: %v", tablename, typ)
	}

	rows, err = db.Query("select index_col(object_name(i.id), i.indid, 1)"+
		"from sysindexes as i where object_name(i.id)=? and status&2048=2048", "test_type")
	if err != nil {
		log.Panicf("sybase query primary key err: %v", err)
	}
	for rows.Next() {
		var primarykey string
		if err := rows.Scan(&primarykey); err != nil {
			log.Panicf("sybase scan primary err: %v", err)
		}
		log.Printf("primary key: %v", primarykey)
	}

	rows, err = db.Query("select c.name, c.type, c.usertype, c.length, c.prec, c.scale "+
		"from sysobjects as o, syscolumns as c "+
		"where o.name=? and o.id=c.id", "test_type")
	if err != nil {
		log.Panicf("sybase query fields err: %v", err)
	}
	for rows.Next() {
		var colName string
		var colType, colExt, colLen int64
		var colPrec, colScale sql.NullInt64
		if err := rows.Scan(&colName, &colType, &colExt, &colLen, &colPrec, &colScale); err != nil {
			log.Panicf("sybase scan fields err: %v", err)
		}
		log.Printf("fields: %v, (%v, %v), %v, (%v, %v)", colName, colType, colExt, colLen, colPrec, colScale)
		// t := ""
		// switch colType {
		// case 0:
		// 	t = "CHAR"
		// case 1:
		// 	t = "SMALLINT"
		// case 2, 258:
		// 	t = "INTEGER"
		// case 3:
		// 	t = "FLOAT"
		// case 4:
		// 	t = "SMALLFLOAT"
		// case 5:
		// 	t = "DECIMAL"
		// case 6, 262:
		// 	t = "SERIAL"
		// case 7:
		// 	t = "DATE"
		// case 8:
		// 	t = "MONEY"
		// case 9:
		// 	t = "NULL"
		// case 10:
		// 	t = "DATETIME"
		// case 11:
		// 	t = "BYTE"
		// case 12:
		// 	t = "TEXT"
		// case 13:
		// 	t = "VARCHAR"
		// case 14:
		// 	t = "INTERVAL"
		// case 15:
		// 	t = "NCHAR"
		// case 16:
		// 	t = "NVARCHAR"
		// case 17:
		// 	t = "INT8"
		// case 18:
		// 	t = "SERIAL8"
		// case 19:
		// 	t = "SET"
		// case 20:
		// 	t = "MULTISET"
		// case 21:
		// 	t = "LIST"
		// case 23:
		// 	t = "COLLECTION"
		// case 40:
		// 	t = "LVARCHAR"
		// case 41:
		// 	t = "BLOB"
		// 	if colExt == 5 {
		// 		t = "BOOLEAN"
		// 	}
		// case 43:
		// 	t = "LVARCHAR"
		// case 45:
		// 	t = "BOOLEAN"
		// case 52:
		// 	t = "BIGINT"
		// case 53:
		// 	t = "BIGSERIAL"
		// }
		// log.Printf("fields: %v, %v, %v", colName, t, colLen)
	}
}
