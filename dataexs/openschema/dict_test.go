package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

var dictID string

type DictRequest struct {
	ID      int64    `json:"id,string"`
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Title   string   `json:"title"`
	Desc    string   `json:"desc"`
	Type    int32    `json:"type"`
	Attach  string   `json:"attach"`
	Fields  []*Field `json:"fields"`
}

type Field struct {
	SrcName    string   `json:"src_name"`
	SrcType    string   `json:"src_type"`
	SrcComment string   `json:"src_comment"`
	LabelEN    string   `json:"label_en"`
	CommentCN  string   `json:"comment_cn"`
	Tags       []string `json:"tags"`
}

func testAddDict(t *testing.T) {
	req := &DictRequest{
		Name:    "dict vote",
		Version: "v1",
		Title:   "vote title",
		Desc:    "dict vote desc",
		Type:    20,
		Fields: []*Field{
			&Field{"field1", "varchar", "", "label1", "comment1", []string{tagName}},
			&Field{"field2", "varchar", "", "label2", "comment2", []string{tagName}},
		},
	}
	resp := e.POST("/dict").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict create response: %v\n", resp.Body())
	dictID = resp.JSON().Object().Value("data").Object().Value("id").String().Raw()
	fmt.Println("dict id: ", dictID)
}

func testGetDictAttach_Definition(t *testing.T) {
	req := &DictRequest{
		Type:   10,
		Attach: dictDefinition,
	}
	resp := e.POST("/dict/attach").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	log.Printf("/dict/attach definition response: %v\n", resp.Body())
}

func testGetDictAttach_Vote(t *testing.T) {
	req := &DictRequest{
		Type:   20,
		Attach: dictVote,
	}
	resp := e.POST("/dict/attach").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	log.Printf("/dict/attach vote response: %v\n", resp.Body())
}
