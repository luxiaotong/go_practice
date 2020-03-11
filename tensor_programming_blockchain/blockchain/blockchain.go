package blockchain

import (
    "fmt"
    "github.com/dgraph-io/badger"
)

const (
    dbPath = "./tmp/blocks/"
)

type BlockChain struct {
    LastHash []byte
    Database *badger.DB
}

type BlockChainIterator struct {
    CurrentHash []byte
    Database *badger.DB
}

func InitBlockChain() *BlockChain {
    var lastHash []byte

    db, err := badger.Open(badger.DefaultOptions(dbPath))
    Handle(err)

    err = db.Update(func(txn *badger.Txn) error {
        if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
            fmt.Println("No existing blockchain found")
            gBlock := Genesis()
            fmt.Println("Genesis proved")

            err := txn.Set(gBlock.Hash, gBlock.Serialize())
            Handle(err)

            err = txn.Set([]byte("lh"), gBlock.Hash)

            lastHash = gBlock.Hash

            return err
        } else {
            item, err := txn.Get([]byte("lh"))
            Handle(err)
            lastHash, err = item.ValueCopy(nil)
            return err
        }
        return err
    })
    Handle(err)

    return &BlockChain{lastHash, db}
}

func (chain *BlockChain) AddBlock(data string) {
    var lastHash []byte
    err := chain.Database.View(func(txn *badger.Txn) error {
        item, err := txn.Get([]byte("lh"))
        Handle(err)
        lastHash, err = item.ValueCopy(nil)
        return err
    })
    Handle(err)

    newBlock := CreateBlock([]byte(data), lastHash)
    err = chain.Database.Update(func(txn *badger.Txn) error {
        err = txn.Set(newBlock.Hash, newBlock.Serialize())
        Handle(err)

        err := txn.Set([]byte("lh"), newBlock.Hash)

        chain.LastHash = newBlock.Hash
        return err
    })
    Handle(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
    iter := &BlockChainIterator{chain.LastHash, chain.Database}
    return iter
}

func (iter *BlockChainIterator) Next() *Block {
    var block *Block
    var encodedBlock []byte
    err := iter.Database.View(func(txn *badger.Txn) error {
        item, err := txn.Get(iter.CurrentHash)
        Handle(err)
        encodedBlock, err = item.ValueCopy(nil)
        block = Deserialize(encodedBlock)
        return err
    })
    Handle(err)

    iter.CurrentHash = block.PrevHash

    return block
}
