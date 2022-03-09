package testdatassets

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func testUploadSample(t *testing.T) {
	f := "/Users/luxiaotong/code/datassets.cn/medias/test/area_sample.json"
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

func testUploadVoucher(t *testing.T) {
	f := "/Users/luxiaotong/code/datassets.cn/medias/test/test.jpg"
	resp := eu.POST("/upload/file").
		WithHeader("Authorization", "Bearer "+tokenValBuyer).
		WithCookie(jwtCookieSecret, tokenKeyBuyer).
		WithMultipart().
		WithFile("file", f).WithFormField("api_type", "voucher").
		Expect().Status(http.StatusOK)
	fmt.Println("/upload/file voucher result: ", resp.Body())
	name := resp.JSON().Object().Value("data").Object().Value("name").String().Raw()
	ss := strings.Split(name, "/")
	voucherTmp = ss[len(ss)-1]
}
