package blockchain

import (
    "bytes"
    "crypto/sha256"
    _ "fmt"
    "encoding/gob"
    "log"
)

/*
type BlockChain struct {
    Blocks []*Block
}
*/

type Block struct {
    Hash []byte
    Data []byte
    PrevHash []byte
    Nonce int64
}

func (b *Block) DeriveHash() {
    info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
    hash := sha256.Sum256(info)
    b.Hash = hash[:]
}

func CreateBlock(data []byte, prevHash []byte) *Block {
    //This code for part1
    //block := &Block{[]byte{}, data, prevHash}
    //block.DeriveHash()

    //This code for part2
    block := &Block{[]byte{}, data, prevHash, 0}
    pow := NewProof(block)
    nonce, hash := pow.Run()
    block.Hash = hash
    block.Nonce = nonce

    return block
}

func Genesis() *Block {
    return CreateBlock([]byte("Genesis"), []byte{})
}

/*
func (chain *BlockChain) AddBlock(data string) {
    prevBlock := chain.Blocks[len(chain.Blocks)-1]
    currBlock := CreateBlock([]byte(data), prevBlock.Hash)
    chain.Blocks = append(chain.Blocks, currBlock)
}

func InitBlockChain() *BlockChain {
    return &BlockChain{[]*Block{Genesis()}}
}
*/

func (b *Block) Serialize() []byte {
    var res bytes.Buffer
    enc := gob.NewEncoder(&res)
    err := enc.Encode(b)
    Handle(err)
    return res.Bytes()
}

func Deserialize(data []byte) *Block {
    var block *Block
    dec := gob.NewDecoder(bytes.NewReader(data))
    err := dec.Decode(&block)
    Handle(err)
    return block
}

func Handle(err error) {
    if err != nil {
        log.Panic(err)
    }
}
