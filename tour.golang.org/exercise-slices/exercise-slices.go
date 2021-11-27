package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    ans := [][]uint8{}
    for i:=0;i<dx;i++ {
        y := make([]uint8, dy)
        for j:=0;j<dy;j++ {
            y[j] = uint8((i+j)/2)
        }
        ans = append(ans, y)
    }
    println(ans)
    return ans 
}

func main() {
    pic.Show(Pic)
}
