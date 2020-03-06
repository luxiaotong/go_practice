package main

import (
    "fmt"
    "sort"
)

func main() {
    ints := []int{8, 3, 5, 5, 6, 2, 1, 4}
    fmt.Println(sort.IntsAreSorted(ints))

    sort.Ints(ints)
    fmt.Println(ints)

    fmt.Println(sort.IntsAreSorted(ints))
}
