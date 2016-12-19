package main

import (
    "fmt"
    "time"
)

func main() {
    requests := make(chan int, 5)
    for i := 0; i < 5; i ++ {
        requests <- i
    }
    close(requests)

    tick := time.Tick(time.Millisecond * 200)

    for r := range requests {
        <- tick
        fmt.Println(time.Now(), "request", r)
    }
}
