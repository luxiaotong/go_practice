package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const CookieSecret = "crc_key"

var (
	cookieVal string
	token     string
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PassRequest struct {
	New string `json:"new"`
	Old string `json:"old"`
}

func testSignIn(t *testing.T) {
	req := SignInRequest{
		Username: "shannon",
		Password: "123456",
	}
	resp := e.POST("/user/signin").WithJSON(req).Expect().Status(http.StatusOK)
	cookieVal = resp.Cookie(CookieSecret).Value().Raw()
	token = resp.JSON().Object().Value("data").Object().Value("token").String().Raw()
	fmt.Println("cookie val: ", cookieVal)
	fmt.Println("token: ", token)
	assert.NotEmpty(t, cookieVal)
	assert.NotEmpty(t, token)
}

func testUserPass(t *testing.T) {
	req := PassRequest{
		Old: "123456",
		New: "123456",
	}
	resp := e.POST("/user/pass").WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("user/pass response: %v\n", resp.Body())
}
