package main

import (
    "fmt"
    "time"
    "sync"
)

type SafeCounter struct {
    //v map[string]int
    v int
    mux sync.Mutex
}

func (sc *SafeCounter) Inc(key string) {
    sc.mux.Lock()
    //sc.v[key] ++
    sc.v ++
    sc.mux.Unlock()
}

func (sc *SafeCounter) Value(key string) int {
    sc.mux.Lock()
    defer sc.mux.Unlock()

    //return sc.v[key]
    return sc.v
}

func main() {
    //sc := SafeCounter{v: make(map[string]int)}
    sc := SafeCounter{v: 0}
    for i := 0; i < 1000; i ++ {
        go sc.Inc("test")
    }
    time.Sleep(time.Second)
    fmt.Println(sc.Value("test"))
}
