package main

import (
	"fmt"
	"net/http"
	"testing"
)

type Group struct {
	ID   int64  `json:"id,string"`
	Name string `json:"name"`
}

func testAddGroup(t *testing.T) {
	req := &Group{
		Name: "testgroup",
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
