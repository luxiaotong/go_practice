package main

import (
    "bytes"
    "crypto/sha256"
    "fmt"
)

type BlockChain struct {
    blocks []*Block
}

type Block struct {
    Hash []byte
    Data []byte
    PrevHash []byte
}

func (b *Block) DeriveHash() {
    info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
    hash := sha256.Sum256(info)
    b.Hash = hash[:]
}

func CreateBlock(data []byte, prevHash []byte) *Block {
    block := &Block{[]byte{}, data, prevHash}
    block.DeriveHash()
    return block
}

func (chain *BlockChain) AddBlock(data string) {
    prevBlock := chain.blocks[len(chain.blocks)-1]
    currBlock := CreateBlock([]byte(data), prevBlock.Hash)
    chain.blocks = append(chain.blocks, currBlock)
}

func Genesis() *Block {
    return CreateBlock([]byte("Genesis"), []byte{})
}

func InitBlockChain() *BlockChain {
    return &BlockChain{[]*Block{Genesis()}}
}

func main() {
    chain := InitBlockChain()
    chain.AddBlock("First Block after Genesis")
    chain.AddBlock("Second Block after Genesis")
    chain.AddBlock("Third Block after Genesis")

    for _,block := range chain.blocks {
        fmt.Printf("Prev Hash: %x\n", block.PrevHash)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
    }
}
