package main

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

var (
	adminCookie string
	adminToken  string
	adminUID    int64
)

type GetUsersRequest struct {
	Query string `json:"q"`
	Type  int32  `json:"type"`
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
		Role:     50,
	}
	resp := e.PUT("/user/info").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("user/info update %d response: %v\n", uid, resp.Body())
}

func testGetApplications(t *testing.T) {
	req := GetUsersRequest{
		// Query: "shannon",
		Type: 30,
	}
	resp := e.POST("/applications").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/applications response: %v\n", resp.Body())
}

func testAuditUser_Voter(t *testing.T) {
	req := GetUsersRequest{
		Type: 10,
	}
	resp := e.POST("/applications").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	list := resp.JSON().Object().Value("data").Object().Value("list").Array()
	var aid string
	for _, val := range list.Iter() {
		userID, _ := strconv.ParseInt(val.Object().Value("id").String().Raw(), 10, 64)
		t := int32(val.Object().Value("application").Object().Value("type").Number().Raw())
		if userID != uid || t != 10 {
			continue
		}
		fmt.Println("user id: ", userID)
		fmt.Println("type: ", t)
		aid = val.Object().Value("application").Object().Value("id").String().Raw()
		fmt.Println("application id: ", aid)
	}
	id, _ := strconv.ParseInt(aid, 10, 64)
	req2 := &UserRequest{
		Apply: &Apply{
			ID:     id,
			Status: 30,
		},
	}
	resp = e.POST("/user/audit").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req2).Expect().Status(http.StatusOK)
	fmt.Printf("/user/audit voter %d response: %v\n", uid, resp.Body())
}

func testAuditUser_Provider(t *testing.T) {
	req := GetUsersRequest{
		Type: 10,
	}
	resp := e.POST("/applications").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	list := resp.JSON().Object().Value("data").Object().Value("list").Array()
	var aid string
	for _, val := range list.Iter() {
		userID, _ := strconv.ParseInt(val.Object().Value("id").String().Raw(), 10, 64)
		t := int32(val.Object().Value("application").Object().Value("type").Number().Raw())
		if userID != uid || t != 20 {
			continue
		}
		fmt.Println("user id: ", userID)
		fmt.Println("type: ", t)
		aid = val.Object().Value("application").Object().Value("id").String().Raw()
		fmt.Println("application id: ", aid)
	}
	id, _ := strconv.ParseInt(aid, 10, 64)
	req2 := &UserRequest{
		Apply: &Apply{
			ID:     id,
			Status: 30,
		},
	}
	resp = e.POST("/user/audit").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req2).Expect().Status(http.StatusOK)
	fmt.Printf("/user/audit provider %d response: %v\n", uid, resp.Body())
}

func testAuditUser_SDK(t *testing.T) {
	req := GetUsersRequest{
		Type: 20,
	}
	resp := e.POST("/applications").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	list := resp.JSON().Object().Value("data").Object().Value("list").Array()
	var aid string
	for _, val := range list.Iter() {
		userID, _ := strconv.ParseInt(val.Object().Value("id").String().Raw(), 10, 64)
		t := int32(val.Object().Value("application").Object().Value("type").Number().Raw())
		if userID != uid || t != 30 {
			continue
		}
		fmt.Println("user id: ", userID)
		fmt.Println("type: ", t)
		aid = val.Object().Value("application").Object().Value("id").String().Raw()
		fmt.Println("application id: ", aid)
	}
	id, _ := strconv.ParseInt(aid, 10, 64)
	req2 := &UserRequest{
		Apply: &Apply{
			ID:     id,
			Status: 30,
		},
	}
	resp = e.POST("/user/audit").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req2).Expect().Status(http.StatusOK)
	fmt.Printf("/user/audit sdk %d response: %v\n", uid, resp.Body())
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
