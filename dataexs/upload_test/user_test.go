package main

import (
	"fmt"
	"net/http"
	"testing"
)

const cookieSecret = "crc_key"

var (
	userCookie string
	userToken  string
)

type SignInRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func testSignIn(t *testing.T) {
	req := SignInRequest{
		Mobile:   "18500022713",
		Password: "123456",
	}
	resp := e.POST("/user/signin").WithJSON(req).Expect().Status(http.StatusOK)
	userCookie = resp.Cookie(cookieSecret).Value().Raw()
	userToken = resp.JSON().Object().Value("data").Object().Value("token").String().Raw()
	fmt.Printf("user cookie: %s, token: %s\n", userCookie, userToken)
}
