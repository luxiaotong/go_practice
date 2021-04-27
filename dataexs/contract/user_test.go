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
	mockPrivateKey := "2e4cbecd9737b8123cad0ebe80d4bcc88dcf0a39d9a6e6b5f789d5fd916a634a"
	mockPublicKey := "041e032bdec78467dc51a4755a33b80bd75718300a5da6e464f13c3dc6c25bf7d0bba4bade32bc81edd1522f3ec5e0e2a44f711593302dd8c9247dc34342e93384"

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
