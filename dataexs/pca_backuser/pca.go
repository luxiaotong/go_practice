package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type provCode struct {
	Name     string      `json:"name"`
	Code     string      `json:"code"`
	Children []*cityCode `json:"children"`
}

type cityCode struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Area string `json:"area_code"`
}

func main() {
	dsn := "user=auth password=authpass host=139.9.119.21 port=5432 dbname=backuser sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	q := "select name, city_id, province_id,area_code from city;"
	cRows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer cRows.Close()

	cMap := make(map[string][]*cityCode)
	for cRows.Next() {
		var name, cityID, provinceID, areaCode string
		_ = cRows.Scan(&name, &cityID, &provinceID, &areaCode)
		fmt.Printf("name: %s, cid: %s, pid: %s, area: %s\n", name, cityID, provinceID, areaCode)
		cMap[provinceID] = append(cMap[provinceID], &cityCode{
			Name: name,
			Code: cityID,
			Area: areaCode,
		})
	}
	// fmt.Printf("cmap: %v", cMap)

	q = "select name,province_id from province;"
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pCode []*provCode
	for rows.Next() {
		var name, provinceID string
		_ = rows.Scan(&name, &provinceID)
		fmt.Printf("id: %s, name: %s\n", provinceID, name)
		pCode = append(pCode, &provCode{
			Name:     name,
			Code:     provinceID,
			Children: cMap[provinceID],
		})
	}
	// fmt.Printf("pCode: %#v", pCode)

	b, err := json.Marshal(pCode)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	if err := ioutil.WriteFile("pca-code-back.json", b, 0644); err != nil {
		fmt.Println("error:", err)
	}
}
