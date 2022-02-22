package testdatassets

import (
	"fmt"
	"net/http"
	"testing"
)

var provUserToken string
var cityUserToken string

type Point struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type Account struct {
	Card string `json:"card"`
	Bank string `json:"bank"`
	Firm string `json:"firm"`
}

type Center struct {
	ID            int32    `json:"id"`
	Name          string   `json:"name"`
	Level         int32    `json:"level"`
	ParentID      int32    `json:"parent_id"`
	ParentName    string   `json:"parent_name"`
	AreaCode      string   `json:"area_code"`
	AreaList      []string `json:"area_list"`
	Point         *Point   `json:"point"`
	ProvinceID    string   `json:"province_id"`
	ContactFirm   string   `json:"contact_firm"`
	ContactName   string   `json:"contact_name"`
	ContactMobile string   `json:"contact_mobile"`
	ContactTel    string   `json:"contact_tel"`
	ContactEmail  string   `json:"contact_email"`
	Account       *Account `json:"Account"`
	CreateTime    string   `json:"create_time"`
	ChildNum      int32    `json:"child_num"`
}

type BackUser struct {
	ID         int32   `json:"id"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Type       int32   `json:"type"`
	CenterID   int32   `json:"center_id"`
	Center     *Center `json:"center"`
	CreateTime string  `json:"create_time"`
	UpdateTime string  `json:"update_time"`
	LoginTime  string  `json:"login_time"`
	Status     bool    `json:"status"`
}

func testLoginBackend(t *testing.T) {
	req := &BackUser{
		Username: "test1",
		Password: "pass@datassets",
	}
	resp := eb.POST("/backuser/login").WithJSON(req).Expect().Status(http.StatusOK)
	cityUserToken = resp.Cookie(backCookie).Value().Raw()
	fmt.Println("city user token in cookie: ", cityUserToken)
	fmt.Println("/backuser/login result: ", resp.Body())

	req = &BackUser{
		Username: "test2",
		Password: "pass@datassets",
	}
	resp = eb.POST("/backuser/login").WithJSON(req).Expect().Status(http.StatusOK)
	provUserToken = resp.Cookie(backCookie).Value().Raw()
	fmt.Println("prov user token in cookie: ", provUserToken)
	fmt.Println("/backuser/login result: ", resp.Body())
}
