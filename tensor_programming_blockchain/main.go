package main

import (
    "github.com/luxiaotong/go_practice/tensor_programming_blockchain/blockchain"
    "fmt"
    "strconv"
)


func main() {
    chain := blockchain.InitBlockChain()
    chain.AddBlock("First Block after Genesis")
    chain.AddBlock("Second Block after Genesis")
    chain.AddBlock("Third Block after Genesis")

    for _,block := range chain.Blocks {
        fmt.Printf("Prev Hash: %x\n", block.PrevHash)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)

        pow := blockchain.NewProof(block)
        fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
    }

}
