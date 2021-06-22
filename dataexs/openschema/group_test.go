package main

import (
	"fmt"
	"net/http"
	"testing"
)

type OpMemberRequest struct {
	ID     int64    `json:"id,string"`
	UIDs   []string `json:"uids"`
	Action int32    `json:"action"`
}

type Group struct {
	ID          int64  `json:"id,string"`
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	SubIndustry string `json:"sub_industry"`
}

type GetGroupsRequest struct {
	PageIndex uint32 `json:"page_index"`
	PageSize  uint32 `json:"page_size"`
}

type GetMembersRequest struct {
	ID        int64  `json:"id,string"`
	Query     string `json:"q"`
	PageIndex uint32 `json:"page_index"`
	PageSize  uint32 `json:"page_size"`
}

func testAddGroup(t *testing.T) {
	req := &Group{
		Name:        "testgroup",
		Industry:    "农、林、牧、鱼",
		SubIndustry: "农业",
	}
	resp := e.POST("/group").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/group add response: %v\n", resp.Body())
}

func testEditGroup(t *testing.T) {
	req := &Group{
		ID:   1401747325717581824,
		Name: "testgroup2",
	}
	resp := e.PUT("/group").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/group edit response: %v\n", resp.Body())
}

func testGetGroups(t *testing.T) {
	req := GetGroupsRequest{}
	resp := e.POST("/groups").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/groups response: %v\n", resp.Body())
}

func testGetGroup(t *testing.T) {
	req := &Group{
		ID: 1401747325717581824,
	}
	resp := e.POST("/group/info").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/group/info response: %v\n", resp.Body())
}

func testJoinGroup(t *testing.T) {
	req := &OpMemberRequest{
		ID:     1401747325717581824,
		UIDs:   []string{"1400649470592421888"},
		Action: 10,
	}
	resp := e.PUT("/group/member").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/group/member join response: %v\n", resp.Body())
}

func testLeaveGroup(t *testing.T) {
	req := &OpMemberRequest{
		ID:     1401747325717581824,
		UIDs:   []string{"1400649470592421888"},
		Action: 20,
	}
	resp := e.PUT("/group/member").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/group/member leave response: %v\n", resp.Body())
}

func testGetMembers(t *testing.T) {
	req := &GetMembersRequest{
		ID: 0,
	}
	resp := e.POST("/group/members").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/group/members response: %v\n", resp.Body())
}
