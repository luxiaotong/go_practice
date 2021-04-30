package main

import (
	"fmt"
	"net/http"
	"testing"
)

var dictID string

type DictRequest struct {
	ID      int64  `json:"id,string"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Type    int32  `json:"type"`
}

func testAddDict(t *testing.T) {
	req := &DictRequest{
		Name:    "dict vote",
		Version: "v1",
		Title:   "vote title",
		Desc:    "dict vote desc",
		Type:    20,
	}
	resp := e.POST("/dict").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict create response: %v\n", resp.Body())
	dictID = resp.JSON().Object().Value("data").Object().Value("id").String().Raw()
	fmt.Println("dict id: ", dictID)
}
