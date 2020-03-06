package main

import (
    "fmt"
    "time"
)

func work(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {

    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for i := 0; i < 3; i ++ {
        go work(i, jobs, results)
    }

    for i := 0; i < 10; i ++ {
        jobs <- i
    }
    close(jobs)

    for i := 0; i < 10; i ++ {
        <-results
    }
}
