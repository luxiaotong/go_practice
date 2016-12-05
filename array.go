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
}
