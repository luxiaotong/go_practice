package main

import (
    "fmt"
)

func foo(arr []int) func() {
    g := func() {
        fmt.Printf("foo arr: %v\n", arr)
    }
    arr[1] = 5
    return g
}

func foo1(arr []int) {
    ch := make(chan int, 1)
    go func(chan int) {
        fmt.Printf("foo1 arr: %v\n", arr)
        ch <- 1
    }(ch)
    <-ch
}

func main() {
    arr := []int{1,2,3,4}
    f := foo(arr)
    f()
    foo1(arr)
}
