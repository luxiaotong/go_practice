package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	_ "github.com/lib/pq"
)

var (
	node *snowflake.Node
	once sync.Once
)

func Gen() int64 {
	once.Do(func() {
		var err error
		node, err = snowflake.NewNode(int64(1))
		if err != nil {
			log.Fatal("error in snowflake.NewNode", err)
		}
	})
	return node.Generate().Int64()
}

func main() {
	dsn := "user=auth password=authpass host=139.9.119.21 port=5432 dbname=schema sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("./schema2.csv")
	if err != nil {
		log.Fatalln("open schema failed")
	}
	defer f.Close()
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		id := Gen()
		q := "INSERT INTO schemas (id, label_cn, label_en, comment_cn, create_time, update_time) VALUES ($1, $2, $3, $4, $5, $5)"
		if _, err := db.Exec(q, id, record[0], record[1], record[2], time.Now().Unix()); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d, %s, %s, %s\n", id, record[0], record[1], record[2])
	}
}
