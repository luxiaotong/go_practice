package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("/Users/luxiaotong/code/go_practice/dataexs/pdf_replace/12121212.qrdoc.pdf")
	if err != nil {
		log.Fatal("read pdf error:", err)
	}
	// log.Printf("b: %v", string(b))
	has := strings.Contains(string(b), "2021")
	log.Printf("has 2021: %v", has)
	has = strings.Contains(string(b), "2022")
	log.Printf("has 2022: %v", has)
	s := bytes.Replace(b, []byte("2021"), []byte("2022"), -1)
	has = strings.Contains(string(s), "2022")
	log.Printf("has 2022: %v", has)
	err = ioutil.WriteFile("/Users/luxiaotong/code/go_practice/dataexs/pdf_replace/replace.qrdoc.pdf", []byte(s), 0644)
	if err != nil {
		log.Fatal("write pdf error:", err)
	}
}
