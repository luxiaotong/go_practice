package main
import "fmt"

func main() {
    var arr [2]int
    fmt.Println(arr)

    arr[0] = 0
    arr[1] = 1
    fmt.Println(arr)

    arr2 := [5]int {1, 2, 3, 4, 5}
    fmt.Println(arr2)

    arr3 := arr2[0:2]
    fmt.Println(arr3)

    arr4 := [...]int {4, 4, 4, 4}
    fmt.Println(arr4)

    for i,v := range arr4 {
        fmt.Println(i, v)
    }
}
