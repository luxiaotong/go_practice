package main

import "log"

type receipt struct {
	next Department
}

func (r *receipt) execute(p *patient) {
	p.receiptDone = true
	log.Printf("patient(%s) receipt done\n", p.name)
	if r.next != nil {
		r.next.execute(p)
	}
}

func (r *receipt) setNext(d Department) {
	r.next = d
}
