package main

import (
    "fmt"
    "os"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
    }()

    _, err := os.Create("/tmp/file/file")
    if err != nil {
        panic(err)
    }

    //Nothing happends here
    fmt.Println("continue...")
}
