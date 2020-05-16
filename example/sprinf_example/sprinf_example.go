package main

import (
    "fmt"
)

func main() {
    a := 1
    fmt.Println(a)
    //前置补0
    fmt.Printf("%03d", a)
    fmt.Println("")
    n := 3
    fmt.Printf("%0*d\n", n, a)
    n = 5
    fmt.Printf("%0*d\n", n, a)
}
