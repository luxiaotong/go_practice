package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tealeg/xlsx"
)

func main() {
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
