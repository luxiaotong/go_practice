package main

import (
	"crypto/md5"
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	SecretKey  = "userpanel@datassets"
	SecretAlgo = "md5"
)

var u = flag.String("u", "", "username")
var p = flag.String("p", "", "password")

var udb, pdb, odb, fdb *sql.DB

func LoadDB() {
	dbType := "postgres"
	var dsn string
	var err error
	dsn = "user=auth password=authpass host=139.9.119.21 port=5432 dbname=userpanel sslmode=disable"
	udb, err = sql.Open(dbType, dsn)
	if err != nil {
		panic(err)
	}
	dsn = "user=auth password=authpass host=139.9.119.21 port=5432 dbname=product sslmode=disable"
	pdb, err = sql.Open(dbType, dsn)
	if err != nil {
		panic(err)
	}
	dsn = "user=auth password=authpass host=139.9.119.21 port=5432 dbname=order sslmode=disable"
	odb, err = sql.Open(dbType, dsn)
	if err != nil {
		panic(err)
	}
	dsn = "user=auth password=authpass host=139.9.119.21 port=5432 dbname=favorite sslmode=disable"
	fdb, err = sql.Open(dbType, dsn)
	if err != nil {
		panic(err)
	}
}

func encrypt(password string) string {
	m := md5.New()
	_, _ = m.Write([]byte(password))
	_, _ = m.Write([]byte(string(SecretKey)))
	cipherStr := m.Sum(nil)
	return fmt.Sprintf("%s:%x", SecretAlgo, cipherStr)
}

func CheckPass(mobile, password string) (int64, int64) {
	password = encrypt(password)
	q := "SELECT id, firm_id FROM users WHERE mobile=$1 AND password=$2 LIMIT 1"
	row := udb.QueryRow(q, mobile, password)
	var userID, firmID int64
	if err := row.Scan(&userID, &firmID); err != nil {
		panic(err)
	}
	return userID, firmID
}

func GetPid(firmID int64) []int64 {
	q := "SELECT id FROM products WHERE firm_id = $1"
	rows, err := pdb.Query(q, firmID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var ret []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			panic(err)
		}
		ret = append(ret, id)
	}
	return ret
}

func GetOid(firmID int64) []int64 {
	q := "SELECT id FROM orders WHERE seller = $1 or buyer = $1"
	rows, err := odb.Query(q, firmID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var ret []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			panic(err)
		}
		ret = append(ret, id)
	}
	return ret
}

func GetFid(userID int64) []int64 {
	q := "SELECT id FROM favorites WHERE user_id = $1"
	rows, err := fdb.Query(q, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var ret []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			panic(err)
		}
		ret = append(ret, id)
	}
	return ret
}

func clear(userID, firmID int64) {
	if _, err := fdb.Exec("DELETE FROM favorites WHERE user_id = $1", userID); err != nil {
		panic(err)
	}
	if _, err := odb.Exec("DELETE FROM orders WHERE buyer = $1 OR seller = $1", firmID); err != nil {
		panic(err)
	}
	if _, err := pdb.Exec("DELETE FROM products WHERE firm_id = $1", firmID); err != nil {
		panic(err)
	}
	if _, err := udb.Exec("DELETE FROM users WHERE id = $1", userID); err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	if len(*u) == 0 || len(*p) == 0 {
		fmt.Println("invalid params")
		return
	}
	LoadDB()
	userID, firmID := CheckPass(*u, *p)
	fmt.Println("user id: ", userID, "firm id:", firmID)
	productID := GetPid(firmID)
	fmt.Println("product id: ", productID)
	orderID := GetOid(firmID)
	fmt.Println("order id: ", orderID)
	favoriteID := GetFid(userID)
	fmt.Println("favorite id: ", favoriteID)

	clear(userID, firmID)
	fmt.Println("Done!")
}
