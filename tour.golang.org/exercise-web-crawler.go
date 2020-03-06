package main

import (
    "fmt"
    "sync"
)

type Record struct {
    v map[string]string
    mux sync.Mutex
}

func (r *Record) Write(url string, body string) {
    r.mux.Lock()
    r.v[url] = body
    r.mux.Unlock()
}

func (r *Record) Look(url string) bool {
    r.mux.Lock()
    defer r.mux.Unlock()

    _, ok := r.v[url]
    return ok
}

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, r *Record, quit chan int) {
    // TODO: Fetch URLs in parallel.
    //ch := make(chan int)
    // TODO: Don't fetch the same URL twice.
    // This implementation doesn't do either:
    if depth <= 0 {
        //quit <- 1
        return
    }
    if ( r.Look(url) ) {
        return
    }
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        r.Write(url, "111")
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    r.Write(url, body)
    for _, u := range urls {
        go Crawl(u, depth-1, fetcher, r, quit)
        //fmt.Println(<-quit)
        //quit <- 1
    }
    return
}

func main() {
    quit := make(chan int)
    r := &Record{v: make(map[string]string)}
    go Crawl("http://golang.org/", 4, fetcher, r, quit)
    fmt.Println(<-quit)

    //for k, v := range r.v {
        //fmt.Println(k, v)
    //}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}

