package main

import "fmt"

func main() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}

	folder1 := &folder{
		name:     "folder1",
		children: []inode{file1},
	}
	folder2 := &folder{
		name:     "folder2",
		children: []inode{folder1, file2, file3},
	}

	fmt.Println("print folder 1")
	folder1.print("    ")

	fmt.Println("print folder 2")
	folder2.print("    ")

	folderc := folder2.clone()
	fmt.Println("print folder clone")
	folderc.print("    ")
}
