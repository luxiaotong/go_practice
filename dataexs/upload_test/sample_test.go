package main

import (
	"fmt"
	"net/http"
	"testing"
)

func testUploadFile_Sample(t *testing.T) {
	resp := u.POST("/upload/file").
		WithHeader("Authorization", "Bearer "+userToken).
		WithCookie(cookieSecret, userCookie).
		WithMultipart().
		WithFile("file", "./sample.json").WithFormField("api_type", "sample").
		Expect().Status(http.StatusOK)
	fmt.Println("/upload/file sample result: ", resp.Body())
}
