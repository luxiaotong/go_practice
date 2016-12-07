package main
import "fmt"

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i ++ {
        c <- x
        x, y = y, x + y
    }
    close(c)
}

func main() {
    n := 10
    ch := make(chan int, n)
    go fibonacci(n, ch)

    for i := range ch {
        fmt.Println(i)
    }
}
