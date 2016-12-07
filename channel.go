package main
import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, i := range s {
        sum += i
    }

    c <- sum
}

func main() {
    s := []int {7, 2, 8, -9, 4, 0}
    c := make(chan int)
    mid := len(s) / 2
    go sum(s[:mid], c)
    go sum(s[mid:], c)

    x, y := <-c, <-c
    fmt.Println(x, y, x+y)
}
