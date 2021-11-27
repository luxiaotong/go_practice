package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        time.Sleep(time.Second)
        ch1 <- 1
    }()

    go func() {
        time.Sleep(time.Second * 2)
        ch1 <- 2
    }()

    for i := 0; i < 2; i ++ {
        select {
            case x := <- ch1:
                fmt.Println(x)
            case y := <- ch2:
                fmt.Println(y)
            default:
                fmt.Println("Waiting...")
                time.Sleep(time.Second/2)
                i -= 1
        }
    }
}
