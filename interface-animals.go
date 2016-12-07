package main
import "fmt"

type Animal interface {
    Speak()
}

type Dog struct {}
func (d Dog) Speak() {
    fmt.Println("Woof!")
}

type Cat struct {}
func (c Cat) Speak() {
    fmt.Println("Meow!")
}

func main() {
    d := Dog{}
    c := Cat{}

    var a Animal
    a = d
    a.Speak()

    a = c
    a.Speak()
}
