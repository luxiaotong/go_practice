package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("Cannot Sqrt negative number: %v\n", float64(e))
}

func Sqrt(x float64) (float64,ErrNegativeSqrt) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }

    z := 1.0

    for i := 0; i < 10; i ++ {
        z = z - ((z*z - x) / (2*z))
        //fmt.Println(z)
    }

    return z, ErrNegativeSqrt(0)
}

func main() {
    fmt.Println(Sqrt(10))
    fmt.Println(Sqrt(-10))
}
