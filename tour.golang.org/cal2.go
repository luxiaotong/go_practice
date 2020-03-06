package main

import "fmt"
import "time"

func main(){
    t1 := time.Now()
    count :=0
    for i :=0; i< 9000000000; i++{
        count =count +i
    }
    t2 := time.Now()
    fmt.Printf("cost:%d,count:%d\n",t2.Sub(t1),count)
}
