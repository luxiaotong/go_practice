package main
import "fmt"
func main() {
    var i, j int = 1, 2
    fmt.Println(i, j)

    var pi *int = &i
    var pj *int = &j
    fmt.Println(*pi, *pj)

    *pi = 100
    *pj = 200
    fmt.Println(i, j)
}
