package main

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

type Field struct {
	ID         int64    `json:"id,string"`
	SrcName    string   `json:"src_name"`
	SrcType    string   `json:"src_type"`
	SrcComment string   `json:"src_comment"`
	LabelEN    string   `json:"label_en"`
	CommentCN  string   `json:"comment_cn"`
	Tags       []string `json:"tags"`
	Status     int32    `json:"status"`
	Votes      int32    `json:"votes"`
}

type VoteFieldRequest struct {
	ID         int64  `json:"id,string"`
	Type       int32  `json:"type"`
	Suggestion string `json:"suggestion"`
}

func testFillField(t *testing.T) {
	id, _ := strconv.ParseInt(fieldID, 10, 64)
	req := &Field{
		ID:        id,
		LabelEN:   "label_fill",
		CommentCN: "comment_fill",
		Tags:      []string{tagName},
	}
	resp := e.PUT("/field/fill").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/field/fill response: %v\n", resp.Body())
}

func testVoteField(t *testing.T) {
	id, _ := strconv.ParseInt(fieldID, 10, 64)
	req := &VoteFieldRequest{
		ID:   id,
		Type: 20,
	}
	resp := e.PUT("/field/vote").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/field/vote response: %v\n", resp.Body())
}
