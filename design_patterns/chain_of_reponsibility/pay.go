package main

import "log"

type pay struct {
	next Department
}

func (r *pay) execute(p *patient) {
	p.payDone = true
	log.Printf("patient(%s) pay done\n", p.name)
	if r.next != nil {
		r.next.execute(p)
	}
}

func (r *pay) setNext(d Department) {
	r.next = d
}
