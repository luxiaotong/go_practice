package main

import (
    "fmt"
    //"time"
)

func main() {
    for i := 1; i < 40; i ++ {
        if ( i % 3 == 0 || i % 10 == 3 ) {
            fmt.Print("过")
        } else {
            fmt.Print(i)
        }
        fmt.Print("  ")
    }
}
