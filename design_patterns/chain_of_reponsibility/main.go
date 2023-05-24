package main

import "log"

func main() {
	log.Println("chain")

	med := &medicine{}

	pay := &pay{}
	pay.setNext(med)

	doc := &doctor{}
	doc.setNext(pay)

	rec := &receipt{}
	rec.setNext(doc)

	rec.execute(&patient{
		name: "test patient",
	})
}
