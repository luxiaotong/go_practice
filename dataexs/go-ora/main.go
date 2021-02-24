package main

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora"
)

func main() {
	fmt.Println("go-ora")
	openString := "oracle://test:datassets@139.9.119.21:51521/xe"
	db, err := sql.Open("oracle", openString)
	if err != nil {
		fmt.Printf("open error: %v\n", err)
		return
	}
	defer db.Close()
	for i := uint32(0); i < 400; i++ {
		list(db, i, 50)
	}
}

func list(db *sql.DB, pageNo, pageSize uint32) {
	q := "SELECT earthquake_id FROM earthquake order by earthquake_id OFFSET :1 ROWS FETCH NEXT :2 ROWS ONLY"
	fmt.Println("query: ", q, pageNo)
	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Printf("prepare error: %v\n", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(pageNo*pageSize, pageSize)
	if err != nil {
		fmt.Printf("query error: %v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var val uint32
		err = rows.Scan(&val)
		if err != nil {
			fmt.Printf("scan error: %v\n", err)
			return
		}
		fmt.Println("id: ", val)
	}
}
