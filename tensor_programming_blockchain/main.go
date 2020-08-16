package main

import (
	"os"

	"github.com/luxiaotong/go_practice/tensor_programming_blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cli := &cli.CommandLine{}
	cli.Run()
}
