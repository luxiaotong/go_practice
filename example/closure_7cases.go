// https://zhuanlan.zhihu.com/p/92634505

package main

import (
    "fmt"
    "time"
)

func foo1(x *int) func() {
    return func() {
        *x = *x + 1
        fmt.Printf("foo1 val = %d\n", *x)
    }
}

func foo2(x int) func() {
    return func() {
        x = x + 1
        fmt.Printf("foo2 val = %d\n", x)
    }
}

func foo3() {
    values := [4]int{1,2,3,5}
    for _, val := range values {
        fmt.Printf("foo3 val = %d\n", val)
    }
}

func show(v interface{}) {
    fmt.Printf("foo4 val = %d\n", v)
}
func foo4() {
    values := []int{1,2,3,5}
    for _, val := range values {
        show(val)
    }
}

func foo5() {
    values := []int{1,2,3,5}
    for _, val := range values {
        go func() {
            fmt.Printf("foo5 val = %d\n", val)
        }()
    }
}

var foo6Chan = make(chan int, 10)
func foo6() {
    for val := range foo6Chan {
        go func() {
            fmt.Printf("foo6 val = %d\n", val)
        }()
    }
}

func foo7(x int) []func() {
    var fs []func()
    values := []int{1,2,3,5}
    for _, val := range values {
        fs = append(fs, func() {
            fmt.Printf("foo7 val = %d\n", x + val)
        })
    }
    return fs
}

/*
Q1：
第一组实验：假设现在有变量x=133，并创建变量f1和f2分别为foo1(&x)和foo2(x)的返回值，请问多次调用f1()和f2()会打印什么？
第二组实验：重新赋值变量x=233，请问此时多次调用f1()和f2()会打印什么？
第三组实验：如果直接调用foo1(&x)()和foo2(x)()多次，请问每次都会打印什么？
Q2：
请问分别调用函数foo3()，foo4()和foo5()，分别会打印什么？
Q3：
第一组实验：如果“几乎同时”往channelfoo6Chan里面塞入一组数据"1,2,3,5"，foo6会打印什么？
第二组实验：如果以间隔纳秒（10^-9秒）的时间往channel里面塞入一组数据，此时foo6又会打印什么？
第三组实验：如果是微秒（10^-6秒）呢？如果是毫秒（10^-3秒）呢？如果是秒呢？
Q4：
请问如果创建变量f7s=foo7(11)，f7s是一个函数集合，遍历f7s会打印什么？
*/

func main() {
    // Q1第一组实验
    x := 133
    f1 := foo1(&x)
    f2 := foo2(x)
    f1() //134
    f2() //134
    f1() //135
    f2() //135
    fmt.Println()

    // Q1第二组
    x = 233
    f1() //234
    f2() //136
    f1() //235
    f2() //137
    fmt.Println()

    // Q1第三组
    foo1(&x)() //236
    foo2(x)()  //237
    foo1(&x)() //237
    foo2(x)()  //238
    foo2(x)()  //238
    fmt.Println()

    // Q4实验：
    f7s := foo7(11)
    for _, f7 := range f7s {
        f7()
    }
    fmt.Println()

    foo3()
    fmt.Println()
    foo4()
    fmt.Println()
    foo5()
    fmt.Println("sleep 500ms")
    time.Sleep(500 * time.Microsecond)
    fmt.Println()

    // Q3第一组实验
    go foo6()
    foo6Chan <- 1
    foo6Chan <- 2
    foo6Chan <- 3
    foo6Chan <- 5
    fmt.Println()
    fmt.Scanln("Next")

    // Q3第二组实验
    foo6Chan <- 11
    time.Sleep(time.Duration(1) * time.Nanosecond)
    foo6Chan <- 12
    time.Sleep(time.Duration(1) * time.Nanosecond)
    foo6Chan <- 13
    time.Sleep(time.Duration(1) * time.Nanosecond)
    foo6Chan <- 15
    fmt.Println()
    fmt.Scanln("Next")

    // Q3第三组实验
    // 微秒
    foo6Chan <- 21
    time.Sleep(time.Duration(1) * time.Microsecond)
    foo6Chan <- 22
    time.Sleep(time.Duration(1) * time.Microsecond)
    foo6Chan <- 23
    time.Sleep(time.Duration(1) * time.Microsecond)
    foo6Chan <- 25
    time.Sleep(time.Duration(10) * time.Second)
    fmt.Println()
    fmt.Scanln("Next")

    // 毫秒
    foo6Chan <- 31
    time.Sleep(time.Duration(1) * time.Millisecond)
    foo6Chan <- 32
    time.Sleep(time.Duration(1) * time.Millisecond)
    foo6Chan <- 33
    time.Sleep(time.Duration(1) * time.Millisecond)
    foo6Chan <- 35
    time.Sleep(time.Duration(10) * time.Second)
    fmt.Println()
    fmt.Scanln("Next")

    // 秒
    foo6Chan <- 41
    time.Sleep(time.Duration(1) * time.Second)
    foo6Chan <- 42
    time.Sleep(time.Duration(1) * time.Second)
    foo6Chan <- 43
    time.Sleep(time.Duration(1) * time.Second)
    foo6Chan <- 45
    time.Sleep(time.Duration(10) * time.Second)

    // 实验完毕，最后记得关闭channel
    close(foo6Chan)
}
