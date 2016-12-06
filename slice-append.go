package main
import "fmt"

func main() {
    var s []int
    printSlice(s)
    if ( s == nil ) {
        fmt.Println("s is nil")
    }

    s = append(s, 0)
    printSlice(s)

    s = append(s, 1, 2, 3, 4)
    printSlice(s)

    var s2 []int = make([]int, 5)
    printSlice(s2)

    s2 = s2[:1]
    printSlice(s2)

    s2 = append(s2, 1, 2)
    printSlice(s2)

    s2 = append(s2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
    printSlice(s2)
}

func printSlice(s[]int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s);
}
