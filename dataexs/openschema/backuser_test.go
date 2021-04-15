package main

import (
	"fmt"
	"net/http"
	"testing"
)

var (
	rootCookie string
	rootToken  string
	rootUID    int64
)

type GetUsersRequest struct {
	Query  string `json:"q"`
	Status int32  `json:"status"`
}

func testRootSignIn(t *testing.T) {
	req := SignInRequest{
		Username: "root",
		Password: "123456",
	}
	resp := e.POST("/user/signin").WithJSON(req).Expect().Status(http.StatusOK)
	rootCookie = resp.Cookie(CookieSecret).Value().Raw()
	rootToken = resp.JSON().Object().Value("data").Object().Value("token").String().Raw()
	fmt.Println("root cookie: ", rootCookie)
	fmt.Println("root token: ", rootToken)
	rootUID = parse(rootToken, rootCookie)
}

func testGetUsers(t *testing.T) {
	req := GetUsersRequest{
		// Query: "abbraaaa",
		// Status: 10,
	}
	resp := e.POST("/users").
		WithHeader("Authorization", "Bearer "+rootToken).
		WithCookie(CookieSecret, rootCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("users response: %v\n", resp.Body())
}
