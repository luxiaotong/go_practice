package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
)

func main() {
	privKey := "312D0E4740B1438739B44AE12EC3B990D3DDEE9BB95FC2F523E06ED197449946"
	pubKey := "7CC3F8ED6D541BB6D02406A51B2DC9689DB459353222F503CD571FDCF5029220BB50A222948554DC47093B19C1154DDFAC867EE0CD79029DB0946E442437B8EC"

	priv, err := sm2.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	d := new(big.Int)
	d.SetString(privKey, 16)
	// log.Debug("bigint:", d)
	x := new(big.Int)
	y := new(big.Int)
	x.SetString(pubKey[0:64], 16)
	y.SetString(pubKey[64:128], 16)
	fmt.Printf("x: %v, y: %v\n", pubKey[0:64], pubKey[64:128])
	priv.D = d
	priv.PublicKey.X = x
	priv.PublicKey.Y = y

	cipher := "040f758411b04b205b7ac5596408c5dd87d331be17c2e9594d374fd0858ed68da920b61234dab411bf8f86fd60fc51663c50fe01ce5b07451353c9189d5da99d3a5b4388c47f7cdda4ed2f2063dda3b957bef4079771e09369e72b99c5d857268a460120fd081450e7268a2b53cd5d34e68ce2d762f51265521511747c91b67c87"
	data, err := hex.DecodeString(cipher)
	if err != nil {
		panic(err)
	}
	decrypted, err := sm2.Decrypt(priv, data, sm2.C1C3C2)
	if err != nil {
		panic(err)
	}
	fmt.Println("dec: ", string(decrypted))
}
