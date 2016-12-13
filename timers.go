package main

import (
    "fmt"
    "time"
)

func main() {
    time1 := time.NewTimer(time.Second)
    <- time1.C
    fmt.Println("time 1 expired")

    time2 := time.NewTimer(time.Second * 2)
    go func() {
        <- time2.C
        fmt.Println("time 2 expired")
    } ()

    stop := time2.Stop()
    if stop {
        fmt.Println("time 2 stopped")
    }
}
