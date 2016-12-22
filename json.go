package main

import (
    "fmt"
    "encoding/json"
)

type Response1 struct {
    Page int
    Fruits []string
}

type Response2 struct {
    Page int
    Fruits []string
}

func main() {
    //Encoding
    r1 := Response1{1, []string{"apple", "apple", "apple"}}
    rJson, _ := json.Marshal(r1)
    fmt.Println(string(rJson))

    //Decoding
    var r2 Response2
    json.Unmarshal(rJson, &r2)
    fmt.Println(r2)
}
