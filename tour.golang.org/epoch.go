package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now)

    secs := now.Unix()
    fmt.Println(secs)

    nanos := now.UnixNano()
    fmt.Println(nanos)

    millis := nanos / (1000*1000)
    fmt.Println(millis)

    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}
