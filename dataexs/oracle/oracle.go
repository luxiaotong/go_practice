package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-oci8"
)

func main() {
	fmt.Println("hello oracle")
	openString := "test/datassets@//139.9.119.21:51521/xe"
	db, err := sql.Open("oci8", openString)
	if err != nil {
		fmt.Printf("open error: %v", err)
		return
	}
	defer db.Close()

	var rows *sql.Rows
	rows, err = db.QueryContext(context.Background(), "select * from T_STU")
	if err != nil {
		fmt.Println("QueryContext error is not nil:", err)
		return
	}
	var result [][]string
	for rows.Next() {
		var id, name string
		_ = rows.Scan(&id, &name)

		result = append(result, []string{id, name})
	}
	fmt.Printf("result: %v", result)
}
