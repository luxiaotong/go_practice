package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    go func() {
        time.Sleep(time.Second * 2)
        ch1 <- "Request 1"
    } ()

    select {
        case v := <-ch1:
            fmt.Println(v)
        case <- time.After(time.Second):
            fmt.Println("Timeout 1")
    }

    ch2 := make(chan string)
    go func() {
        time.Sleep(time.Second * 2)
        ch2 <- "Request 2"
    } ()

    select {
        case v := <-ch2:
            fmt.Println(v)
        case <- time.After(time.Second * 3):
            fmt.Println("Timeout 2")
    }
}
