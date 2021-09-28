package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
)

const privKey = "9b495adae2d43dd8c1041709de972906a0c9773583ca36a4339cce42201acc83"
const pubKey = "04e50e6f0e821534f8ac20ea960b6c7fa318f569b5350332e94e8205f774120be4745dd2c5daaa405a6c89acdfc3d2889c86824e21baa1ae61468821e5611d7e3e"

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
	u := url.URL{Scheme: "ws", Host: "139.9.119.21:58121", Path: "/SCIDE/SCExecutor"}
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
	id := "TestLicense"
	param := "http://www.baidu.com"
	result, err := wsExecContract(c, privKey, pubKey, id, param)
	if err != nil {
		log.Fatal("exec contract error:", err)
	}
	log.Println("result: ", result)
}

func leftPad(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return strings.Repeat("0", n-len(s)) + s
}

func SignSM2KeyPairHex(privKey, pubKey, data string) (string, error) {
	priv, err := sm2.GenerateKey()
	if err != nil {
		return "", err
	}
	d := new(big.Int)
	d.SetString(privKey, 16)
	// log.Println("bigint:", d)
	x := new(big.Int)
	y := new(big.Int)
	x.SetString(pubKey[2:66], 16)
	y.SetString(pubKey[66:130], 16)
	// log.Println("x, y: ", x, y)
	priv.D = d
	priv.PublicKey.X = x
	priv.PublicKey.Y = y

	r, s, err := sm2.Sign(priv, []byte(data))
	if err != nil {
		return "", err
	}
	signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
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
