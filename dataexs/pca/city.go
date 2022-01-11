package pca

import (
	"database/sql"
)

type City struct {
	Id         int    `json:"_id"`
	ProvinceId string `json:"province_id"`
	CityId     string `json:"city_id"`
	Name       string `json:"name"`
	AreaCode   string `json:"area_code,omitempty"`
}

func FetchCities(db *sql.DB) map[string][]*City {
	q, err := db.Query("SELECT _id, name, city_id, province_id, area_code FROM city ORDER BY province_id, city_id")
	if err != nil {
		// log.Error("FetchAll return error: %v", err)
		return nil
	}
	ret := make(map[string][]*City)
	for q.Next() {
		p := City{}
		err := q.Scan(&p.Id, &p.Name, &p.CityId, &p.ProvinceId, &p.AreaCode)
		if err != nil {
			// log.Error("FetchAll scan return error: %v", err)
			return nil
		}
		cities := ret[p.ProvinceId]
		cities = append(cities, &p)
		ret[p.ProvinceId] = cities
	}
	return ret
}
