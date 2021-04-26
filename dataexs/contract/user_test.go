package main

import (
	"encoding/json"
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/tjfoc/gmsm/sm2"
)

type LoginRequest struct {
	Action    string `json:"action"`
	PubKey    string `json:"pubKey"`
	Signature string `json:"signature"`
}

func wsLogin(t *testing.T) {
	var err error
	priv, err = sm2.GenerateKey()
	if err != nil {
		log.Println("sm2 generate error: ", err)
		return
	}
	mockManager(priv)

	privateKey, publicKey := getKeyPairHex(priv)
	log.Printf("private key: %s \n public key: %s", privateKey, publicKey)

	//Signature
	r, s, err := sm2.Sign(priv, []byte(sessionID))
	if err != nil {
		log.Println("sm2 sign error: ", err)
		return
	}
	signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	log.Printf("signature: %s", signature)

	req := &LoginRequest{
		Action:    "login",
		PubKey:    publicKey,
		Signature: signature,
	}
	b, err := json.Marshal(req)
	if err != nil {
		log.Println("json encode error:", err)
		return
	}
	if err = c.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println("write error:", err)
		return
	}
	pubKey = publicKey
	log.Println("websocket request login")

	resp, err := getResponse(c)
	if err != nil {
		log.Println("response error:", err)
		return
	}
	if !strings.Contains(resp.Data, "NodeManager") {
		log.Println("login error:", resp.Data)
		return
	}
}

func mockManager(priv *sm2.PrivateKey) {
	mockPrivateKey := "4345aed948b0547e70ce4038c797680541fcd2f2d9a60e49f41891c6fcca6f69"
	mockPublicKey := "04a64b55e6575ea1bbb17cd7e2be739f3deff7ef47f5ca38b44f6d84aedc321a369757e150476967b45c9d8c82f8c9af8a2528dadb01a5d0669966300b2ea5d12f"

	d := new(big.Int)
	d.SetString(mockPrivateKey, 16)
	// log.Println("bigint:", d)
	x := new(big.Int)
	y := new(big.Int)
	x.SetString(mockPublicKey[2:66], 16)
	y.SetString(mockPublicKey[66:130], 16)
	// log.Println("x, y: ", x, y)

	priv.D = d
	priv.PublicKey.X = x
	priv.PublicKey.Y = y
}

func leftPad(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return strings.Repeat("0", n-len(s)) + s
}

func getKeyPairHex(priv *sm2.PrivateKey) (string, string) {
	privateKey := leftPad(priv.D.Text(16), 64)
	publicKey := "04" + leftPad(priv.PublicKey.X.Text(16), 64) + leftPad(priv.PublicKey.Y.Text(16), 64)
	return privateKey, publicKey
}
