package main
import "fmt"
import "unsafe"

func main() {
    var a bool = true
    fmt.Printf("true is %d byte\n", unsafe.Sizeof(a))
}
