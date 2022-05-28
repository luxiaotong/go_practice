package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tjfoc/gmsm/sm2"
)

var (
	sessionID string
	pubKey    string
	priv      *sm2.PrivateKey
)

type WSResponse struct {
	Action  string `json:"action"`
	Data    string `json:"data"`
	Session string `json:"session"`
}

type LoginRequest struct {
	Action    string `json:"action"`
	PubKey    string `json:"pubKey"`
	Signature string `json:"signature"`
}

type ApplyNodeRoleRequest struct {
	Action string `json:"action"`
	PubKey string `json:"pubKey"`
	Role   string `json:"role"`
}

type AuthNodeRoleRequest struct {
	Action string `json:"action"`
	PubKey string `json:"pubKey"`
	Accept bool   `json:"isAccept"`
}

type Contract struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"contractStatus"`
	Function string `json:"function"`
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

type ExecuteContractRequest struct {
	Action     string `json:"action"`
	PubKey     string `json:"owner"`
	RequestID  string `json:"requestID"`
	ContractID string `json:"contractID"`
	Arg        string `json:"arg"`
	Signature  string `json:"signature"`
}

type KillContractRequest struct {
	Action     string `json:"action"`
	RequestID  string `json:"requestID"`
	ContractID string `json:"id"`
}

func main() {
	// httpPing()
	wsConn()
}

func wsConn() {
	done := make(chan int)
	sessDone := make(chan int)
	loginDone := make(chan int)
	startDone := make(chan int)
	execDone := make(chan int)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	contract := &Contract{
		Name: "PostgreSQLSample",
	}

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/SCIDE/SCExecutor"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer c.Close()

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			// log.Printf("websocket response: %s", message)
			var resp WSResponse
			if err = json.Unmarshal(message, &resp); err != nil {
				log.Println("json decode error:", err)
				return
			}
			log.Printf("resp: %v", resp)
			if resp.Action == "onSessionID" {
				// TODO merge sessionID and sessDone
				sessionID = resp.Session
				sessDone <- 1
			}
			if resp.Action == "onLogin" {
				loginDone <- 1
			}
			if resp.Action == "onListContractProcess" {
				var data []*Contract
				if err = json.Unmarshal([]byte(resp.Data), &data); err != nil {
					log.Println("json decode data error: ", err)
					return
				}
				for _, d := range data {
					if d.Name == contract.Name && d.Status == "Ready" {
						contract.Status = d.Status
						contract.ID = d.ID
						startDone <- 1
					}
				}
			}
			if resp.Action == "onExecuteResult" {
				execDone <- 1
			}
		}
	}()

	wsPing(c)
	wsSession(c)

	for {
		select {
		case <-done:
			return
		case <-sessDone:
			log.Println("session id: ", sessionID)
			wsLogin(c, sessionID)
		case <-loginDone:
			wsApplyNodeRole(c, "ContractProvider")
			wsApplyNodeRole(c, "ContractUser")
			wsApplyNodeRole(c, "ContractInstanceManager")
			wsAuthNodeRole(c)
			wsStartContract(c, contract)
		case <-startDone:
			wsExecContract(c, contract)
		case <-execDone:
			wsKillContract(c, contract)
			return
		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			if err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
				log.Println("write close:", err)
			}
			return
		}
	}
}

func wsPing(c *websocket.Conn) {
	if err := c.WriteMessage(websocket.TextMessage, []byte(`{"action":"ping"}`)); err != nil {
		log.Println("write error:", err)
		return
	}
	log.Println("websocket request ping")
}

func wsSession(c *websocket.Conn) {
	if err := c.WriteMessage(websocket.TextMessage, []byte(`{"action":"getSessionID"}`)); err != nil {
		log.Println("write error:", err)
		return
	}
	log.Println("websocket request getSessionId")
}

func wsLogin(c *websocket.Conn, sessionID string) {
	var err error
	priv, err = sm2.GenerateKey(nil)
	if err != nil {
		log.Println("sm2 generate error: ", err)
		return
	}
	// TODO: DEL
	mockManager(priv)

	privateKey, publicKey := getKeyPairHex(priv)
	log.Printf("private key: %s \n public key: %s", privateKey, publicKey)

	//Signature
	// r, s, err := sm2.Sign(priv, []byte(sessionID))
	// if err != nil {
	// 	log.Println("sm2 sign error: ", err)
	// 	return
	// }
	// signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	s, err := priv.Sign(nil, []byte(sessionID), nil)
	if err != nil {
		log.Println("sm2 sign error: ", err)
		return
	}
	signature := hex.EncodeToString(s)
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

func mockManager(priv *sm2.PrivateKey) {
	mockPrivateKey := "930d8bbb78a3a52cea4e534967b1b5d7599f18dc5b8496da8711f163bcfce50b"
	mockPublicKey := "04d46cbfa3fae734906099b9b84a385f15592a1d405b74589c711f706c41ab036d7cb5434370e623ed1cfafa8cc9263838177c7b30a7523418d9b50d9df8214d48"

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

// wsApplyNodeRole apply node role
//  param: *websocket.Conn c
//  param: string role "ContractProvider", "ContractUser", "ContractInstanceManager"
//  return
func wsApplyNodeRole(c *websocket.Conn, role string) {
	req := &ApplyNodeRoleRequest{
		Action: "applyNodeRole",
		PubKey: pubKey,
		Role:   role,
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
	log.Println("websocket request applyNodeRole: ", role)
}

func wsAuthNodeRole(c *websocket.Conn) {
	req := &AuthNodeRoleRequest{
		Action: "authNodeRole",
		PubKey: pubKey,
		Accept: true,
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
	log.Println("websocket request authNodeRole")
}

func wsStartContract(c *websocket.Conn, contract *Contract) {
	script := "empty"
	//Signature
	// r, s, err := sm2.Sign(priv, []byte("Algorithm|"+script+"|"+pubKey))
	// if err != nil {
	// 	log.Println("sign contract error: ", err)
	// 	return
	// }
	// signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	s, err := priv.Sign(nil, []byte("Algorithm|"+script+"|"+pubKey), nil)
	if err != nil {
		log.Println("sign contract error: ", err)
		return
	}
	signature := hex.EncodeToString(s)
	log.Printf("signature: %s", signature)
	req := &StartContractRequest{
		Action:    "startContract",
		PubKey:    pubKey,
		RequestID: fmt.Sprintf("%d", time.Now().Unix()),
		Private:   false,
		Path:      fmt.Sprintf("/%s/manifest.json", contract.Name),
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
	log.Println("websocket request startContract")
}

func wsExecContract(c *websocket.Conn, contract *Contract) {
	arg := `{"action":"selectStudent", "arg":""}`
	//Signature
	// r, s, err := sm2.Sign(priv, []byte(contract.ID+"|"+arg+"|"+pubKey))
	// if err != nil {
	// 	log.Println("sign contract arg error: ", err)
	// 	return
	// }
	// signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	s, err := priv.Sign(nil, []byte(contract.ID+"|"+arg+"|"+pubKey), nil)
	if err != nil {
		log.Println("sign contract error: ", err)
		return
	}
	signature := hex.EncodeToString(s)
	log.Printf("signature: %s", signature)
	req := &ExecuteContractRequest{
		Action:     "executeContract",
		PubKey:     pubKey,
		RequestID:  fmt.Sprintf("%d", time.Now().Unix()),
		ContractID: contract.ID,
		Arg:        arg,
		Signature:  signature,
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
	log.Println("websocket request executeContract")
}

func wsKillContract(c *websocket.Conn, contract *Contract) {
	req := &KillContractRequest{
		Action:     "killContractProcess",
		RequestID:  fmt.Sprintf("%d", time.Now().Unix()),
		ContractID: contract.ID,
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
	log.Println("websocket request killContract")
}

// func heartbeat() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()
// 	for range ticker.C {
// 		// TODO
// 	}
// }

// func httpPing() {
// 	resp, err := http.Get("http://127.0.0.1:8080/SCIDE/SCManager?action=ping")
// 	if err != nil {
// 		log.Printf("ping error: %v", err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("read error: %v", err)
// 	}
// 	log.Printf("ping response: %v", string(body))
// }
