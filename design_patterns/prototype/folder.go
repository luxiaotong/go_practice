package main

import "fmt"

type folder struct {
	name     string
	children []inode
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, c := range f.children {
		c.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cp := &folder{name: f.name + "_clone"}
	cc := make([]inode, 0, len(f.children))
	for _, c := range f.children {
		i := c.clone()
		cc = append(cc, i)
	}
	cp.children = cc
	return cp
}
