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
