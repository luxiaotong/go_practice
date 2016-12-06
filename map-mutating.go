package main
import "fmt"

func main() {
    var m = make(map[string]string)
    m["a"] = "aaa"
    m["b"] = "bbb"

    item,ok := m["a"]
    fmt.Println(ok, item)

    item,ok = m["c"]
    fmt.Println(ok, item)
}
