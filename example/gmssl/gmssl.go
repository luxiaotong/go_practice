package main

import (
	"fmt"

	"github.com/guanzhi/GmSSL/go/gmssl"
)

// gmssl ecparam -genkey -name sm2p256v1 -out priv.pem
// gmssl sm2 -pubout -in priv.pem -out pub.pem
// echo "f31fdf2bec0c44a09d4edf089a40ae30" | gmssl sm2utl -encrypt -pubin -inkey pub.pem -out ciphertext.sm2
// gmssl sm2utl -decrypt -inkey priv.pem -in ciphertext.sm2

func main() {
	// cipher := "040f758411b04b205b7ac5596408c5dd87d331be17c2e9594d374fd0858ed68da920b61234dab411bf8f86fd60fc51663c50fe01ce5b07451353c9189d5da99d3a5b4388c47f7cdda4ed2f2063dda3b957bef4079771e09369e72b99c5d857268a460120fd081450e7268a2b53cd5d34e68ce2d762f51265521511747c91b67c87"
	// cipher := "400E8D8EC5F129D0547374AACC581D2CBB44389BCE165CD4C7F07C27E9D8EB0F7ED92A8EB54E83F8D2D5A46AB1D6BFB58B4667A5589CD9B79AFF5E095B90F26A0B6E9A9714BF95A6019107447170C821C075367392223E66128F526E240D3398898A099197F2C3B6DA683686C4159F99"
	// cipher := "0430818a02207b6cf290b8613931d496ad7f545f0ac292faf68e875bd5154dd1b0051584ccfb02210081c1efbaca3566912b28938f0163ca528136817a86601ddb530ea97e109e908a0420fea4e0d6e3aa3d7482521ea3ff5d357d7cc59f142aacd0a126d9945748a3d8f6042169dc118c46558a37fa3da7db8550e7526b8945392b68689b806c819462937ff674"
	// b, err := hex.DecodeString(cipher)
	// if err != nil {
	// 	panic(err)
	// }
	// ioutil.WriteFile("/Users/luxiaotong/code/go_practice/example/gmssl/ciphertext2.sm2", b, 0644)
	pem := "/Users/luxiaotong/code/go_practice/example/gmssl/pub.pem"
	pub, err := gmssl.NewPublicKeyFromPEM(pem)
	if err != nil {
		panic(err)
	}
	s, err := pub.GetText()
	if err != nil {
		panic(err)
	}
	fmt.Println("pubkey: ", s)
}
