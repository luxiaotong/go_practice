package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tjfoc/gmsm/sm2"
)

type CompileRequest struct {
	Action  string `json:"action"`
	Path    string `json:"path"`
	Private bool   `json:"isPrivate"`
}

type StartContractRequest struct {
	Action    string `json:"action"`
	PubKey    string `json:"owner"`
	RequestID string `json:"requestID"`
	Private   bool   `json:"isPrivate"`
	Path      string `json:"path"`
	Script    string `json:"script"`
	Signature string `json:"signature"`
}

type QueryContractLogByDateRequest struct {
	Action string `json:"action"`
	Start  int64  `json:"start"`
}

var ypk string

func wsCompile(t *testing.T) {
	req := &CompileRequest{
		Action:  "compile",
		Path:    "Datassets",
		Private: false,
	}
	b, err := json.Marshal(req)
	if err != nil {
		log.Println("json encode error:", err)
		return
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println("write error:", err)
		return
	}
	log.Println("websocket request startContract")

	resp, err := getResponse(c)
	if err != nil {
		log.Println("response error:", err)
		return
	}
	if resp.Action != "onCompile" {
		log.Println("wrong action response: ", resp.Action)
		return
	}
	log.Println("compile success: ", resp)
	ypk = resp.Result
}

func wsStartContractByYPK(t *testing.T) {
	script := "empty"
	//Signature
	r, s, err := sm2.Sign(priv, []byte("Algorithm|"+script+"|"+pubKey))
	if err != nil {
		log.Println("sign contract error: ", err)
		return
	}
	signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	log.Printf("signature: %s", signature)
	req := &StartContractRequest{
		Action:    "startContractByYPK",
		PubKey:    pubKey,
		RequestID: fmt.Sprintf("%d", time.Now().Unix()),
		Private:   false,
		Path:      "/" + ypk,
		Script:    script,
		Signature: signature,
	}
	b, err := json.Marshal(req)
	if err != nil {
		log.Println("json encode error:", err)
		return
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println("write error:", err)
		return
	}
	log.Println("websocket request startContractByYPK")

	for {
		resp, err := getResponse(c)
		if err != nil {
			log.Println("response error:", err)
			return
		}
		if resp.Action == "onListContractProcess" {
			log.Println("contract process: ", resp)
		}
		if resp.Action == "onStartContract" {
			log.Println("start contract by ypk response: ", resp)
		}
	}
}

func wsQueryContractLogByDate(t *testing.T) {
	req := &QueryContractLogByDateRequest{
		Action: "queryContractLogByDate",
		Start:  int64(1646644863531),
	}
	b, err := json.Marshal(req)
	if err != nil {
		log.Println("json encode error:", err)
		return
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println("write error:", err)
		return
	}
	log.Println("websocket request queryContractLogByDate")

	resp, err := getResponse(c)
	if err != nil {
		log.Println("response error:", err)
		return
	}
	if resp.Action != "onQueryContractLogByDate" {
		log.Println("wrong action response: ", resp.Action)
		return
	}
	log.Println("onQueryContractLogByDate success: ", resp)
}
