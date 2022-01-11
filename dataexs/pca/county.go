package pca

import (
	"database/sql"
)

type County struct {
	Id       int    `json:"_id"`
	CityId   string `json:"city_id"`
	CountyId string `json:"county_id"`
	Name     string `json:"name"`
}

func FetchCounties(db *sql.DB) map[string][]*County {
	q, err := db.Query("SELECT _id, name, county_id, city_id FROM county ORDER BY city_id, county_id")
	if err != nil {
		// log.Error("FetchAll return error: %v", err)
		return nil
	}
	ret := make(map[string][]*County)
	for q.Next() {
		var p County
		err := q.Scan(&p.Id, &p.Name, &p.CountyId, &p.CityId)
		if err != nil {
			// log.Error("FetchAll scan return error: %v", err)
			return nil
		}
		cities := ret[p.CityId]
		cities = append(cities, &p)
		ret[p.CityId] = cities
	}
	return ret
}
