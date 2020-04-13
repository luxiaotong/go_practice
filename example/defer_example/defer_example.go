package main

import "fmt"

func main() {
	defer d(f())
}

var v int = 1

func f() int {
	fmt.Println("In Defer")
	v = 2
	return v
}

func d(i int) {
	fmt.Println("Defer: ", i)
}
