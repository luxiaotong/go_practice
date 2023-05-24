package main

import "log"

type doctor struct {
	next Department
}

func (r *doctor) execute(p *patient) {
	p.doctorDone = true
	log.Printf("patient(%s) doctor done\n", p.name)
	if r.next != nil {
		r.next.execute(p)
	}
}

func (r *doctor) setNext(d Department) {
	r.next = d
}
