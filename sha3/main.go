package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pubkeys := []string{
		"04e50e6f0e821534f8ac20ea960b6c7fa318f569b5350332e94e8205f774120be4745dd2c5daaa405a6c89acdfc3d2889c86824e21baa1ae61468821e5611d7e3e",
	}
	for _, pk := range pubkeys {
		b, err := hex.DecodeString(pk)
		if err != nil {
			panic(err)
		}
		d := crypto.Keccak256(b)

		fmt.Printf("pub key: %x\n", b)
		fmt.Printf("32bytes: %x\n", d)
		fmt.Printf("20bytes: %x\n", d[12:])
	}
}
