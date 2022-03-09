package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/tjfoc/gmsm/sm2"
)

type WSResponse struct {
	Action  string `json:"action"`
	Data    string `json:"data"`
	Session string `json:"session"`
	Status  bool   `json:"status"`
	Result  string `json:"result"`
}

var (
	c         *websocket.Conn
	sessionID string
	pubKey    string
	priv      *sm2.PrivateKey
)

func init() {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:58150", Path: "/SCIDE/SCExecutor"}
	var err error
	c, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
}

func getResponse(c *websocket.Conn) (*WSResponse, error) {
	_, message, err := c.ReadMessage()
	if err != nil {
		return nil, err
	}
	fmt.Printf("websocket response: %s\n", message)
	var resp WSResponse
	if err = json.Unmarshal(message, &resp); err != nil {
		return nil, err
	}
	// log.Debug("resp: %v", resp)
	return &resp, nil
}

func wsPing(t *testing.T) {
	if err := c.WriteMessage(websocket.TextMessage, []byte(`{"action":"ping"}`)); err != nil {
		log.Println("write error:", err)
		return
	}
	log.Println("websocket request ping")
	resp, _ := getResponse(c)
	log.Println("ping response:", resp)
}

func wsSession(t *testing.T) {
	if err := c.WriteMessage(websocket.TextMessage, []byte(`{"action":"getSessionID"}`)); err != nil {
		log.Println("write error:", err)
		return
	}
	log.Println("websocket request getSessionId")
	resp, err := getResponse(c)
	if err != nil {
		log.Println("get session error: ", err)
		return
	}
	if resp.Action != "onSessionID" {
		log.Println("wrong action response: ", resp.Action)
		return
	}
	sessionID = resp.Session
}

func TestAll(t *testing.T) {
	defer c.Close()
	t.Run("wsPing", wsPing)
	t.Run("wsSession", wsSession)
	t.Run("wsLogin", wsLogin)
	// t.Run("wsCompile", wsCompile)
	// t.Run("wsStartContractByYPK", wsStartContractByYPK)
	t.Run("wsQueryContractLogByDate", wsQueryContractLogByDate)
}
