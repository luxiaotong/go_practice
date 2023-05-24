package main

import "log"

type medicine struct {
	next Department
}

func (r *medicine) execute(p *patient) {
	p.medicineDone = true
	log.Printf("patient(%s) medicine done\n", p.name)
	if r.next != nil {
		r.next.execute(p)
	}
}

func (r *medicine) setNext(d Department) {
	r.next = d
}
