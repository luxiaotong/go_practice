package main
import "fmt"

func main() {
    a := make([]int, 5)
    printSlice(a)

    b := make([]int, 0, 5)
    printSlice(b)

    c := b[:2]
    printSlice(c)

    d := c[2:5]
    printSlice(d)
}

func printSlice(s[]int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s);
}
