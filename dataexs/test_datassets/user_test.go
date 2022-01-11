package testdatassets

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tokenKey, tokenVal string

type SignInRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Vcode    string `json:"vcode"`
}

func testSignIn(t *testing.T) {
	req := SignInRequest{
		Mobile:   "18500022713",
		Password: "123456",
	}

	resp := ep.POST("/user/signin").WithJSON(req).Expect().Status(http.StatusOK)
	tokenKey = resp.Cookie(jwtCookieSecret).Value().Raw()
	tokenVal = resp.JSON().Object().Value("data").Object().Value("token").String().Raw()
	fmt.Println("token key: ", tokenKey)
	fmt.Println("token val: ", tokenVal)
	assert.NotEmpty(t, tokenKey)
	assert.NotEmpty(t, tokenVal)
}
