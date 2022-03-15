package ma

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
)

const maURL = "http://124.126.76.161:31026/"

const (
	appKey    = "6A9FD3C243F84D6281534A0130FB1E6D"
	appSecret = "2C6F6137C4734FB686FC752D1425A228"
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
	Name string `json:"name"`
	Age  string `json:"age"`
}

type saveRequest struct {
	AnalysisPath string    `json:"analysisPath"`
	Data         *saveData `json:"data"`
	DataAttr     int32     `json:"dataAttr"`
	DataType     int32     `json:"dataType"`
	TemplateID   int64     `json:"templateId"`
}

func testSaveAnalysis(t *testing.T) {
	req := &saveRequest{
		AnalysisPath: "MA.110.9902.XXCY001_122-11/DA11111",
		Data: &saveData{
			Name: "testdatassetsname",
			Age:  "19",
		},
		DataAttr:   0,
		DataType:   1,
		TemplateID: 11,
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
		WithQuery("analysisUrl", "MA.110.9902.XXCY001_122-11/DA11111").
		Expect().Status(http.StatusOK)
	fmt.Println("/ma-api/analysis/ma result: ", resp.Body())
}

func testNodeInfo(t *testing.T) {
	resp := ema.GET("/ma-api/analysis/nodeInfo").
		WithHeader("token", token).
		WithQuery("analysisUrl", "MA.110.9902.XXCY001_122-11/DA11111").
		Expect().Status(http.StatusOK)
	fmt.Println("/ma-api/analysis/nodeInfo result: ", resp.Body())
}
