package main

type Department interface {
	execute(p *patient)
	setNext(d Department)
}
