package main

import (
	"fmt"
	"net/http"
	"testing"
)

var (
	adminCookie string
	adminToken  string
	adminUID    int64
)

type GetUsersRequest struct {
	Query  string `json:"q"`
	Status int32  `json:"status"`
}

func testAdminSignIn(t *testing.T) {
	req := SignInRequest{
		Username: "root",
		Password: "123456",
	}
	resp := e.POST("/user/signin").WithJSON(req).Expect().Status(http.StatusOK)
	adminCookie = resp.Cookie(CookieSecret).Value().Raw()
	adminToken = resp.JSON().Object().Value("data").Object().Value("token").String().Raw()
	fmt.Println("admin cookie: ", adminCookie)
	fmt.Println("admin token: ", adminToken)
	adminUID = parse(adminToken, adminCookie)
}

func testGetUsers(t *testing.T) {
	req := GetUsersRequest{
		// Query: "abbraaaa",
		// Status: 10,
	}
	resp := e.POST("/users").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("users response: %v\n", resp.Body())
}

func testAdminGetUser(t *testing.T) {
	req := &UserRequest{
		ID: uid,
	}
	resp := e.POST("/user/info").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("user/info %d response: %v\n", uid, resp.Body())
}

func testAdminUpdateUser(t *testing.T) {
	req := &UserRequest{
		ID:       uid,
		Mobile:   "18500022713",
		Name:     "shannon",
		Email:    "shannon@datassets.cn",
		FirmName: "firm_name_2",
		FirmAbbr: "firm_abbr_2",
	}
	resp := e.PUT("/user/info").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("user/info update %d response: %v\n", uid, resp.Body())
}

func testAuditUser(t *testing.T) {
	req := &UserRequest{
		ID:     uid,
		Status: 30,
	}
	resp := e.POST("/user/audit").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("audit %d response: %v\n", uid, resp.Body())
}

func testFreezeUser(t *testing.T) {
	req := &UserRequest{
		ID:      uid,
		Enabled: false,
	}
	resp := e.POST("/user/freeze").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("freeze %d response: %v\n", uid, resp.Body())
}
