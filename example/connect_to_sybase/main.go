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
		t := ""
		switch colType {
		case 48:
			t = "tinyint"
		case 52:
			t = "smallint"
		case 65:
			t = "usmallint"
		case 56:
			t = "int"
		case 38:
			switch colExt {
			case 5:
				t = "tinyint"
			case 6:
				t = "smallint"
			case 7:
				t = "int"
			case 43:
				t = "bigint"
			default:
				t = "intn"
			}
		case 66, 68:
			t = "uint"
			if colExt == 46 {
				t = "ubigint"
			}
		case 191:
			t = "bigint"
		case 67:
			t = "ubigint"
		case 262:
			t = "serial"
		case 62:
			t = "float"
		case 109:
			switch colExt {
			case 8:
				t = "double"
			case 23:
				t = "real"
			default:
				t = "float"
			}
		case 59:
			t = "real"
		case 63, 108:
			t = "numeric"
		case 55, 106:
			t = "decimal"
		case 60, 110:
			t = "money"
		case 122:
			t = "smallmoney"
		case 50:
			t = "bit"
		case 61:
			t = "datetime"
		case 111:
			switch colExt {
			case 12:
				t = "datetime"
			case 22:
				t = "smalldatetime"
			default:
				t = "datetime"
			}
		case 58:
			t = "smalldatetime"
		case 37:
			t = "timestamp"
			switch colExt {
			case 3:
				t = "binary"
			case 4:
				t = "varbinary"
			case 80:
				t = "timestamp"
			}
		case 49, 123:
			t = "date"
		case 51, 147:
			t = "time"
		case 39:
			t = "varchar"
			if colExt == 1 {
				t = "char"
			}
		case 47:
			t = "char"
		case 45:
			t = "binary"
		}
		log.Printf("fields: %v, %v", colName, t)
	}
}
