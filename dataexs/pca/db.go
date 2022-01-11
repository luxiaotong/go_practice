package pca

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/pkg/errors"
)

func SetDB(host string, port int, user, pass, dbname string) (*Impl, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		user, pass, host, port, dbname)
	fmt.Printf("open sample db %s\n", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "open sample db error")
	}
	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping error")
	}
	return &Impl{DB: db}, nil
}

var lock sync.Mutex

var db *Impl

func LoadDB() *Impl {
	lock.Lock()
	defer lock.Unlock()
	if db != nil {
		return db
	}
	var err error
	db, err = SetDB("127.0.0.1", 5432, "dbuser", "test", "userpanel")
	if err != nil {
		log.Panic("load db panic", err)
	}
	return db
}

func ProvinceCities(db *sql.DB) ([]*Province, map[string][]*City, error) {
	provinces := FetchProvinces(db)
	if len(provinces) == 0 {
		return nil, nil, errors.New("fetch province error")
	}
	cities := FetchCities(db)
	if len(cities) == 0 {
		return nil, nil, errors.New("fetch city error")
	}

	return provinces, cities, nil
}

func Cities(db *sql.DB) (map[string]*City, error) {
	_, cities, err := ProvinceCities(db)
	if err != nil {
		return nil, err
	}
	cityMap := make(map[string]*City)
	for _, cc := range cities {
		for _, city := range cc {
			cityMap[city.CityId] = city
		}
	}
	return cityMap, nil
}

func Counties(db *sql.DB) (map[string][]*County, error) {
	counties := FetchCounties(db)
	if len(counties) == 0 {
		return nil, errors.New("fetch county error")
	}
	return counties, nil
}
