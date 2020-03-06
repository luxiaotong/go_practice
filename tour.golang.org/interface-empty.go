package main
import (
    "fmt"
    "math"
)

type I interface{}

func main() {
    var i I
    describe(i)

    i = float64(math.Pi)
    describe(i)

    i = "Hello"
    describe(i)
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
