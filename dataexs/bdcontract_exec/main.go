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

const (
	privKey = "d048300bd1645a7f4abe10329167bfa1a85a5d6b73562796d2abbfa34261e7e1"
	pubKey  = "04395d3c655637c4f06cbeabc1f061815a1e85fedd2c6f6d23afba9412e2bd8bf387786cbd659b588bca646ccc7076caf31923749db51a3f966c4e28594a9be0e5"
)

const contract = "Datassets"

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

type ListProjectsRequest struct {
	Action string `json:"action"`
}

type CreateProjectRequest struct {
	Action    string `json:"action"`
	IsPrivate bool   `json:"isPrivate"`
	IsFolder  bool   `json:"isFolder"`
	Dir       string `json:"dir"`
	Name      string `json:"name"`
	ProjTpl   string `json:"projectTemplate"`
}

type SaveFileRequest struct {
	Action    string `json:"action"`
	IsAppend  bool   `json:"isAppend"`
	IsPrivate bool   `json:"isPrivate"`
	Path      string `json:"path"`
	Content   string `json:"content"`
}

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

type Process struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"contractStatus"`
}

func main() {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:58150", Path: "/SCIDE/SCExecutor"}
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

	// List Projects
	pp, err := wsListProjects(c)
	if err != nil {
		log.Fatal("list projects error:", err)
	}
	var contractExist bool
	for _, p := range pp {
		if p == contract {
			contractExist = true
			break
		}
	}

	// Create Project
	if !contractExist {
		pp, err := wsCreateProject(c)
		if err != nil {
			log.Fatal("create project error:", err)
		}
		contractExist = false
		for _, p := range pp {
			if p == contract {
				contractExist = true
				break
			}
		}
		if !contractExist {
			log.Fatal("create project error:", err)
		}
		if err := UploadFile(privKey, pubKey, "postgres.yjs", "/Users/luxiaotong/code/go_practice/dataexs/bdcontract_exec/postgres.yjs"); err != nil {
			log.Fatal("upload postgres.yjs error:", err)
		}
		if err := UploadFile(privKey, pubKey, "postgresql-9.1-901-1.jdbc4.jar", "/Users/luxiaotong/code/go_practice/dataexs/bdcontract_exec/postgresql-9.1-901-1.jdbc4.jar"); err != nil {
			log.Fatal("upload postgresql-9.1-901-1.jdbc4.jar error:", err)
		}
		if err := UploadFile(privKey, pubKey, "postsample.jar", "/Users/luxiaotong/code/go_practice/dataexs/bdcontract_exec/postgres.yjs"); err != nil {
			log.Fatal("upload postgres.yjs error:", err)
		}
	}

	// Save Contract File
	if err := wsSaveFile(c); err != nil {
		log.Fatal("create project error:", err)
	}

	// Compile
	ypk, err := wsCompile(c)
	if err != nil {
		log.Fatal("compile contract error:", err)
	}

	// Start
	id, err := wsStartContractByYPK(c, privKey, pubKey, ypk)
	if err != nil {
		log.Fatal("start contract error:", err)
	}

	// Execute
	// {"host":"139.9.119.21","port":5432,"db":"target3","user":"test","pass":"datassets"}
	tdb := struct {
		Host string `json:"host"`
		Port uint32 `json:"port"`
		DB   string `json:"db"`
		User string `json:"user"`
		Pass string `json:"pass"`
	}{
		Host: "139.9.119.21",
		Port: 5432,
		DB:   "target3",
		User: "test",
		Pass: "datassets",
	}
	param, _ := json.Marshal(tdb)
	result, err := wsExecContract(c, privKey, pubKey, id, string(param))
	if err != nil {
		log.Fatal("exec contract error:", err)
	}
	log.Println("result: ", result)
}

func SignSM2KeyPairHex(privKey, data string) (string, error) {
	priv, err := x509.ReadPrivateKeyFromHex(privKey)
	if err != nil {
		log.Fatal("ReadPrivateKeyFromHex err:", err)
	}
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
	signature, err := SignSM2KeyPairHex(privKey, sessionID)
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
		Action: "main",
		Arg:    param,
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return "", err
	}
	argJSON := string(b)
	data := id + "|" + argJSON + "|" + pubKey
	signature, err := SignSM2KeyPairHex(privKey, data)
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

func wsListProjects(c *websocket.Conn) ([]string, error) {
	req := &ListProjectsRequest{
		Action: "listProjects",
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return nil, err
	}
	resp, err := getResponse(c)
	if err != nil {
		return nil, err
	}
	log.Println("list projects json: ", resp.Data)
	if len(resp.Data) == 0 {
		return nil, errors.Errorf("list projects error: %v", resp.Data)
	}
	var ret []string
	if err = json.Unmarshal([]byte(resp.Data), &ret); err != nil {
		return nil, err
	}
	log.Println("list projects array: ", ret)
	return ret, nil
}

func wsCreateProject(c *websocket.Conn) ([]string, error) {
	// {"isPrivate":false,"dir":"./","action":"createFile","isFolder":true,"name":"test","projectTemplate":"空白项目"}
	req := &CreateProjectRequest{
		Action:    "createFile",
		IsPrivate: false,
		IsFolder:  true,
		Dir:       "./",
		Name:      contract,
		ProjTpl:   "空白项目",
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return nil, err
	}
	resp, err := getResponse(c)
	if err != nil {
		return nil, err
	}
	log.Println("list created projects json: ", resp.Data)
	if len(resp.Data) == 0 {
		return nil, errors.Errorf("list created projects error: %v", resp.Data)
	}
	var ret []string
	if err = json.Unmarshal([]byte(resp.Data), &ret); err != nil {
		return nil, err
	}
	log.Println("list created projects array: ", ret)
	return ret, nil
}

func wsSaveFile(c *websocket.Conn) error {
	// {"isAppend":false,"isPrivate":false,"path":"/Datassets/Datassets.yjs","action":"saveFile","content":"contract Datassets{\naaaa\n}"}
	req := &SaveFileRequest{
		Action:    "saveFile",
		IsPrivate: false,
		IsAppend:  false,
		Path:      "/Datassets/Datassets.yjs",
		Content: `import "postgres.yjs";
		contract Data{
		  export function main(arg) {
			var args = JSON.parse(arg);
			var conn = PGConnect(args.host, args.port, args.db, args.user, args.pass);
			var  sql ="select count(1) from area_statistics where level=3";
			var statement = conn.createStatement();
			var resultSet = statement.executeQuery(sql);
			ret  = [];
			var meta = resultSet.getMetaData();
			for (;resultSet.next();){
			  var line = {
			  };
			  for (var j=1;j<=meta.getColumnCount();j++){
				line[meta.getColumnName(j)] = resultSet.getString(j);
			  }
			  ret.push(line);
			}
			return JSON.stringify(ret);
		  }
		}`,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return err
	}
	resp, err := getResponse(c)
	if err != nil {
		return err
	}
	if resp.Data != "successs" {
		return errors.Errorf("save file error: %v", resp.Data)
	}
	return nil
}

func wsCompile(c *websocket.Conn) (string, error) {
	req := &CompileRequest{
		Action:  "compile",
		Path:    contract,
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
		return "", errors.Errorf("wrong compile action: %v", resp.Action)
	}
	return resp.Result, nil
}

func wsStartContractByYPK(c *websocket.Conn, privKey, pubKey, ypk string) (string, error) {
	script := "empty"
	data := "Algorithm|" + script + "|" + pubKey
	signature, err := SignSM2KeyPairHex(privKey, data)
	if err != nil {
		return "", err
	}
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
		return "", err
	}
	if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
		return "", err
	}
	process := make([]*Process, 0, 1)
	for {
		resp, err := getResponse(c)
		if err != nil {
			return "", err
		}
		switch resp.Action {
		case "onListContractProcess":
			log.Printf("onListContractProcess in start contract response : %v", resp.Data)
			if err = json.Unmarshal([]byte(resp.Data), &process); err != nil {
				return "", err
			}
			if len(process) != 1 || process[0].Status != "RUNNING" {
				return "", errors.Errorf("listProcess error: %#v", process)
			}
			goto StartedByYPK
		case "onStartContract":
			log.Printf("onStartContract in start contract response : %v", resp.Data)
			var ret ContractResponse
			if err = json.Unmarshal([]byte(resp.Data), &ret); err != nil {
				return "", err
			}
			if ret.Status != "Success" {
				return "", errors.Errorf("startContract error: %v", resp)
			}
		}
	}
StartedByYPK:
	return process[0].ID, nil
}
