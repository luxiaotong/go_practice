package main
import "fmt"

func main() {
    var i interface{}

    i = "Hello"
    t := i.(string)
    fmt.Println(t)

    a, ok := i.(float64)
    if ( ok ) {
        fmt.Println(a)
    } else {
        fmt.Println(ok)
    }
}
