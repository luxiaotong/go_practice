package main

import (
    "fmt"
    "bytes"
)

func main() {
    name := [][]byte{[]byte("Shannon"), []byte("Lu")}
    fmt.Printf("First name: %s\n", name[0])
    fmt.Printf("Second name: %s\n", name[1])
    sep := []byte("-")
    full_name := bytes.Join(name, sep)
    fmt.Printf("Full name: %s\n", full_name)

    fmt.Printf("Full name: %s\n", bytes.Join(name, []byte{}))
    fmt.Printf("Full name: %s\n", bytes.Join(name, []byte("")))
}
