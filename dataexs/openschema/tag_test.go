package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

var (
	tagID int64
)

const tagName = "tag1"

type TagRequest struct {
	ID     int64  `json:"id,string"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Status int32  `json:"status"`
	Reason string `json:"reason"`
}

type GetTagsRequest struct {
	Query  string `json:"q"`
	Status int32  `json:"status"`
}

func testAddTag(t *testing.T) {
	req := &TagRequest{
		Name: tagName,
		Desc: "desc1",
	}
	resp := e.POST("/tag").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/tag create response: %v\n", resp.Body())
}

func testOpTag(t *testing.T) {
	req := &TagRequest{
		ID:     tagID,
		Status: 30,
		// Reason: "refuse",
	}
	resp := e.POST("/tag/status").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/tag/status op response: %v\n", resp.Body())
}

func testGetTags(t *testing.T) {
	req := &GetTagsRequest{}
	resp := e.POST("/tags").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/tags get response: %v\n", resp.Body())

	list := resp.JSON().Object().Value("data").Object().Value("list")
	// fmt.Println("list: ", list)
	var id string
	for _, val := range list.Array().Iter() {
		if val.Object().Value("name").String().Raw() == tagName {
			id = val.Object().Value("id").String().Raw()
			break
		}
	}
	tagID, _ = strconv.ParseInt(id, 10, 64)
	fmt.Println("tag id: ", tagID)
}

func clearTags() {
	dsn := "host=139.9.119.21 port=5432 user=auth password=authpass dbname=openschema sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("open error: ", err)
		return
	}
	defer db.Close()

	if _, err := db.Exec("delete from tags where name=$1", tagName); err != nil {
		fmt.Println("delete error: ", err)
		return
	}
}
