package main
import "fmt"

type Vertex struct {
    x int
    y int
}

func main() {
    m := make(map[string]Vertex)
    m["Bell Labs"] = Vertex{1,2}

    fmt.Println(m)
}
