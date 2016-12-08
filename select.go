package main

import "fmt"

func main() {
    ch := make(chan int)
    quit := make(chan int)

    go func() {
        for i := 0; i < 10; i ++ {
            fmt.Println(<-ch)
        }
        quit <- 1
    }()

    fibonacci(ch, quit)
}

func fibonacci(ch, quit chan int) {
    x, y := 0, 1
    for {
        select {
            case ch <- x:
                x, y = y, x + y
            case <- quit:
                fmt.Println("quit")
                return
        }
    }
}
