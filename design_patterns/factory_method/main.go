package main

import "fmt"

func main() {
	fmt.Println("factory method")
	g := getGun("ak47")
	fmt.Printf("gun name: %v, power: %v\n", g.getName(), g.getPower())
	g = getGun("musket")
	fmt.Printf("gun name: %v, power: %v\n", g.getName(), g.getPower())
}
