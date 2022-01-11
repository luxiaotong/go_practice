package pca

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type provCode struct {
	Name     string     `json:"name"`
	Code     string     `json:"code"`
	Children []cityCode `json:"children"`
}

type cityCode struct {
	Name     string     `json:"name"`
	Code     string     `json:"code"`
	Children []areaCode `json:"children"`
}

type areaCode struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func ExportJSON() {
	db = LoadDB()
	provinces, cities, err := ProvinceCities(db.DB)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	countyMap, err := Counties(db.DB)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	var pCode []provCode
	for _, p := range provinces {
		var cCode []cityCode
		for _, c := range cities[p.ProvinceId] {
			var aCode []areaCode
			for _, a := range countyMap[c.CityId] {
				// fmt.Println(a.CountyId, a.Name)
				aCode = append(aCode, areaCode{Name: a.Name, Code: a.CountyId})
			}
			cCode = append(cCode, cityCode{Name: c.Name, Code: c.CityId, Children: aCode})
		}

		pCode = append(pCode, provCode{Name: p.Name, Code: p.ProvinceId, Children: cCode})
	}
	// fmt.Printf("result: %v\n", pCode)

	b, err := json.Marshal(pCode)
	if err != nil {
		fmt.Println("error:", err)
	}
	// fmt.Println(string(b))

	if err := ioutil.WriteFile("pca-code.json", b, 0644); err != nil {
		fmt.Println("error:", err)
	}
}
