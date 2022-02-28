package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const host = "http://139.9.119.21:58131"

const (
	sendTxAPI = "/v0/ledgers/test/transactions"
	getTxAPI  = "/v0/ledgers/test/transaction"
)

type transaction struct {
	Type  int32  `json:"type"`
	From  string `json:"from"`
	Data  string `json:"data"`
	Nonce int32  `json:"nonce"`
}

type sendTxRequest struct {
	Tx *transaction `json:"transaction"`
}

type sendTxResponse struct {
	Hash string `json:"hash"`
}

func main() {
	hash := sendTx()
	fmt.Println("send time: ", time.Now().Format("2006-01-02 15:04:05"))
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	quit := time.After(60 * time.Second)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("tick time: ", t.Format("2006-01-02 15:04:05"))
			getTx(hash)
		case t := <-quit:
			fmt.Println("quit time: ", t.Format("2006-01-02 15:04:05"))
			return
		}
	}
	// for c := range ticker.C {
	// 	fmt.Println("time: ", c.Format("2006-01-02 15:04:05"))
	// 	getTx(hash)
	// }
}

func sendTx() string {
	u := host + sendTxAPI
	now := time.Now().Format("2006-01-02 15:04:05")
	req := &sendTxRequest{
		Tx: &transaction{
			Type:  0,
			From:  "8A3K/vANyv7wDcr+8A3K/vANyv4=",
			Data:  base64.StdEncoding.EncodeToString([]byte(now)),
			Nonce: 0,
		},
	}
	b, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(b)
	fmt.Println("send tx request: ", buf.String())
	// buf := bytes.NewBufferString(`{"transaction":{"type":0,"from":"8A3K/vANyv7wDcr+8A3K/vANyv4=","nonce":52,"data":"lQItWZKS5hlUn6V/DMKKwvZXxvM="}}`)
	resp, err := http.Post(u, "application/json", buf)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var bout bytes.Buffer
	if _, err := io.Copy(&bout, resp.Body); err != nil {
		panic(err)
	}
	var res sendTxResponse
	if err := json.Unmarshal(bout.Bytes(), &res); err != nil {
		panic(err)
	}
	fmt.Println("tx hash: ", res.Hash)
	return res.Hash
}

func getTx(hash string) {
	u := host + getTxAPI + "?hash=" + url.QueryEscape(hash)
	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var bout bytes.Buffer
	if _, err := io.Copy(&bout, resp.Body); err != nil {
		panic(err)
	}
	fmt.Println("get tx response: ", bout.String())
}
