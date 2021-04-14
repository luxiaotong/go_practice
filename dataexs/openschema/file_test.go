package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
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
	logo := mediasDIR + fmt.Sprintf("firm/%d.logo.jpg", uid)
	_ = os.Remove(logo)
}
