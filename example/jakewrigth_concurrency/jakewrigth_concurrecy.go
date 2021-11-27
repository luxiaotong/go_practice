// https://www.youtube.com/watch?v=LvgVSSpwND8&t=442s
package main

import (
    "fmt"
    "time"
    _ "sync"
)

func count(str string) {
    for i:=1;true;i++ {
        fmt.Println(i, str)
        time.Sleep(time.Millisecond * 500)
    }
}

func count_by_waitgroup(str string) {
    for i:=1; i<=5 ;i++ {
        fmt.Println(i, str)
        time.Sleep(time.Millisecond * 500)
    }
}

func count_by_channel(str string, ch chan string) {
    for i:=1; i<=5 ;i++ {
        ch <- str
        time.Sleep(time.Millisecond * 500)
    }
    close(ch)
}

func worker(jobs <-chan int, results chan<- int) {
    for n := range jobs {
        results <- fib(n)
    }
}

func fib(n int) int {
    if n <= 1 {
        return 1
    }
    return fib(n-1) + fib(n-2)
}

func main() {
    /*
    go count("sheep")
    go count("fish")
    time.Sleep(time.Second * 2)
    fmt.Scanln()
    */

    //WaitGroup
    /*
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        count_by_waitgroup("sheep")
        wg.Done()
    }()
    wg.Wait()
    */

    //Channel
    /*
    ch := make(chan string)
    go count_by_channel("sheep", ch)
    for {
        msg, open := <-ch
        if !open {
            break
        }
        fmt.Println(msg)
    }
    for msg := range ch {
        fmt.Println(msg)
    }
    */

    //Channel Buffer
    /*
    ch := make(chan string, 2)
    //ch := make(chan string)
    //go func() {
    //    ch <- "hello"
    //}()
    ch <- "hello"
    ch <- "world"
    //ch <- "three" //Wrong
    msg := <-ch
    fmt.Println(msg)
    msg = <-ch
    fmt.Println(msg)
    */

    //Channel Select
    /*
    go func() {
        for {
            ch1 <- "Every 500ms"
            time.Sleep(time.Millisecond * 500)
        }
    }()
    go func() {
        for {
            ch2 <- "Every 2s"
            time.Sleep(time.Second * 2)
        }
    }()

    for {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
    */

    //Worker Pool
    /*
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    go worker(jobs, results)
    //go worker(jobs, results)
    //go worker(jobs, results)
    //go worker(jobs, results)
    for i:=0;i<100;i++ {
        jobs <- i
    }
    close(jobs)
    for j:=0;j<100;j++ {
        fmt.Println(<-results)
    }
    */
}
