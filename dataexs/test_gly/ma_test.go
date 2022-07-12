package testgly

import (
	"fmt"
	"net/http"
	"testing"
)

type SearchMARequest struct {
	Query     string `json:"q"`
	PageIndex uint32 `json:"page_index"`
	PageSize  uint32 `json:"page_size"`
}

func testSearchMA(t *testing.T) {
	req := &SearchMARequest{
		PageSize: 10,
	}
	resp := eb.POST("/mas").
		WithCookie(backCookie, rootUserToken).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/mas pass result: ", resp.Body())
	resp.JSON().Object().Value("err_code").Equal(0)
	resp.JSON().Object().Value("data").Object().Value("list").Array().Length().Equal(10)
}
