package pca

import (
	"database/sql"
)

const Country = "000"

type Province struct {
	Id         int    `json:"_id"`
	ProvinceId string `json:"province_id"`
	Name       string `json:"name"`
}

func FetchProvinces(db *sql.DB) []*Province {
	q, err := db.Query("SELECT _id, name, province_id FROM province ORDER BY province_id")
	if err != nil {
		// log.Error("FetchAll return error: %v", err)
		return nil
	}
	var ret []*Province
	for q.Next() {
		p := Province{}
		err := q.Scan(&p.Id, &p.Name, &p.ProvinceId)
		if err != nil {
			// log.Error("FetchAll scan return error: %v", err)
			return nil
		}
		ret = append(ret, &p)
	}
	return ret
}
