package main
import "fmt"
func main() {
    for i := 0; i < 10; i ++ {
        if remainder := i % 2; remainder == 0 {
            fmt.Println(i, "is even");
        } else {
            fmt.Println(i, "is odd");
        }
    }
}
