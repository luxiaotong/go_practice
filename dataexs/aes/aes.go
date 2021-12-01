package main

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/forgoer/openssl"
)

func main() {
	secret := "31c241309a9231f585bca20c9873b49a"
	n := len(secret) / 2
	bb := make([]byte, n)
	k := 0
	for i := 0; i < n; i++ {
		var high, low int64
		var h, l byte
		high, _ = strconv.ParseInt(string(secret[k]), 16, 64)
		low, _ = strconv.ParseInt(string(secret[k+1]), 16, 64)
		h = byte(high) & 0xff
		l = byte(low) & 0xff
		bb[i] = byte(h<<4 | l)
		k += 2
	}
	fmt.Println("bb: ", bb)
	cc, _ := hex.DecodeString(secret)
	fmt.Println("cc: ", cc)
	src := []byte("20211117001:31c241309a9231f585bca20c9873b49a:1638343975067")
	dst, _ := openssl.AesECBEncrypt(src, bb, openssl.PKCS5_PADDING)
	fmt.Printf("dst: %x\n", dst)
}
