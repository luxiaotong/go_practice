package main

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

type GetFieldsRequest struct {
	DictID      int64   `json:"dict_id,string"`
	Query       string  `json:"q"`
	Industry    string  `json:"industry"`
	SubIndustry string  `json:"sub_industry"`
	Status      []int32 `json:"status"`
	PageIndex   uint32  `json:"page_index"`
	PageSize    uint32  `json:"page_size"`
}

type Field struct {
	ID         int64    `json:"id,string"`
	SrcName    string   `json:"src_name"`
	SrcType    string   `json:"src_type"`
	SrcComment string   `json:"src_comment"`
	LabelEN    string   `json:"label_en"`
	LabelCN    string   `json:"label_cn"`
	CommentCN  string   `json:"comment_cn"`
	Tags       []string `json:"tags"`
	Status     int32    `json:"status"`
	VoteCount  int32    `json:"vote_count"`
}

type RecommendRequest struct {
	FieldID   int64    `json:"field_id,string"`
	LabelEN   string   `json:"label_en"`
	LabelCN   string   `json:"label_cn"`
	CommentCN string   `json:"comment_cn"`
	Tags      []string `json:"tags"`
}

type VoteFieldRequest struct {
	ID   int64 `json:"id,string"`
	Type int32 `json:"type"`
}

type GetVotesRequest struct {
	FieldID   int64  `json:"field_id,string"`
	Type      int32  `json:"type"`
	Query     string `json:"q"`
	PageIndex uint32 `json:"page_index"`
	PageSize  uint32 `json:"page_size"`
}

func testGetFields(t *testing.T) {
	id, _ := strconv.ParseInt(dictID, 10, 64)
	req := &GetFieldsRequest{DictID: id}
	resp := e.POST("/fields").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/fields %v response: %v\n", id, resp.Body())

	fieldID = resp.JSON().Object().Value("data").
		Object().Value("list").Array().Element(0).
		Object().Value("id").String().Raw()
	fmt.Println("field id: ", fieldID)
}

func testFillField(t *testing.T) {
	id, _ := strconv.ParseInt(dictID, 10, 64)
	req := &GetFieldsRequest{DictID: id}
	resp := e.POST("/fields").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	list := resp.JSON().Object().Value("data").Object().Value("list").Array()
	for _, val := range list.Iter() {
		fieldID := val.Object().Value("id").String().Raw()
		fmt.Println("field id: ", fieldID)
		id, _ := strconv.ParseInt(fieldID, 10, 64)
		req := &RecommendRequest{
			FieldID:   id,
			LabelEN:   "label_fill",
			LabelCN:   "本属性的中文",
			CommentCN: "comment_fill",
			Tags:      []string{tagName},
		}
		resp := e.POST("/field/recommend").
			WithHeader("Authorization", "Bearer "+token).
			WithCookie(CookieSecret, cookieVal).
			WithJSON(req).Expect().Status(http.StatusOK)
		fmt.Printf("/field/recommend %v response: %v\n", id, resp.Body())
	}
	// req := &Field{
	// 	ID:        1405069047694888960,
	// 	LabelEN:   "label_fill",
	// 	LabelCN:   "本属性的中文",
	// 	CommentCN: "comment_fill",
	// 	Tags:      []string{tagName},
	// }
	// resp := e.POST("/field/recommend").
	// 	WithHeader("Authorization", "Bearer "+token).
	// 	WithCookie(CookieSecret, cookieVal).
	// 	WithJSON(req).Expect().Status(http.StatusOK)
	// fmt.Printf("/field/recommend response: %v\n", resp.Body())
}

func testVoteField(t *testing.T) {
	// id, _ := strconv.ParseInt(dictID, 10, 64)
	// req := &GetFieldsRequest{DictID: id}
	// resp := e.POST("/fields").
	// 	WithHeader("Authorization", "Bearer "+token).
	// 	WithCookie(CookieSecret, cookieVal).
	// 	WithJSON(req).Expect().Status(http.StatusOK)
	// list := resp.JSON().Object().Value("data").Object().Value("list").Array()
	// for _, val := range list.Iter() {
	// 	fieldID := val.Object().Value("id").String().Raw()
	// 	fmt.Println("field id: ", fieldID)
	// 	id, _ := strconv.ParseInt(fieldID, 10, 64)
	// 	req := &VoteFieldRequest{
	// 		ID:   id,
	// 		Type: 20,
	// 	}
	// 	resp := e.PUT("/field/vote").
	// 		WithHeader("Authorization", "Bearer "+token).
	// 		WithCookie(CookieSecret, cookieVal).
	// 		WithJSON(req).Expect().Status(http.StatusOK)
	// 	fmt.Printf("/field/vote response: %v\n", resp.Body())
	// }
	req := &VoteFieldRequest{
		ID:   1408361135484178432,
		Type: 20,
	}
	resp := e.PUT("/recommend/vote").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/recommend/vote response: %v\n", resp.Body())
}

func testGetVotes(t *testing.T) {
	id, _ := strconv.ParseInt(fieldID, 10, 64)
	req := &GetVotesRequest{
		FieldID: id,
		Type:    20,
	}
	resp := e.POST("/field/votes").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/field/votes %v response: %v\n", id, resp.Body())
}

func testGetRecords(t *testing.T) {
	req := &GetVotesRequest{
		Type:  20,
		Query: "field",
	}
	resp := e.POST("/field/records").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/field/records response: %v\n", resp.Body())
}

func testOpField(t *testing.T) {
	req := &Field{
		ID:     1396027332208103424,
		Status: 50,
	}
	resp := e.PUT("/field/status").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/field/status %v response: %v\n", req.ID, resp.Body())
}

func testGetDefinitions(t *testing.T) {
	req := &GetFieldsRequest{
		// DictID: 1407622954618982400,
		Industry:    "农、林、牧、渔业",
		SubIndustry: "农业",
		// Status:      []int32{20},
	}
	resp := e.POST("/fields/definitions").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/fields/definitions response: %v\n", resp.Body())
}

func testGetRecommends(t *testing.T) {
	req := &GetFieldsRequest{
		// DictID: 1405069047543894016,
	}
	resp := e.POST("/recommends").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/recommends response: %v\n", resp.Body())
}
