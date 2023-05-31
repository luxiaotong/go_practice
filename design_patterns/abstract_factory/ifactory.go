package main

type ifactory interface {
	makeShoes(int32) ishoes
	makeShirt(int32) ishirt
}
