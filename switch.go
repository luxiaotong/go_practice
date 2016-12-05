package main

import (
    "fmt"
    "runtime"
)

func main() {
    switch os := runtime.GOOS; os {
        case "darwin" : fmt.Println("OSX");
        case "linux" : fmt.Println("Linux");
        default : fmt.Println(os);
    }
}
