package main

import (
	"fmt"
	"strings"
)

func matchParen(s string) string {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s[i:], ")")
		if j >= 0 {
			return s[i+1 : j+i]
		}
	}
	return ""
}

func main() {

	s := `int(10)`
	fmt.Println(matchParen(s))

	s = `double(10,2)`
	fmt.Println(matchParen(s))

	// r := regexp.MustCompile(`\((.*?)\)`)
	// s := `SomeText(10)`
	// rs := r.FindStringSubmatch(s)
	// fmt.Println(rs)

	// s = `int(10)`
	// rs = r.FindStringSubmatch(s)
	// fmt.Println(rs)

	// s = `text`
	// rs = r.FindStringSubmatch(s)
	// fmt.Println(rs)

}
