package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testing"
)

var (
	dictID  string
	fieldID string
)

type DictRequest struct {
	ID          int64    `json:"id,string"`
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Title       string   `json:"title"`
	Desc        string   `json:"desc"`
	Industry    string   `json:"industry"`
	SubIndustry string   `json:"sub_industry"`
	Tags        []string `json:"tags"`
	Type        int32    `json:"type"`
	Attach      string   `json:"attach"`
	Fields      []*Field `json:"fields"`
	Status      int32    `json:"status"`
	Reason      string   `json:"reason"`
}

type GetDictsRequest struct {
	Query       string `json:"q"`
	Industry    string `json:"industry"`
	SubIndustry string `json:"sub_industry"`
	Type        int32  `json:"type"`
}

func testAddDict_Definition(t *testing.T) {
	req := &DictRequest{
		Name:    "dict vote",
		Version: "v1",
		Title:   "vote title",
		Desc:    "dict vote desc",
		Type:    10,
		Fields: []*Field{
			&Field{
				SrcName:    "field1",
				SrcType:    "varchar",
				SrcComment: "",
				LabelEN:    "",
				CommentCN:  "",
				Tags:       []string{},
			},
			&Field{
				SrcName:    "field2",
				SrcType:    "varchar",
				SrcComment: "",
				LabelEN:    "",
				CommentCN:  "",
				Tags:       []string{},
			},
		},
		Industry:    "农、林、牧、鱼",
		SubIndustry: "农业",
	}
	resp := e.POST("/dict").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict create response: %v\n", resp.Body())
	dictID = resp.JSON().Object().Value("data").Object().Value("id").String().Raw()
	fmt.Println("dict id: ", dictID)
}

func testAddDict_Vote(t *testing.T) {
	req := &DictRequest{
		Name:    "dict vote",
		Version: "v1",
		Title:   "vote title",
		Desc:    "dict vote desc",
		Type:    20,
		Fields: []*Field{
			&Field{
				SrcName:    "field1",
				SrcType:    "varchar",
				SrcComment: "",
				LabelEN:    "label1",
				LabelCN:    "label中文",
				CommentCN:  "comment1",
				Tags:       []string{tagName},
			},
			&Field{
				SrcName:    "field2",
				SrcType:    "varchar",
				SrcComment: "",
				LabelEN:    "label2",
				LabelCN:    "label中文",
				CommentCN:  "comment2",
				Tags:       []string{tagName},
			},
		},
		Industry:    "农、林、牧、鱼",
		SubIndustry: "农业",
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
	fmt.Printf("/dicts/gets response: %v\n", resp.Body())
}

func testSearchDicts(t *testing.T) {
	req := &GetDictsRequest{}
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

func testOpDict_Close(t *testing.T) {
	id, _ := strconv.ParseInt(dictID, 10, 64)
	req := &DictRequest{
		ID:     id,
		Status: 60,
	}
	resp := e.POST("/dict/status").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict/status close %v response: %v\n", dictID, resp.Body())
}

func testOpDict_Reopen(t *testing.T) {
	id, _ := strconv.ParseInt(dictID, 10, 64)
	req := &DictRequest{
		ID:     id,
		Status: 40,
	}
	resp := e.POST("/dict/status").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/dict/status reopen %v response: %v\n", dictID, resp.Body())
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
