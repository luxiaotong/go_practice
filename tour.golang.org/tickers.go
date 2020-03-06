package main

import (
    "fmt"
    "time"
)

func main() {
    tick := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range tick.C {
            fmt.Println(t)
        }
    } ()

    time.Sleep(time.Millisecond * 10000)
    tick.Stop()
    fmt.Println("Ticker stopped")
}
