package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/luxiaotong/go_practice/pprof_example/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/luxiaotong"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
