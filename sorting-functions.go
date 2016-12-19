package main

import (
    "fmt"
    "sort"
)

type ByLength []string

func (bl ByLength) Len() int {
    return len(bl)
}

func (bl ByLength) Swap(i, j int) {
    bl[i], bl[j] = bl[j], bl[i]
}

func (bl ByLength) Less(i, j int) bool {
    return len(bl[i]) < len(bl[j])
}

func main() {
    str := []string{"aaaaaa", "aaa", "bbbbbbbb"}
    sort.Sort(ByLength(str))
    fmt.Println(str)
}
