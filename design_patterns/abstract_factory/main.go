package main

import "fmt"

func main() {
	adidasFactory := buildFactory("adidas")
	shirt := adidasFactory.makeShirt(175)
	shoes := adidasFactory.makeShoes(42)
	fmt.Printf("shirt logo: %v, size: %v\n", shirt.getLogo(), shirt.getSize())
	fmt.Printf("shoes logo: %v, size: %v\n", shoes.getLogo(), shoes.getSize())

	nikeFactory := buildFactory("nike")
	shirt = nikeFactory.makeShirt(175)
	shoes = nikeFactory.makeShoes(42)
	fmt.Printf("shirt logo: %v, size: %v\n", shirt.getLogo(), shirt.getSize())
	fmt.Printf("shoes logo: %v, size: %v\n", shoes.getLogo(), shoes.getSize())
}
