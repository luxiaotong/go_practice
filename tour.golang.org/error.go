package main
import "fmt"

type TmpMap map[string]string

func (tm TmpMap) Error() string {
    return fmt.Sprintf("Not found")
}

func main() {
    tm := make(TmpMap)
    tm["a"] = "aaa"
    tm["b"] = "bbb"
    fmt.Println(tm)

    tmp, ok := tm["c"]
    fmt.Println(ok, tmp)
}
