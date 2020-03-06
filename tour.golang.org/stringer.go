package main
import "fmt"

type Person struct {
    Name string
    Age int
}

func (p Person)String() string{
    return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func main() {
    p1 := Person{"John", 21}
    p2 := Person{"Frank", 22}

    fmt.Print(p1)
    fmt.Print(p2)
}
