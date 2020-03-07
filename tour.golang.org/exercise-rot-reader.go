package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r13r *rot13Reader) Read(b []byte) (n int, err error) {
    n,err = r13r.r.Read(b)
    for i:=0;i<n;i++ {
        c := b[i]
        if c>='a' && c<='n' || c>='A' && c<='N' {
            b[i] = c+13
        }
        if c>='m' && c<='z' || c>='M' && c<='Z' {
            b[i] = c-13
        }
    }
    return
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
