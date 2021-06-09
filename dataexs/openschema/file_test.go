package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/rs/zerolog/log"
)

const mediasDIR = "/Users/luxiaotong/code/datassets.cn/medias/"

func testUploadFile(t *testing.T) {
	f, _ := ioutil.TempFile("", "*.jpg")
	fmt.Println("tmp refund voucher: ", f.Name())
	defer os.Remove(f.Name())
	defer f.Close()
	alpha := image.NewAlpha(image.Rect(0, 0, 100, 100))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			alpha.Set(x, y, color.Alpha{uint8(x % 256)})
		}
	}
	_ = jpeg.Encode(f, alpha, nil)

	resp := e.POST("/upload/file").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithMultipart().
		WithFile("file", f.Name()).WithFormField("api_type", "logo").
		Expect().Status(http.StatusOK)
	fmt.Printf("/upload/file result: %v\n", resp.Body())
}

func clearLogo(uid int64) {
	logo := mediasDIR + fmt.Sprintf("user/%d.logo.jpg", uid)
	_ = os.Remove(logo)
}

var (
	dictDefinition string
	dictVote       string
)

func testUploadFile_Definition(t *testing.T) {
	f, _ := ioutil.TempFile("", "definition.*.csv")
	log.Print("tmp definition csv: ", f.Name())
	defer f.Close()
	w := csv.NewWriter(f)
	data := [][]string{
		{"字段名称", "字段类型", "备注"},
		{"field1", "int", ""},
		{"field2", "varchar", ""},
	}
	_ = w.WriteAll(data)

	resp := e.POST("/upload/file").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithMultipart().
		WithFile("file", f.Name()).WithFormField("api_type", "definition_fields").
		Expect().Status(http.StatusOK)
	log.Print("/upload/file definition result: ", resp.Body())
	dictDefinition = resp.JSON().Object().Value("data").Object().Value("name").String().Raw()
	dictDefinition = dictDefinition[6:]
}

func testUploadFile_Vote(t *testing.T) {
	f, _ := ioutil.TempFile("", "vote.*.csv")
	log.Print("tmp vote csv: ", f.Name())
	defer f.Close()
	w := csv.NewWriter(f)
	data := [][]string{
		{"字段名称", "字段类型", "备注", "schema字段英文名称", "schema字段中文名称", "schema字段中文注释", "分类"},
		{"field1", "int", "", "label1", "label cn", "comment cn", "tag1"},
		{"field2", "varchar", "", "label2", "label cn", "comment cn", "tag1"},
	}
	_ = w.WriteAll(data)

	resp := e.POST("/upload/file").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithMultipart().
		WithFile("file", f.Name()).WithFormField("api_type", "vote_fields").
		Expect().Status(http.StatusOK)
	log.Print("/upload/file vote result: ", resp.Body())
	dictVote = resp.JSON().Object().Value("data").Object().Value("name").String().Raw()
	dictVote = dictVote[6:]
}
