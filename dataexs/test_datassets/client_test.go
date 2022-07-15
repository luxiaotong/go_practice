package testdatassets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
)

const stepDone = 3

var clientSession string

type clientLoginRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
	IsLocal  bool   `json:"is_local"`
}

type clientLoginResponse struct {
	Err     string `json:"err_msg"`
	Session string `json:"session"`
	Updated int32  `json:"updated"`
}

type dbConnect struct {
	Type int32  `json:"type"`
	Host string `json:"host"`
	Port int32  `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Db   string `json:"db"`
}

type dbMessage struct {
	Src *dbConnect `json:"source_db"`
	Dst *dbConnect `json:"target_db"`
}

type setDBConnRequest struct {
	Session string     `json:"session"`
	Info    *dbMessage `json:"info"`
}

type opRequest struct {
	Session string `json:"session"`
}

type fieldMapping struct {
	FieldName   string `json:"field_name"`
	Mapping     string `json:"mapping"`
	Description string `json:"description"`
}

type mappingInfo struct {
	TableName   string          `json:"table_name"`
	Description string          `json:"description"`
	Mapping     []*fieldMapping `json:"mapping"`
}
type setMappingRequest struct {
	Session string         `json:"session"`
	Info    []*mappingInfo `json:"info"`
}

type setDatassetsApplyRequest struct {
	Session     string   `json:"session"`
	Title       string   `json:"title"`
	Description string   `json:"datassets_description"`
	From        []string `json:"datassets_from"`
	Other       string   `json:"datassets_from_other"`
}

type Progress struct {
	Duration     int32 `json:"duration"`
	SuccessCount int32 `json:"success_count"`
	FailCount    int32 `json:"fail_count"`
	RecordSize   int32 `json:"record_size"`
	Step         int32 `json:"step"`
}

type GenerateProcessResponse struct {
	Err  string    `json:"err_msg"`
	Body *Progress `json:"body"`
}

type getContractRequest struct {
	Session string `json:"session"`
	OrderId string `json:"order_id"`
}

type executeRequest struct {
	Session string `json:"session"`
	Id      string `json:"id"`
	Pass    string `json:"pass"`
}

type auditRequest struct {
	Session  string `json:"session"`
	Id       string `json:"id"`
	Approved bool   `json:"approved"`
	Reason   string `json:"reason"`
}
type distributeRequest struct {
	Session string `json:"session"`
	Id      string `json:"id"`
	Node    string `json:"trust_server"`
}

func testLoginClient(t *testing.T) {
	req := &clientLoginRequest{
		Username: "18500022713",
		Password: "123456",
		IsLocal:  false,
	}
	resp := ec.POST("/v1.ClientService/Login").WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/Login result: ", resp.Body())
	raw := resp.Body().Raw()
	var res clientLoginResponse
	if err := json.Unmarshal([]byte(raw), &res); err != nil {
		panic(err)
	}
	clientSession = res.Session
	fmt.Println("client user session: ", clientSession)
}

func testNewDatassets(t *testing.T) {
	req := opRequest{Session: clientSession}
	resp := ec.POST("/v1.ClientService/New").WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/New result: ", resp.Body())
}

func testListDatassets(t *testing.T) {
	req := opRequest{Session: clientSession}
	resp := ec.POST("/v1.ClientService/List").WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/List result: ", resp.Body())
}

func testClientStatus(t *testing.T) {
	resp := ec.POST("/v1.ClientService/Status").WithJSON(struct{}{}).Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/Status result: ", resp.Body())
}

func testSetDBConn(t *testing.T) {
	req := &setDBConnRequest{
		Session: clientSession,
		Info: &dbMessage{
			Src: &dbConnect{
				Type: 2,
				Host: "139.9.119.21",
				Port: 5432,
				User: "test",
				Pass: "datassets",
				Db:   "test",
			},
			Dst: &dbConnect{
				Type: 2,
				Host: "139.9.119.21",
				Port: 5432,
				User: "test",
				Pass: "datassets",
				Db:   "target3",
			},
		},
	}
	resp := ec.POST("/v1.ClientService/SetDBConn").WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/SetDBConn result: ", resp.Body())
}

func testGetTableList(t *testing.T) {
	resp := ec.POST("/v1.ClientService/GetTableList").
		WithJSON(&opRequest{clientSession}).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/GetTableList result: ", resp.Body())
}

func testSetMapping(t *testing.T) {
	b, err := ioutil.ReadFile("/Users/luxiaotong/code/datassets.cn/medias/test/area_mapping.json")
	if err != nil {
		panic(err)
	}
	var mm []*mappingInfo
	if err := json.Unmarshal(b, &mm); err != nil {
		panic(err)
	}
	req := &setMappingRequest{
		Session: clientSession,
		Info:    mm,
	}
	resp := ec.POST("/v1.ClientService/SetMappingInfo").
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/SetMappingInfo result: ", resp.Body())
}

func testSchematize(t *testing.T) {
	resp := ec.POST("/v1.ClientService/Schematize").
		WithJSON(&opRequest{clientSession}).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/Schematize result: ", resp.Body())
}

func testSchematizeProcess(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		fmt.Println("schematize process time: ", t.Format("2006-01-02 15:14:02"))
		resp := ec.POST("/v1.ClientService/SchematizeProcess").
			WithJSON(&opRequest{clientSession}).
			Expect().Status(http.StatusOK)
		fmt.Println("/v1.ClientService/SchematizeProcess result: ", resp.Body())
		raw := resp.Body().Raw()
		var res GenerateProcessResponse
		if err := json.Unmarshal([]byte(raw), &res); err != nil {
			panic(err)
		}
		if res.Body.SuccessCount+res.Body.FailCount >= res.Body.RecordSize {
			break
		}
	}
	fmt.Println("schematize step done")
}

func testSetDatassetsApply(t *testing.T) {
	testSchematizeProcess(t)
	req := &setDatassetsApplyRequest{
		Session:     clientSession,
		Title:       "河南省新乡市统计局数据2",
		Description: "数据资产描述2",
		From:        []string{"webCrawler", "other"},
		Other:       "testtesttest",
	}
	resp := ec.POST("/v1.ClientService/SetDatassetsApply").
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/SetDatassetsApply result: ", resp.Body())
}

func testGenerate(t *testing.T) {
	resp := ec.POST("/v1.ClientService/Generate").
		WithJSON(&opRequest{clientSession}).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/Generate result: ", resp.Body())
}

func testGenerateProcess(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		fmt.Println("genereate process time: ", t.Format("2006-01-02 15:14:02"))
		resp := ec.POST("/v1.ClientService/GenerateProcess").
			WithJSON(&opRequest{clientSession}).
			Expect().Status(http.StatusOK)
		fmt.Println("/v1.ClientService/GenerateProcess result: ", resp.Body())
		raw := resp.Body().Raw()
		var res GenerateProcessResponse
		if err := json.Unmarshal([]byte(raw), &res); err != nil {
			panic(err)
		}
		step := res.Body.Step
		fmt.Println("generate process step: ", step)
		if step == stepDone {
			break
		}
	}
	fmt.Println("generate step done")
}

func testGenerateDatassetsApplyPdf(t *testing.T) {
	testGenerateProcess(t)
	resp := ec.POST("/v1.ClientService/GenerateDatassetsApplyPdf").
		WithJSON(&opRequest{clientSession}).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/GenerateDatassetsApplyPdf result: ", resp.Body())
}

func testUploadDatassetsApplyPdf(t *testing.T) {
	resp := ec.POST("/v1.ClientService/UploadDatassetsApplyPdf").
		WithJSON(&opRequest{clientSession}).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/UploadDatassetsApplyPdf result: ", resp.Body())
}

func testGetOrders_Client(t *testing.T) {
	resp := ec.POST("/v1.ClientService/GetOrders").
		WithJSON(&opRequest{clientSession}).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/GetOrders result: ", resp.Body())
}

func testGetContract(t *testing.T) {
	req := &getContractRequest{
		Session: clientSession,
		OrderId: strconv.FormatInt(orderID, 10),
	}
	resp := ec.POST("/v1.ClientService/GetContract").
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/GetContract result: ", resp.Body())
}

func testKeyPair(t *testing.T) {
	resp := ec.POST("/v1.ClientService/KeyPair").
		WithJSON(&opRequest{clientSession}).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/KeyPair result: ", resp.Body())
}

func testExecute(t *testing.T) {
	// orderID = 1531486245002153984
	req := &executeRequest{
		Session: clientSession,
		Id:      strconv.FormatInt(orderID, 10),
		Pass:    "datassets",
	}
	resp := ec.POST("/v1.ClientService/Execute").
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/Execute result: ", resp.Body())
}

func testAudit(t *testing.T) {
	req := &auditRequest{
		Session:  clientSession,
		Id:       strconv.FormatInt(orderID, 10),
		Approved: true,
	}
	resp := ec.POST("/v1.ClientService/Audit").
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/Audit result: ", resp.Body())
}

func testDistribute(t *testing.T) {
	req := &distributeRequest{
		Session: clientSession,
		Id:      strconv.FormatInt(orderID, 10),
	}
	resp := ec.POST("/v1.ClientService/Distribute").
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/v1.ClientService/Distribute result: ", resp.Body())
}
