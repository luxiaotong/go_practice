package testdatassets

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tokenKeySeller string
	tokenValSeller string
	tokenKeyBuyer  string
	tokenValBuyer  string
)

type SignInRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Vcode    string `json:"vcode"`
}

type AuditFirm struct {
	ID       int64  `json:"id,string"`
	Approved bool   `json:"approved"`
	Refuse   string `json:"refuse"`
}

type GetFirmsRequest struct {
	Query     string `json:"q"`
	Status    int32  `json:"status"`
	PageIndex uint32 `json:"page_index"`
	PageSize  uint32 `json:"page_size"`
}

func testSignInSeller(t *testing.T) {
	req := SignInRequest{
		Mobile:   "18500022713",
		Password: "123456",
	}
	resp := ep.POST("/user/signin").WithJSON(req).Expect().Status(http.StatusOK)
	tokenKeySeller = resp.Cookie(jwtCookieSecret).Value().Raw()
	tokenValSeller = resp.JSON().Object().Value("data").Object().Value("token").String().Raw()
	fmt.Println("seller token key: ", tokenKeySeller)
	fmt.Println("seller token val: ", tokenValSeller)
	assert.NotEmpty(t, tokenKeySeller)
	assert.NotEmpty(t, tokenValSeller)
}

func testSignInBuyer(t *testing.T) {
	req := SignInRequest{
		Mobile:   "15101501908",
		Password: "123456",
	}
	resp := ep.POST("/user/signin").WithJSON(req).Expect().Status(http.StatusOK)
	tokenKeyBuyer = resp.Cookie(jwtCookieSecret).Value().Raw()
	tokenValBuyer = resp.JSON().Object().Value("data").Object().Value("token").String().Raw()
	fmt.Println("buyer token key: ", tokenKeyBuyer)
	fmt.Println("buyer token val: ", tokenValBuyer)
	assert.NotEmpty(t, tokenKeyBuyer)
	assert.NotEmpty(t, tokenValBuyer)
}

func testAuditFirm(t *testing.T) {
	firmID := int64(1281908739774877696)
	req := &AuditFirm{
		ID:       firmID,
		Approved: true,
	}
	resp := eb.POST("/user/firm/status").
		WithCookie(backCookie, provUserToken).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/user/firm/status pass result: ", resp.Body())
	resp.JSON().Object().Value("err_code").Equal(0)
}

func testGetFirms(t *testing.T) {
	req := &GetFirmsRequest{
		// PageSize: 10,
	}
	resp := eb.POST("/user/firms").
		WithCookie(backCookie, rootUserToken).
		WithJSON(req).
		Expect().Status(http.StatusOK)
	fmt.Println("/user/firms pass result: ", resp.Body())
	resp.JSON().Object().Value("err_code").Equal(0)
	list := resp.JSON().Object().Value("data").Object().Value("list")
	list.Array().Length().Equal(10)
}
