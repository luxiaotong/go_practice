package main

import "fmt"

func main() {
    queue := make(chan int, 5)
    queue <- 1
    queue <- 2
    close(queue)

    for i := range queue {
        fmt.Println(i)
    }
}
