package main

import (
    "golang.org/x/tour/wc"
    "strings"
    "fmt"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    fmt.Printf("words:%q\n", words)
    ans := make(map[string]int)
    for _,word := range words {
        i, ok := ans[word]
        if ok {
            ans[word] = int(i+1)
        } else {
            ans[word] = int(1)
        }
    }
    fmt.Printf("ans:%q\n", ans)
    return ans
}

func main() {
    wc.Test(WordCount)
}
