package main
import "fmt"

func main() {
    var arr[5]int = [5]int {1, 2, 3, 4, 5}
    fmt.Println(arr)

    var slc1[]int = arr[0:1]
    fmt.Println(slc1)

    var slc2[]int = arr[:5]
    fmt.Println(slc2)
}
