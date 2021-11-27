package main
import "fmt"

var x, y = 1, 5;
//x, y := 1, 5;

func main() {

    fmt.Println(sum(x, y));

    x, y = swap(x, y);
    fmt.Println(x, y);
}

func sum(x int, y int) int {
    return x+y;
}

func swap(x int, y int) (int, int) {
    return y, x;
}
