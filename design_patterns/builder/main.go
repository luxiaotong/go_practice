package main

import "fmt"

func main() {
	normalBuilder := newNormalBuilder()
	iglooBuilder := newIglooBuilder()

	d := newDirector(normalBuilder)
	h := d.buildHouse()
	fmt.Printf("normal: %#v\n", h)

	d = newDirector(iglooBuilder)
	h = d.buildHouse()
	fmt.Printf("igloo: %#v\n", h)
}
