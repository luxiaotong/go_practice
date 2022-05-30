package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/x509"
)

const privKey = "229160f61c36a22b3ac5ea200c634aa4b8c1b47b5997f971f0270e60e6536968"
const pubKey = "043dc667a68073f08f570df96e1eddd34948dfa5812618c6f25e383e3e003911b36e6a78112ed9e503e379b382816c8995b8add4bdc9bda28b59ceb4a577079f1d"

type SessionRequest struct {
	Action string `json:"action"`
}

type LoginRequest struct {
	Action    string `json:"action"`
	PubKey    string `json:"pubKey"`
	Signature string `json:"signature"`
}

type WSResponse struct {
	Action  string `json:"action"`
	Data    string `json:"data"`
	Session string `json:"session"`
	Result  string `json:"result"`
}

type ContractArg struct {
	Action string `json:"action"`
	Arg    string `json:"arg"`
}

type ExecuteContractRequest struct {
	Action     string `json:"action"`
	PubKey     string `json:"owner"`
	RequestID  string `json:"requestID"`
	ContractID string `json:"contractID"`
	Arg        string `json:"arg"`
	Signature  string `json:"signature"`
}

type ContractResponse struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func main() {
	// u := url.URL{Scheme: "ws", Host: "139.9.119.21:58121", Path: "/SCIDE/SCExecutor"}
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:21030", Path: "/SCIDE/SCExecutor"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer c.Close()

	sessionID, err := wsSession(c)
	if err != nil {
		log.Fatal("session error:", err)
	}
	if err := wsLogin(c, privKey, pubKey, sessionID); err != nil {
		log.Fatal("login error:", err)
	}
	// id := "TestLicense"
	// param := "http://www.baidu.com"
	// result, err := wsExecContract(c, privKey, pubKey, id, param)
	// if err != nil {
	// 	log.Fatal("exec contract error:", err)
	// }
	// log.Println("result: ", result)
}

// nolint: unused, deadcode
func leftPad(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return strings.Repeat("0", n-len(s)) + s
}

func SignSM2KeyPairHex(privKey, pubKey, data string) (string, error) {
	priv, err := x509.ReadPrivateKeyFromHex(privKey)
	if err != nil {
		log.Fatal("ReadPrivateKeyFromHex err:", err)
	}

	// Sign #1
	// r, s, err := Sign(priv, []byte(data))
	// if err != nil {
	// 	return "", err
	// }
	// signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	// log.Println("s1: ", signature)

	// Sign #2
	// r, s, err := sm2.Sm2Sign(priv, []byte(data), nil, nil)
	// if err != nil {
	// 	return "", err
	// }
	// s2 := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	// signature := s2

	// Sign #3
	s3, err := priv.Sign(nil, []byte(data), nil)
	if err != nil {
		return "", err
	}
	signature := hex.EncodeToString(s3)
	log.Println("s3: ", signature)

	return signature, nil
}

func getResponse(c *websocket.Conn) (*WSResponse, error) {
	_, message, err := c.ReadMessage()
	if err != nil {
		return nil, err
	}
	log.Println("websocket response: ", string(message))
	var resp WSResponse
	if err = json.Unmarshal(message, &resp); err != nil {
		return nil, err
	}
	// log.Println("resp: %v", resp)
	return &resp, nil
}

func wsSession(c *websocket.Conn) (string, error) {
	req := &SessionRequest{
		Action: "getSessionID",
	}
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return "", err
	}
	resp, err := getResponse(c)
	if err != nil {
		return "", err
	}
	if resp.Action != "onSessionID" {
		return "", errors.Errorf("wrong action response: %v", resp.Action)
	}
	return resp.Session, nil
}

func wsLogin(c *websocket.Conn, privKey, pubKey, sessionID string) error {
	signature, err := SignSM2KeyPairHex(privKey, pubKey, sessionID)
	if err != nil {
		return err
	}
	req := &LoginRequest{
		Action:    "login",
		PubKey:    pubKey,
		Signature: signature,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	if err = c.WriteMessage(websocket.TextMessage, b); err != nil {
		return err
	}
	resp, err := getResponse(c)
	if err != nil {
		return err
	}
	if !strings.Contains(resp.Data, "NodeManager") {
		return errors.Errorf("login contract error: %v", resp.Data)
	}
	return nil
}

// nolint: unused,deadcode
func wsExecContract(c *websocket.Conn, privKey, pubKey, id, param string) (string, error) {
	arg := &ContractArg{
		Action: "get",
		Arg:    param,
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return "", err
	}
	argJSON := string(b)
	data := id + "|" + argJSON + "|" + pubKey
	signature, err := SignSM2KeyPairHex(privKey, pubKey, data)
	if err != nil {
		return "", err
	}
	req := &ExecuteContractRequest{
		Action:     "executeContract",
		PubKey:     pubKey,
		RequestID:  fmt.Sprintf("%d", time.Now().Unix()),
		ContractID: id,
		Arg:        argJSON,
		Signature:  signature,
	}
	b, err = json.Marshal(req)
	if err != nil {
		return "", err
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return "", err
	}
	resp, err := getResponse(c)
	if err != nil {
		return "", err
	}
	var ret ContractResponse
	if err = json.Unmarshal([]byte(resp.Data), &ret); err != nil {
		return "", err
	}
	if ret.Status != "Success" {
		return "", errors.Errorf("execute contract error: %v", ret.Result)
	}
	return ret.Result, nil
}
