package pca

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	sdb "datassets.cn/service/db"
	"datassets.cn/share/db/postgresql"
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

var lock sync.Mutex

var db *postgresql.Impl

func LoadDB() *postgresql.Impl {
	lock.Lock()
	defer lock.Unlock()
	if db != nil {
		return db
	}
	var err error
	db, err = sdb.SetDB("127.0.0.1", 5432, "dbuser", "test", "userpanel")
	if err != nil {
		log.Panic("load db panic", err)
	}
	return db
}

func ExportJSON() {
	db = LoadDB()
	provinces, cities, err := sdb.CachedProvinceCities(db.DB)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	countyMap, err := sdb.CachedCounties(db.DB)
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
