package main

import (
    "math/big"
    "fmt"
)

func main() {
    target := big.NewInt(1)
    target.Lsh(target, uint(10))
    fmt.Printf("Target is: %b\n", target)
}
