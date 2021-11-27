package main

import (
    "fmt"
)

func main() {
    arr1 := [4]byte {1,2,3,4}
    arr2 := arr1[:]
    arr2[0] = 5
    fmt.Printf("arr1: %v\n", arr1)
    fmt.Printf("arr1: %T\n", arr1)
    fmt.Printf("arr1: %p\n", arr1)
    fmt.Printf("arr2: %v\n", arr2)
    fmt.Printf("arr2: %T\n", arr2)
    fmt.Printf("arr2: %p\n", arr2)
}
