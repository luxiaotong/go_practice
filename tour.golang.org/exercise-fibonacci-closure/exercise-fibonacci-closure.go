package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    i := 0
    j := 1
    return func() int {
        tmp := i + j
        i = j
        j = tmp
        return tmp
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
