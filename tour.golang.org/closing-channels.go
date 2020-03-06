package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan int)

    go func() {
        for {
            v, more := <-jobs
            if more {
                fmt.Println("Received job", v)
            } else {
                fmt.Println("Received all jobs")
                done <- 1
                return
            }
        }
    } ()

    for i := 1; i <= 3; i ++ {
        jobs <- i
        fmt.Println("Sent job")
    }
    close(jobs)
    fmt.Println("Sent all jobs")

    <-done
}
