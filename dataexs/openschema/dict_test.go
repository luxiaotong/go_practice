package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	Status  int32    `json:"status"`
	Reason  string   `json:"reason"`
}

type Field struct {
	SrcName    string   `json:"src_name"`
	SrcType    string   `json:"src_type"`
	SrcComment string   `json:"src_comment"`
	LabelEN    string   `json:"label_en"`
	CommentCN  string   `json:"comment_cn"`
	Tags       []string `json:"tags"`
}

type GetDictsRequest struct {
	Query  string `json:"q"`
	Tag    string `json:"tag"`
	Status int32  `json:"status"`
	Type   int32  `json:"type"`
}

type GetFieldsRequest struct {
	DictID int64 `json:"dict_id,string"`
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

func testGetDicts(t *testing.T) {
	req := &GetDictsRequest{}
	resp := e.POST("/dicts/gets").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dicts/admin response: %v\n", resp.Body())
}

func testSearchDicts(t *testing.T) {
	req := &GetDictsRequest{
		Tag:  tagName,
		Type: 20,
	}
	resp := e.POST("/dicts/search").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dicts/search response: %v\n", resp.Body())
}

func testGetDict(t *testing.T) {
	resp := e.GET("/dict/info/"+dictID).
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		Expect().Status(http.StatusOK)
	fmt.Printf("/dict/info dictID response: %v\n", resp.Body())
}

func testGetFields(t *testing.T) {
	id, _ := strconv.ParseInt(dictID, 10, 64)
	req := &GetFieldsRequest{id}
	resp := e.POST("/dict/fields").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict/fields %v response: %v\n", id, resp.Body())
}

func testOpDict(t *testing.T) {
	id, _ := strconv.ParseInt(dictID, 10, 64)
	req := &DictRequest{
		ID:     id,
		Status: 20,
	}
	resp := e.POST("/dict/status").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict/status %v response: %v\n", dictID, resp.Body())
}

func testEditDict(t *testing.T) {
	id, _ := strconv.ParseInt(dictID, 10, 64)
	req := &DictRequest{
		ID:      id,
		Name:    "dict vote222",
		Version: "v1222",
		Title:   "vote title222",
		Desc:    "dict vote desc22",
	}
	resp := e.PUT("/dict").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict edit %v response: %v\n", dictID, resp.Body())

	req = &DictRequest{
		ID:     id,
		Status: 20,
	}
	resp = e.POST("/dict/status").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict/status %v response: %v\n", dictID, resp.Body())
}

func testAuditDict(t *testing.T) {
	id, _ := strconv.ParseInt(dictID, 10, 64)
	req := &DictRequest{
		ID:     id,
		Status: 40,
	}
	resp := e.POST("/dict/audit").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict/audit %v response: %v\n", dictID, resp.Body())
}

func testDeleteDict(t *testing.T) {
	req := &DictRequest{
		ID: 1390963635957796864,
	}
	resp := e.DELETE("/dict").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict delete response: %v\n", resp.Body())
}
