package main

import (
	"fmt"
	"time"
)

func main() {
	drsPre := "deal:drs:"
	amtPre := "deal:amt:"
	t := time.Now()
	drs := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 8, 10, 15, 21, 25, 30, 33, 35, 37, 40, 41, 42, 45, 47, 49, 50, 52}
	amt := []float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.5, 0.8, 1.0, 1.2, 1.5, 2.0, 3.0, 3.5, 4.1, 5.2, 6.0, 6.5, 7.0, 7.3, 7.8, 8.0, 9.2, 10}
	cmd1 := "mset"
	cmd2 := "mset"
	for h := 0; h < 16; h++ {
		min0 := time.Date(t.Year(), t.Month(), t.Day(), h, 0, 0, 0, time.Local).Format("200601021504")
		min30 := time.Date(t.Year(), t.Month(), t.Day(), h, 30, 0, 0, time.Local).Format("200601021504")
		cmd1 += fmt.Sprintf(" %s%s %d", drsPre, min0, drs[h*2])
		cmd1 += fmt.Sprintf(" %s%s %d", drsPre, min30, drs[h*2+1])
		cmd2 += fmt.Sprintf(" %s%s %.2f", amtPre, min0, amt[h*2])
		cmd2 += fmt.Sprintf(" %s%s %.2f", amtPre, min30, amt[h*2+1])
	}
	fmt.Println(cmd1)
	fmt.Println(cmd2)
}
