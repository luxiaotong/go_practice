package main

import (
	"os"

	"github.com/luxiaotong/go_practice/tensor_programming_blockchain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
