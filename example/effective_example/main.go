package main

import (
	"fmt"
	"runtime"
	"sort"
)

func main() {
	v := 1
	// v := 2 // Wrong!
	v, v2 := 2, 3 // Right!
	fmt.Println(v, v2)

	a := []int{1, 2, 3}
	fmt.Printf("%T\n", a)
	// for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	for i, j := 0, len(a)-1; i < j; {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	fmt.Println(a)

	//Switch and Break
A:
	switch {
	case v == 2:
		break A
	}
	fmt.Println("Switch Break")

	//For and Break
B:
	for i := 0; i < 3; i++ {
		if v == 2 {
			break B
		}
		fmt.Println(i, v)
	}
	fmt.Println("For Break")

	// for i := 0; i < 5; i++ {
	// 	defer fmt.Printf("%d ", i)
	// }

	b()
	const Enone = 0
	const Eio = 1
	const Einval = 2
	e := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	fmt.Println(e)

	//New & Make
	f := new(int)
	fmt.Printf("new: %d\n", *f)
	g := make([]int, 3, 5)
	g1 := append(g, []int{0, 0}...)
	fmt.Println("make g : ", g)
	fmt.Println("g cap: ", cap(g))
	fmt.Printf("g addr %p: ", g)
	fmt.Println("make g1 : ", g1)
	fmt.Printf("g1 addr %p: \n", g1)

	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	//Printf
	type T struct {
		a int
		b float64
		c string
	}
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", timeZone)

	fmt.Printf("string: %#q\n", "test string")

	//Const
	type ByteSize float64
	const (
		_           = iota // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println("KB: ", KB)
	fmt.Println("1<<10: ", 1<<10)
	fmt.Println("MB: ", MB)
	fmt.Println("1<<20: ", 1<<20)

	// Interface
	h := Sequence{9, 8, 7, 6, 5, 4, 3, 2, 2}
	sort.Sort(h)
	fmt.Println("sort: ", h)

	// Concurrency
	// sem := make(chan int, 2)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println("send:", i)
	// 		sem <- i
	// 	}
	// 	close(sem)
	// }()
	// for s := range sem {
	// 	fmt.Println("receive: ", s)
	// }

	sem := make(chan int, 2)
	go func() {
		sem <- 1
		fmt.Println("send:", 1)
		sem <- 2
		fmt.Println("send:", 2)
		sem <- 3
		fmt.Println("send:", 3)
		sem <- 4
		fmt.Println("send:", 4)
		sem <- 5
		fmt.Println("send:", 5)
		close(sem)
	}()

	for s := range sem {
		fmt.Println("receive: ", s)
	}
	fmt.Println("NumCPU: ", runtime.NumCPU())

}

// Defer

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

// Sequence Interface
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
// func (s Sequence) String() string {
// 	s = s.Copy() // Make a copy; don't overwrite argument.
// 	sort.Sort(s)
// 	str := "["
// 	for i, elem := range s { // Loop is O(N²); will fix that in next example.
// 		if i > 0 {
// 			str += " | "
// 		}
// 		str += fmt.Sprint(elem)
// 	}
// 	return str + "]"
// }

func (s Sequence) String() string {
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}
