package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
	"github.com/tealeg/xlsx"
)

func main() {
	toDB()
	// toJson()
}

var pdb *sql.DB

func loadDB() {
	dbType := "postgres"
	var err error
	dsn := "user=auth password=authpass host=139.9.119.21 port=5432 dbname=unittest sslmode=disable"
	pdb, err = sql.Open(dbType, dsn)
	if err != nil {
		panic(err)
	}
}

func toDB() {
	loadDB()
	name := "./国民经济行业分类信息.xlsx"
	xlFile, err := xlsx.OpenFile(name)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	sheet := xlFile.Sheets[0]
	var currA, currB, currC, currD string
	var nameA, nameB, nameC, nameD string
	for _, row := range sheet.Rows[1:] {
		for i, cell := range row.Cells[:4] {
			id := cell.String()
			name := row.Cells[4].String()
			if id != "" {
				switch i {
				case 0:
					currA = id
					nameA = name
					q := "INSERT INTO industry_class_a(name, class_a_id) VALUES ($1, $2)"
					if _, err := pdb.Exec(q, nameA, currA); err != nil {
						panic(err)
					}
				case 1:
					currB = id
					nameB = name
					q := "INSERT INTO industry_class_b(name, parent, class_b_id, class_a_id) VALUES ($1, $2, $3, $4)"
					if _, err := pdb.Exec(q, nameB, nameA, currB, currA); err != nil {
						panic(err)
					}
				case 2:
					currC = id
					nameC = name
					q := "INSERT INTO industry_class_c(name, parent, class_c_id, class_b_id) VALUES ($1, $2, $3, $4)"
					if _, err := pdb.Exec(q, nameC, nameB, currC, currB); err != nil {
						panic(err)
					}
				case 3:
					currD = id
					nameD = name
					q := "INSERT INTO industry_class_d(name, parent, class_d_id, class_c_id) VALUES ($1, $2, $3, $4)"
					if _, err := pdb.Exec(q, nameD, nameC, currD, currC); err != nil {
						panic(err)
					}
				}
				fmt.Printf("currA:%s, currB:%s, currC:%s, currD:%s\n", currA, currB, currC, currD)
				fmt.Printf("nameA:%s, nameB:%s, nameC:%s, nameD:%s\n", nameA, nameB, nameC, nameD)
			}
		}
	}
}

// nolint:deadcode,unused
func toJson() {
	name := "./国民经济行业分类信息.xlsx"
	xlFile, err := xlsx.OpenFile(name)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	sheet := xlFile.Sheets[0]

	nameMap := make(map[string]string)
	arrA := make([]string, 0)
	mapB := make(map[string][]string)
	mapC := make(map[string][]string)
	mapD := make(map[string][]string)
	var currA, currB, currC, currD, key string
	for _, row := range sheet.Rows[1:] {
		for i, cell := range row.Cells[:4] {
			id := cell.String()
			if id != "" {
				switch i {
				case 0:
					currA = id
					arrA = append(arrA, currA)
					key = currA
				case 1:
					currB = id
					mapB[currA] = append(mapB[currA], currB)
					key = currA + "_" + currB
				case 2:
					currC = id
					mapC[currB] = append(mapC[currB], currC)
					key = currA + "_" + currB + "_" + currC
				case 3:
					currD = id
					mapD[currC] = append(mapD[currC], currD)
					key = currA + "_" + currB + "_" + currC + "_" + currD

				}
				nameMap[key] = row.Cells[4].String()
			}
		}
	}

	type Industry struct {
		Name string      `json:"name"`
		Sub  []*Industry `json:"sub"`
	}
	var industryA []*Industry
	for _, idA := range arrA {
		var industryB []*Industry
		for _, idB := range mapB[idA] {
			var industryC []*Industry
			for _, idC := range mapC[idB] {
				var industryD []*Industry
				for _, idD := range mapD[idC] {
					key = idA + "_" + idB + "_" + idC + "_" + idD
					fmt.Println(idA, idB, idC, idD, nameMap[key])
					industryD = append(industryD, &Industry{
						Name: nameMap[key],
						Sub:  make([]*Industry, 0),
					})
				}
				key = idA + "_" + idB + "_" + idC
				industryC = append(industryC, &Industry{
					Name: nameMap[key],
					Sub:  industryD,
				})
			}
			key = idA + "_" + idB
			industryB = append(industryB, &Industry{
				Name: nameMap[key],
				Sub:  industryC,
			})
		}
		key = idA
		industryA = append(industryA, &Industry{
			Name: nameMap[key],
			Sub:  industryB,
		})
	}
	industry, _ := json.Marshal(industryA)
	_ = ioutil.WriteFile("industry.json", industry, 0644)
}
