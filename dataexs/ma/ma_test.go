package ma

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
)

const maURL = "http://124.126.76.161:31026/"

const (
	appKey    = "C15CEA486F7B4D9BBFB37BFADCBA863D"
	appSecret = "3BB0B491417E4204965A65BE80CA6A12"
)

var ctx context.Context

var ema *httpexpect.Expect

var token string

func TestMain(m *testing.M) {
	ctx = context.Background()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestAll(t *testing.T) {
	ema = httpexpect.New(t, maURL)
	t.Run("testGetToken", testGetToken)
	t.Run("testSaveAnalysis", testSaveAnalysis)
	t.Run("testAnalysis", testAnalysis)
	t.Run("testNodeInfo", testNodeInfo)
}

func testGetToken(t *testing.T) {
	resp := ema.GET("/token/get").
		WithHeader("appKey", appKey).
		WithHeader("appSecret", appSecret).
		Expect().Status(http.StatusOK)
	fmt.Println("/token/get result: ", resp.Body())
	token = resp.JSON().Object().Value("data").Object().Value("token").String().Raw()
	fmt.Println("token: ", token)
}

type saveData struct {
	Name    string   `json:"name"`
	Firm    string   `json:"firm"`
	DRs     int64    `json:"drs"`
	Sources []string `json:"sources"`
	BlockID string   `json:"block_id"`
}

type saveRequest struct {
	AnalysisPath string `json:"analysisPath"`
	Data         string `json:"data"`
	DataAttr     int32  `json:"dataAttr"`
	DataType     int32  `json:"dataType"`
	TemplateID   int64  `json:"templateId,omitempty"`
}

func testSaveAnalysis(t *testing.T) {
	d := &saveData{
		Name:    "海尔智家2019年度经营状况数据",
		Firm:    "海尔电器集团",
		DRs:     int64(12643987),
		Sources: []string{"提供应用服务取得"},
		BlockID: "14763ef63742a03643572742ec5ee470d30e599d",
	}
	b, _ := json.Marshal(d)
	req := &saveRequest{
		AnalysisPath: "MA.110.9902.2222/DA11111",
		Data:         string(b),
		DataAttr:     0,
		DataType:     1,
	}
	resp := ema.POST("/ma-api/analysis/save").
		WithHeader("token", token).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/ma-api/analysis/save result: ", resp.Body())
}

func testAnalysis(t *testing.T) {
	resp := ema.GET("/ma-api/analysis/ma").
		WithHeader("token", token).
		WithQuery("analysisUrl", "MA.110.9902.2222/DA11111").
		Expect().Status(http.StatusOK)
	fmt.Println("/ma-api/analysis/ma result: ", resp.Body())
}

func testNodeInfo(t *testing.T) {
	resp := ema.GET("/ma-api/analysis/nodeInfo").
		WithHeader("token", token).
		WithQuery("analysisUrl", "MA.110.9902.2222/DA11111").
		Expect().Status(http.StatusOK)
	fmt.Println("/ma-api/analysis/nodeInfo result: ", resp.Body())
}
