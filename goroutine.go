package main
import (
    "fmt"
    "time"
)

func say(sth string) {
    for i:=5; i >= 0; i -- {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(sth)
    }
}

func main() {
    go say("World")
    say("Hello")
}
