package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
)

var mockPair = map[string]map[string]string{
	"node": {
		"private": "9b495adae2d43dd8c1041709de972906a0c9773583ca36a4339cce42201acc83",
		"public":  "04e50e6f0e821534f8ac20ea960b6c7fa318f569b5350332e94e8205f774120be4745dd2c5daaa405a6c89acdfc3d2889c86824e21baa1ae61468821e5611d7e3e",
	},
	"node2": {
		"private": "ccb323cdb9a8ea971d86c756c15d9ea6c42170672c49aaae7760415f56151597",
		"public":  "046cbf152d057e326193e4740db4cfa1159031c82d183d71fb3a19b1afb2258a49d09922ff9ad06d70b3460a00c51955a3be538569dec1dc070c73cb13589482af",
	},
	"center": {
		"private": "29e3b93b58c17f02648bd0d4d95afc00b917a604d55f5c616c53af636f52940d",
		"public":  "04c9229cf5837ee369e8432d55e4611e344eda324990f035273835b0c9ba072662ea3763d35e30bf1f337f95081c790bda6f1764465ec7dc587f8f667ef1d5ea67",
	},
}

var (
	priv *sm2.PrivateKey
)

type WSResponse struct {
	Action   string `json:"action"`
	Session  string `json:"session"`
	Data     string `json:"data"`
	Result   string `json:"result"`
	Progress string `json:"progress"`
}

type LoginRequest struct {
	Action    string `json:"action"`
	PubKey    string `json:"pubKey"`
	Signature string `json:"signature"`
}

type CompileRequest struct {
	Action  string `json:"action"`
	Path    string `json:"path"`
	Private bool   `json:"privateTab"`
}

type ListTrustUnitsRequest struct {
	Action string `json:"action"`
	PubKey string `json:"pubKey"`
}

type Node struct {
	Name   string `json:"nodeName"`
	PubKey string `json:"pubKey"`
}

type TrustUnit struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

type ListTrustUnitsResponse struct {
	Action string       `json:"action"`
	Data   []*TrustUnit `json:"data"`
}

type LoadNodeConfigRequest struct {
	Action string `json:"action"`
}

type LoadNodeConfigResponse struct {
	Action string `json:"action"`
	Data   *Node  `json:"data"`
}

type DistributeRequest struct {
	Action    string `json:"action"`
	PeersID   string `json:"peersID"`
	Project   string `json:"projectName"`
	Private   bool   `json:"isPrivate"`
	Sponsor   string `json:"sponsorName"`
	Signature string `json:"signature"`
}

type StartContractP2PRequest struct {
	Action    string `json:"action"`
	PeersID   string `json:"peersID"`
	Project   string `json:"projectName"`
	Private   bool   `json:"isPrivate"`
	Type      string `json:"type"`
	Sponsor   string `json:"sponsorPeerID"`
	Signature string `json:"signature"`
}

type StartContractP2PTrustfullyRequest struct {
	Action    string `json:"action"`
	PeersID   string `json:"peersID"`
	Private   bool   `json:"isPrivate"`
	Signature string `json:"signature"`
	Path      string `json:"path"`
	Owner     string `json:"owner"`
}

type ExecuteContractRequest struct {
	Action     string `json:"action"`
	ContractID string `json:"contractID"`
	Operation  string `json:"operation"`
	Arg        string `json:"arg"`
	PubKey     string `json:"pubkey"`
	Signature  string `json:"signature"`
}

type ListContractRequest struct {
	Action string `json:"action"`
}

type Contract struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	u := url.URL{Scheme: "ws", Host: "139.9.119.21:58119", Path: "/NodeCenterWS"}
	centerC, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer centerC.Close()

	centerSID, err := wsSession(centerC)
	if err != nil {
		log.Fatal("center session error:", err)
	}
	log.Println("center session id: ", centerSID)

	if err := wsLogin(centerC, "center", centerSID); err != nil {
		log.Fatal("center login error:", err)
	}

	names, peers, err := wsListTrustUnits(centerC)
	if err != nil {
		log.Fatal("list trust units error:", err)
	}
	log.Println("names: ", names)
	log.Println("peers: ", peers)

	u = url.URL{Scheme: "ws", Host: "139.9.119.21:58121", Path: "/SCIDE/SCExecutor"}
	nodeC, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer nodeC.Close()

	if err := wsPing(nodeC); err != nil {
		log.Fatal("ping error:", err)
	}

	nodeSID, err := wsSession(nodeC)
	if err != nil {
		log.Fatal("node session error:", err)
	}
	log.Println("node session id: ", nodeSID)

	if err := wsLogin(nodeC, "node", nodeSID); err != nil {
		log.Fatal("node login error:", err)
	}

	// node, err := wsLoadNodeConfig(nodeC)
	// if err != nil {
	// 	log.Fatal("load node config error:", err)
	// }

	// ypk, err := wsCompile(nodeC)
	// if err != nil {
	// 	log.Fatal("compile error:", err)
	// }
	ypk := "AAStorage_2020-12-04-05:32:37.ypk"
	log.Println("ypk: ", ypk)

	// others := make([]string, 0)
	// for _, n := range names {
	// 	if n != node.Name {
	// 		others = append(others, n)
	// 	}
	// }
	// othersName := strings.Join(others, ",")
	// if err := wsDistribute(centerC, othersName, ypk, node.Name); err != nil {
	// 	log.Fatal("distribute contract error:", err)
	// }

	u = url.URL{Scheme: "ws", Host: "139.9.119.21:58123", Path: "/SCIDE/SCExecutor"}
	node2C, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer node2C.Close()
	node2SID, err := wsSession(node2C)
	if err != nil {
		log.Fatal("node2 session error:", err)
	}
	log.Println("node2 session id: ", node2SID)
	if err := wsLogin(node2C, "node2", node2SID); err != nil {
		log.Fatal("node2 login error:", err)
	}

	if err := wsStartContractMultiPoint(node2C, "node2", ypk, peers); err != nil {
		log.Fatal("start contract multi point error:", err)
	}
	// if err := wsStartContractP2PTrustfully(nodeC, "node", ypk, peers); err != nil {
	// 	log.Fatal("start contract multi point error:", err)
	// }

	// contractID, err := wsListContract(node2C)
	// if err != nil {
	// 	log.Fatal("list contract error:", err)
	// }
	// arg := `{"from":"191377654ab2e7ab09d58bf3627a241ed78ece82","to":"","order_id":"1324968619020390400","buyer":"1289495985193488384","seller":"1281908739774877696"}`
	// if err := wsExecuteContract(node2C, "node2", contractID, "setOrder", arg); err != nil {
	// 	log.Fatal("execute contract setOrder error:", err)
	// }
	// arg = `{"from":"191377654ab2e7ab09d58bf3627a241ed78ece82","to":"","order_id":"1324968619020390400"}`
	// if err := wsExecuteContract(nodeC, "node", contractID, "getOrder", arg); err != nil {
	// 	log.Fatal("execute contract getOrder error:", err)
	// }
}

func wsPing(c *websocket.Conn) error {
	if err := c.WriteMessage(websocket.TextMessage, []byte(`{"action":"ping"}`)); err != nil {
		return err
	}
	_, message, err := c.ReadMessage()
	if err != nil {
		return err
	}
	log.Println("ping response: ", string(message))
	return nil
}

func getResponse(c *websocket.Conn) (*WSResponse, error) {
	_, message, err := c.ReadMessage()
	if err != nil {
		return nil, err
	}
	// log.Println("websocket response: ", string(message))
	var resp WSResponse
	if err = json.Unmarshal(message, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func wsSession(c *websocket.Conn) (string, error) {
	if err := c.WriteMessage(websocket.TextMessage, []byte(`{"action":"getSessionID"}`)); err != nil {
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

func wsLogin(c *websocket.Conn, role, sessionID string) error {
	var err error
	priv, err = sm2.GenerateKey(nil)
	if err != nil {
		return err
	}
	mockManager(priv, role)
	privateKey, publicKey := getKeyPairHex(priv)
	log.Println("private key: ", privateKey)
	log.Println("public key: ", publicKey)

	//Signature
	// r, s, err := sm2.Sign(priv, []byte(sessionID))
	// if err != nil {
	// 	return err
	// }
	// signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	s, err := priv.Sign(nil, []byte(sessionID), nil)
	if err != nil {
		return err
	}
	signature := hex.EncodeToString(s)
	// log.Println("signature: ", signature)

	req := &LoginRequest{
		Action:    "login",
		PubKey:    publicKey,
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
	if resp.Action != "onLogin" {
		return errors.Errorf("wrong action response: %v", resp.Action)
	}
	return nil
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

func mockManager(priv *sm2.PrivateKey, role string) {
	mockPrivateKey := mockPair[role]["private"]
	mockPublicKey := mockPair[role]["public"]

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

func wsCompile(c *websocket.Conn) (string, error) {
	req := &CompileRequest{
		Action:  "compile",
		Path:    "AAStorage",
		Private: false,
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
	if resp.Action != "onCompile" {
		return "", errors.Errorf("wrong action response: %v", resp.Action)
	}
	return resp.Result, nil
}

func wsListTrustUnits(c *websocket.Conn) ([]string, []string, error) {
	_, pubKey := getKeyPairHex(priv)
	req := &ListTrustUnitsRequest{
		Action: "listTrustUnits",
		PubKey: pubKey,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return nil, nil, err
	}
	_, message, err := c.ReadMessage()
	if err != nil {
		return nil, nil, err
	}
	// log.Println("websocket response: ", string(message))
	var resp ListTrustUnitsResponse
	if err = json.Unmarshal(message, &resp); err != nil {
		return nil, nil, err
	}
	if resp.Action != "onListTrustUnits" {
		return nil, nil, errors.Errorf("wrong action response: %v", resp.Action)
	}
	var val string
	for _, d := range resp.Data {
		if strings.HasSuffix(d.Key, "cluster012") {
			val = d.Val
			break
		}
	}
	var nn []*Node
	if err := json.Unmarshal([]byte(val), &nn); err != nil {
		return nil, nil, err
	}
	names := make([]string, 0)
	peers := make([]string, 0)
	for _, n := range nn {
		names = append(names, n.Name)
		peers = append(peers, n.PubKey)
	}
	return names, peers, nil
}

func wsLoadNodeConfig(c *websocket.Conn) (*Node, error) {
	req := &LoadNodeConfigRequest{
		Action: "loadNodeConfig",
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return nil, err
	}
	_, message, err := c.ReadMessage()
	if err != nil {
		return nil, err
	}
	log.Println("websocket response: ", string(message))
	var resp LoadNodeConfigResponse
	if err = json.Unmarshal(message, &resp); err != nil {
		return nil, err
	}
	if resp.Action != "onLoadNodeConfig" {
		return nil, errors.Errorf("wrong action response: %v", resp.Action)
	}
	return resp.Data, nil
}

func wsDistribute(c *websocket.Conn, others, project, sponsor string) error {
	mockManager(priv, "center")
	_, publicKey := getKeyPairHex(priv)
	signStr := fmt.Sprintf("DistributeContract|%s|%s", project, publicKey)
	// r, s, err := sm2.Sign(priv, []byte(signStr))
	// if err != nil {
	// 	return err
	// }
	// signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	s, err := priv.Sign(nil, []byte(signStr), nil)
	if err != nil {
		return err
	}
	signature := hex.EncodeToString(s)

	req := &DistributeRequest{
		Action:    "distributeContract",
		PeersID:   others,
		Project:   project,
		Private:   false,
		Sponsor:   sponsor,
		Signature: signature,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return err
	}
	for {
		resp, err := getResponse(c)
		if err != nil {
			return err
		}
		log.Println("distribute response: ", resp)
		if resp.Action == "onDistributeFinish" {
			break
		}
	}
	return nil
}

func wsStartContractMultiPoint(c *websocket.Conn, role, project string, peers []string) error {
	mockManager(priv, role)
	_, publicKey := getKeyPairHex(priv)
	signStr := fmt.Sprintf("Trusted|%s|%s", project, publicKey)
	// r, s, err := sm2.Sign(priv, []byte(signStr))
	// if err != nil {
	// 	return err
	// }
	// signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	s, err := priv.Sign(nil, []byte(signStr), nil)
	if err != nil {
		return err
	}
	signature := hex.EncodeToString(s)

	peersID := strings.Join(peers, ",")
	req := &StartContractP2PRequest{
		Action:    "startContractMultiPoint",
		PeersID:   peersID,
		Project:   project,
		Private:   false,
		Type:      "RequestAllResponseAll",
		Sponsor:   "",
		Signature: signature,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	log.Println("request: ", string(b))
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return err
	}
	for i := 0; i < len(peers)+1; i++ {
		resp, err := getResponse(c)
		if err != nil {
			return err
		}
		log.Println("start contract multi point response: ", resp)
	}
	return nil
}

func wsStartContractP2PTrustfully(c *websocket.Conn, role, ypk string, peers []string) error {
	path := "/AAStorage/" + ypk
	mockManager(priv, role)
	_, publicKey := getKeyPairHex(priv)
	signStr := fmt.Sprintf("Trusted|%s|%s", path, publicKey)
	// r, s, err := sm2.Sign(priv, []byte(signStr))
	// if err != nil {
	// 	return err
	// }
	// signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	s, err := priv.Sign(nil, []byte(signStr), nil)
	if err != nil {
		return err
	}
	signature := hex.EncodeToString(s)

	peersID := strings.Join(peers, ",")
	req := &StartContractP2PTrustfullyRequest{
		Action:    "startContractP2PTrustfully",
		PeersID:   peersID,
		Path:      path,
		Private:   false,
		Signature: signature,
		Owner:     mockPair[role]["public"],
	}
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	log.Println("request: ", string(b))
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return err
	}
	_, message, err := c.ReadMessage()
	if err != nil {
		return err
	}
	log.Println("start contract p2p trustfully response: ", message)
	return nil
}

func wsListContract(c *websocket.Conn) (string, error) {
	req := &ListContractRequest{
		Action: "listContractProcess",
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
	if resp.Action != "onListContractProcess" {
		return "", errors.Errorf("wrong action response: %v", resp.Action)
	}
	log.Println("list contract response: ", resp.Data)
	var cc []*Contract
	if err = json.Unmarshal([]byte(resp.Data), &cc); err != nil {
		return "", err
	}
	for _, c := range cc {
		log.Println("contract: ", c)
	}
	return cc[0].ID, nil
}

func wsExecuteContract(c *websocket.Conn, role, contractID, operation, arg string) error {
	mockManager(priv, role)
	_, publicKey := getKeyPairHex(priv)
	signStr := fmt.Sprintf("%s|%s|%s|%s", contractID, operation, arg, publicKey)
	// r, s, err := sm2.Sign(priv, []byte(signStr))
	// if err != nil {
	// 	return err
	// }
	// signature := leftPad(r.Text(16), 64) + leftPad(s.Text(16), 64)
	s, err := priv.Sign(nil, []byte(signStr), nil)
	if err != nil {
		return err
	}
	signature := hex.EncodeToString(s)

	req := &ExecuteContractRequest{
		Action:     "executeContract",
		ContractID: contractID,
		Operation:  operation,
		Arg:        arg,
		PubKey:     mockPair[role]["public"],
		Signature:  signature,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	// log.Println("request: ", string(b))
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return err
	}
	resp, err := getResponse(c)
	if err != nil {
		return err
	}
	if resp.Action != "onExecuteResult" {
		return errors.Errorf("wrong action response: %v", resp.Action)
	}
	// log.Printf("execute contract response: %#v", resp)
	var data WSResponse
	if err = json.Unmarshal([]byte(resp.Data), &data); err != nil {
		return err
	}
	log.Println("execute contract response: ", data.Result)
	return nil
}
