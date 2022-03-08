package testdatassets

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func testUploadSample(t *testing.T) {
	f := "/Users/luxiaotong/code/datassets.cn/medias/test/sample.json"
	resp := eu.POST("/upload/file").
		WithHeader("Authorization", "Bearer "+tokenValSeller).
		WithCookie(jwtCookieSecret, tokenKeySeller).
		WithMultipart().
		WithFile("file", f).WithFormField("api_type", "sample").
		Expect().Status(http.StatusOK)
	fmt.Println("/upload/file sample result: ", resp.Body())
	name := resp.JSON().Object().Value("data").Object().Value("name").String().Raw()
	ss := strings.Split(name, "/")
	sampleTmp = ss[len(ss)-1]
}

func testUploadApplication(t *testing.T) {
	f := "/Users/luxiaotong/code/datassets.cn/medias/test/SetDatassetsApply.pdf"
	resp := eu.POST("/upload/application_pdf").
		WithHeader("Authorization", "Bearer "+tokenValSeller).
		WithCookie(jwtCookieSecret, tokenKeySeller).
		WithMultipart().WithFile("pdf", f).
		Expect().Status(http.StatusOK)
	fmt.Println("/upload/file application result: ", resp.Body())
	name := resp.JSON().Object().Value("data").Object().Value("name").String().Raw()
	ss := strings.Split(name, "/")
	applicationTmp = ss[len(ss)-1]
}
