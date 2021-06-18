package main

import (
	"fmt"
	"net/http"
	"testing"
)

type GetReleasesRequest struct {
	PageIndex uint32 `json:"page_index"`
	PageSize  uint32 `json:"page_size"`
}

type GetSchemasRequest struct {
	Version     int64  `json:"version,string"`
	Query       string `json:"q"`
	Industry    string `json:"industry"`
	SubIndustry string `json:"sub_industry"`
	Token       string `json:"token"`
	PageIndex   uint32 `json:"page_index"`
	PageSize    uint32 `json:"page_size"`
}

func testGetReleases(t *testing.T) {
	req := &GetReleasesRequest{}
	resp := e.POST("/releases").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/releases response: %v\n", resp.Body())
}

func testGetSchemas(t *testing.T) {
	req := &GetSchemasRequest{}
	resp := e.POST("/schemas/gets").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/schemas/gets response: %v\n", resp.Body())
}

func testDownloadSchemas(t *testing.T) {
	req := &GetSchemasRequest{
		// Version: 26,
	}
	resp := e.POST("/schemas/download").
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/schemas/download response: %v\n", resp.Body())

	schemas := resp.JSON().Object().Value("data").Object().Value("schemas").String().Raw()
	resp = e.GET(schemas).
		WithHeader("Authorization", "Bearer "+adminToken).
		WithCookie(CookieSecret, adminCookie).
		Expect().Status(http.StatusOK)
	fmt.Printf(schemas+" result: %v", resp.Body())
}

func testSearchSchemas(t *testing.T) {
	req := &GetSchemasRequest{
		Token: "twP_r5mD7T9tH-WBnAOVZ6F8uHqooxpJVxAPnFZUYDM=",
	}
	resp := e.POST("/schemas/search").
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/schemas/search response: %v\n", resp.Body())
}

func testStats(t *testing.T) {
	resp := e.GET("/stats").Expect().Status(http.StatusOK)
	fmt.Printf("/stats response: %v\n", resp.Body())
}

func testStatsSDK(t *testing.T) {
	req := &GetReleasesRequest{}
	resp := e.POST("/stats/sdk").
		WithHeader("Authorization", "Bearer "+token).
		WithCookie(CookieSecret, cookieVal).
		WithJSON(req).Expect().Status(http.StatusOK)
	fmt.Printf("/stats/sdk response: %v\n", resp.Body())
}
