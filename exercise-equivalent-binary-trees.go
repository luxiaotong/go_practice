package main

import (
    "fmt"
    "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if ( t.Left != nil ) {
        Walk(t.Left, ch)
    }
    ch <- t.Value
    if ( t.Right != nil ) {
        Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    rst := true
    walk_ch1 := make(chan int)
    walk_ch2 := make(chan int)

    go Walk(t1, walk_ch1)
    go Walk(t2, walk_ch2)

    for i := 0; i < 10; i ++ {
        x := <- walk_ch1
        y := <- walk_ch2

        if ( x != y ) {
            rst = false
            break
        }
    }

    return rst
}

func main() {
    ch := make(chan int)

    go Walk(tree.New(1), ch)
    for i := 0; i < 10; i ++ {
        fmt.Printf("%v ", <-ch)
    }
    fmt.Println()

    go Walk(tree.New(2), ch)
    for i := 0; i < 10; i ++ {
        fmt.Printf("%v ", <-ch)
    }
    fmt.Println()

    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
    fmt.Println(Same(tree.New(2), tree.New(2)))
}
