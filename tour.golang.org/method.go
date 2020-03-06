package main
import(
    "fmt"
    "math"
)

type Vertex struct {
    x, y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.x*v.x+v.y*v.y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}
