package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

// Take the data from block

// Create a counter(nonce) which starts at 0

// Create a hash of data plus the counter

// Check the hash if it meets a set of requirements

// Requirements:
// The first few bytes must contain 0s

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	//fmt.Printf("target is: %b\n", target)

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.Data,
			pow.Block.PrevHash,
			ToHex(nonce),
			ToHex(Difficulty),
		},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	//err := binary.Write(buff, binary.LittleEndian, num)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//fmt.Printf("hex: %x\n", buff.Bytes())

	return buff.Bytes()
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var intHash big.Int
	var hash [32]byte
	var nonce int64 = 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		//fmt.Printf("hash:")
		fmt.Printf("\r%x", hash)

		//fmt.Println()
		intHash.SetBytes(hash[:])
		//fmt.Printf("int hash:")
		//fmt.Printf("\r%d", hash)
		//fmt.Println()

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	//fmt.Printf("validate int hash: %b\n", intHash)
	//fmt.Printf("validate target : %b\n", pow.Target)
	return intHash.Cmp(pow.Target) == -1
}
